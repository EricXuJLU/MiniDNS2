package service

import (
	"MiniDNS2/library"
	"MiniDNS2/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

//本文件包含HTTP, Gin, gRPC, Go-kit的中间件

///http
type httpMiddleware func(http.HandlerFunc) http.HandlerFunc

func HttpLogger() httpMiddleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			value, err := readRequestBody(r) //记录数据
			library.Check(err, "readRequestBody error in web.Middlewares.httpLogger")
			next(w, r) //执行下一个中间件
			//统一打印日志
			log.Println(r.Method, r.URL, r.Header.Get("Content-Type"))
			if len(r.URL.Query()) != 0 {
				log.Print("URL Params: ")
				for i, j := range r.URL.Query() {
					log.Print(i, " - ", j, "  ")
				}
			}
			if len(value) != 0 {
				log.Print("Body Params: ")
				for i, j := range value {
					log.Print(i, " - ", j, "  ")
				}
			}
			log.Println("响应时间：", time.Since(start).Microseconds(), "微秒")
		}
	}
}

func HttpInterceptor(method string) httpMiddleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if strings.EqualFold(r.Method, method) {
				next(w, r) //继续执行下一个
				return
			}
			w.WriteHeader(http.StatusBadRequest) //同时拦截！
			_, err := fmt.Fprintln(w, "请求方法错误！期望收到", method, "实际收到", r.Method)
			library.Check(err, "fmt.Fprintln error in web.Middlewares.httpInterceptor")
		}
	}
}

func HttpChain(f http.HandlerFunc, middlewares ...httpMiddleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func readRequestBody(r *http.Request) (data map[string]interface{}, err error) { //application/json
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	_ = r.Body.Close()
	r.Body = ioutil.NopCloser(bytes.NewReader(buf))
	if len(buf) != 0 {
		err = json.Unmarshal(buf, &data)
		if err != nil {
			return map[string]interface{}{}, err
		}
	}
	return data, nil
}

///Gin
func GinLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()                            //记录时间
		value, err := readRequestBody(context.Request) //拷贝Body
		library.Check(err, "readRequestBody error in web.Middlewares.ginLogger")
		context.Next() //先执行下一项
		//统一打印日志
		log.Println(context.Request.Method, context.Request.URL, context.Request.Header.Get("Content-Type"))
		if len(context.Request.URL.Query()) != 0 {
			log.Print("URL Params: ")
			for i, j := range context.Request.URL.Query() {
				log.Print(i, " - ", j, "  ")
			}
		}
		if len(value) != 0 {
			log.Print("Body Params: ")
			for i, j := range value {
				log.Print(i, " - ", j, "  ")
			}
		}
		log.Println("响应时间：", time.Since(start).Microseconds(), "微秒")
	}
}

func ginInterceptor(method string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if strings.EqualFold(context.Request.Method, method) {
			context.Next()
		} else {
			context.Writer.WriteHeader(http.StatusBadRequest)
			_, err := context.Writer.WriteString("请求错误！期望收到 " + method + "实际收到 " + context.Request.Method)
			library.Check(err, "gin.Context.Writer.WriteString error in web.Middlewares.ginInterceptor")
		}
	}
}

///gRPC
func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	start := time.Now()
	defer func() {
		in, _ := json.Marshal(req)
		out, _ := json.Marshal(reply)
		inStr, outStr := string(in), string(out)
		duration := int64(time.Since(start) / time.Millisecond)

		log.Println("grpc", method, "in", inStr, "out", outStr, "err", err, "duration/ms", duration)

	}()

	return invoker(ctx, method, req, reply, cc, opts...)
}

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	remote, _ := peer.FromContext(ctx)
	remoteAddr := remote.Addr.String()

	in, _ := json.Marshal(req)
	inStr := string(in)
	log.Println("ip", remoteAddr, "access_start", info.FullMethod, "in", inStr)

	start := time.Now()
	defer func() {
		out, _ := json.Marshal(resp)
		outStr := string(out)
		duration := int64(time.Since(start) / time.Millisecond)
		log.Println("ip", remoteAddr, "access_end", info.FullMethod, "in", inStr, "out", outStr, "err", err, "duration/ms", duration)
	}()

	resp, err = handler(ctx, req)

	return
}

///go-kit
func GoKitLogger(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		defer func(begin time.Time) {
			typeOfReq := reflect.TypeOf(request)
			switch typeOfReq {
			case reflect.TypeOf(model.GetReq{}): //GetIP request
				req := request.(model.GetReq)
				log.Printf("GetIP\tDomain:%s\n", req.Domain)
			case reflect.TypeOf(model.InsertReq{}): //Insert request
				req := request.(model.InsertReq)
				log.Printf("Insert\tDomain:%s\tIP:%s\n", req.Domain, req.IP)
			case reflect.TypeOf(model.UpdateReq{}): //Update request
				req := request.(model.UpdateReq)
				log.Printf("Update\tDmsrc:%s\tIPsrc:%s\tDmdst:%s\tIPdst:%s\n", req.Domainsrc, req.IPsrc, req.Domaindst, req.IPdst)
			case reflect.TypeOf(model.DeleteReq{}): //Delete request
				req := request.(model.DeleteReq)
				log.Printf("Delete\tDomain:%s\tIP:%s\n", req.Domain, req.IP)
			default:
				log.Printf("Unkown Request!!")
				log.Println(request)
			}
			log.Println("Response Time:", time.Since(begin).Microseconds(), "us")
		}(time.Now())

		resp, err := next(ctx, request)
		log.Println("Response: ", resp)
		return resp, err
	}
}
