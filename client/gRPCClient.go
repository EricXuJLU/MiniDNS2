package client

import (
	"MiniDNS2/library"
	"MiniDNS2/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"strings"
)
const (
	SPLIT = "`"
)
type gRPCClient struct {
	conn *grpc.ClientConn
	Client proto.DNSClient
}

func GRPCClient(addr string) {
	//客户端连接服务端
	client := &gRPCClient{}
	var err error
	client.conn, err = grpc.Dial(addr, grpc.WithInsecure())
	library.Check(err, "grpc.Dial error in client.GRPCClient")
	defer  client.conn.Close()
	//获得grpc句柄
	client.Client = proto.NewDNSClient(client.conn)
	fmt.Printf("命令以 %s 字符分割\n", SPLIT)
	for {
		client.handle()
	}

}

func (c gRPCClient)handle(){
	var order string
	_, err := fmt.Scan(&order)
	library.Check(err, "fmt.Scanln error in client.handle")
	orders := strings.Split(order, SPLIT)
	switch strings.ToLower(orders[0]) {
	case "get": handleGet(c.Client, orders[1]);
	case "insert": handleInsert(c.Client, orders[1], orders[2]);
	case "delete": handleDelete(c.Client, orders[1], orders[2]);
	case "update": handleUpdate(c.Client, orders[1], orders[2], orders[3], orders[4]);
	default:
		fmt.Println("Invalid Order!")
	}
}

func handleGet(c proto.DNSClient, domain string) {
	re1, err := c.GetIP(context.Background(), &proto.GetReq{Domain: domain})
	library.Check(err, "proto.DNSClient.GetIP error in client.handleGet")
	if len(re1.IPs) == 0 {
		fmt.Printf("未找到域名：%q\n", re1.Domain)
		return
	}
	fmt.Printf("域名%q对应的IP有：\n", re1.Domain)
	for _, i := range re1.IPs{
		fmt.Println(i)
	}
}

func handleInsert(c proto.DNSClient, domain, ip string){
	re2, err := c.Insert(context.Background(), &proto.InsertReq{Domain: domain, IP: ip})
	library.Check(err, "proto.DNSClient.Insert error in client.handleInsert")
	fmt.Println(re2.Result)
}

func handleDelete(c proto.DNSClient, domain, ip string){
	re2 ,err := c.Delete(context.Background(), &proto.DeleteReq{
		Domain: domain,
		IP:   ip,
	})
	library.Check(err, "proto.DNSClient.Delete error in client.handleDelete")
	fmt.Println(re2.Result)
}

func handleUpdate(c proto.DNSClient, dmsrc, ipsrc, dmdst, ipdst string){
	re2 ,err := c.Update(context.Background(), &proto.UpdateReq{
		Domainsrc: dmsrc,
		IPsrc: ipsrc,
		Domaindst: dmdst,
		IPdst: ipdst,
	})
	library.Check(err, "proto.DNSClient.Update error in client.handleUpdate")
	fmt.Println(re2.Result)
}