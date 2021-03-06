package service

import (
	"MiniDNS2/dao"
	"MiniDNS2/library"
	"MiniDNS2/model"
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var Srvc = &Service{}

func InitService() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})
	_, err := rdb.Ping(context.Background()).Result()
	library.Check(err, "redis init error in web.init")
	db := library.OpenTheDB()
	db.AutoMigrate(&model.DNS{})
	Srvc.Dao = dao.NewDao(db, rdb)
}

type Service struct {
	Dao *dao.Dao
}

func (srvs *Service) GetIP(ctx context.Context, req *model.GetReq) *model.GetResp {
	resp := new(model.GetResp)
	resp.Domain = req.Domain
	resp.IPs = srvs.Dao.GetIP(ctx, req.Domain)
	return resp
}

func (srvs *Service) Insert(ctx context.Context, req *model.InsertReq) *model.InsertResp {
	resp := new(model.InsertResp)
	resp.Domain = req.Domain
	resp.IP = req.IP
	if req.Domain == "" || !library.IsIP(req.IP) { //输入值约束
		resp.Result = "插入失败，不合理的请求"
		return resp
	}
	resp.Result = srvs.Dao.Insert(ctx, req.Domain, req.IP)
	return resp
}

func (srvs *Service) Update(ctx context.Context, req *model.UpdateReq) *model.UpdateResp {
	resp := new(model.UpdateResp)
	if req.Domainsrc == "" || !library.IsIP(req.IPsrc) || req.Domaindst == "" || !library.IsIP(req.IPdst) {
		resp.Affected = 0
		resp.Result = "更新失败，不合理的请求"
		return resp
	}
	resp.Affected = srvs.Dao.Update(ctx, req.Domainsrc, req.IPsrc, req.Domaindst, req.IPdst)
	resp.Result = strconv.FormatInt(int64(resp.Affected), 10) + "条记录被更新"
	return resp
}

func (srvs *Service) Delete(ctx context.Context, req *model.DeleteReq) *model.DeleteResp {
	resp := new(model.DeleteResp)
	if req.Domain == "" || !library.IsIP(req.IP) {
		resp.Affected = 0
		resp.Result = "删除失败，不合理的请求"
		return resp
	}
	resp.Affected = srvs.Dao.Delete(ctx, req.Domain, req.IP)
	resp.Result = strconv.FormatInt(int64(resp.Affected), 10) + "条记录被删除"
	return resp
}
