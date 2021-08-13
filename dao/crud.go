package dao

import (
	. "MiniDNS2/library"
	"MiniDNS2/model"
	"context"
)

///所有操作均不接受*值和空值
///因domain和ip存在一对多的关系，故在redis中采用集合存储，domain是键，ip是值

//同时插入redis和mysql，mysql插入之前先查表，不会重复插入
func (dao *Dao)Insert(ctx context.Context, domain, ip string) (string) {
	if domain=="" || !IsIP(ip) {
		return "不合理的请求"
	}
	//redis
	r := dao.r
	err := r.SAdd(ctx, domain, ip).Err()
	if Check(err, "dao.Insert redis.SAdd error") {
		return "Redis插入失败"
	}

	//mysql
	db := dao.db
	var tmp []model.DNS
	db.Model(&model.DNS{}).Where("domain=? and ip=?", domain, ip).Find(&tmp)
	if len(tmp) != 0 {
		return "Mysql此条目已存在"
	}
	dns := model.DNS{Domain: domain, IP: ip}
	db.Create(&dns)
	return "插入成功"
}

//优先查询redis，查不到再查mysql
func (dao *Dao)GetIP(ctx context.Context, domain string) (ips []string) {
	//redis
	r := dao.r
	ips, err := r.SMembers(ctx, domain).Result()
	if !Check(err, "dao.GetIP redis.Get error") {	//如果没出错，就将redis中的查询结果返回
		return
	}

	//mysql
	db := dao.db
	var tmp []model.DNS
	db.Model(&model.DNS{}).Where("domain=?", domain).Find(&tmp)
	for _, i := range tmp {
		ips = append(ips, i.IP)
	}
	return
}

//同时更新redis和mysql，mysql中允许重复，redis用的是集合，不会重复
func (dao *Dao)Update(ctx context.Context, domainsrc, ipsrc, domaindst, ipdst string) (int) {
	if domainsrc=="" || !IsIP(ipsrc) || domaindst=="" || !IsIP(ipdst) {
		return 0
	}
	//redis
	r := dao.r
	err := r.SRem(ctx, domainsrc, ipsrc).Err()
	if Check(err, "r.Srem error in dao.Update") {	//更新失败（删除失败）
		return 0
	}
	err = r.SAdd(ctx, domaindst, ipdst).Err()
	if Check(err, "r.SAdd error in dao.Update") {	//更新失败（插入失败），回滚
		err = r.SAdd(ctx, domainsrc, ipsrc).Err()	//回滚
		if Check(err, "r.SAdd error in dao.Update 2") {
			panic("Redis和Mysql不一致！！！")	//回滚失败，严重错误
		}
		return 0
	}

	//mysql
	var src []model.DNS
	db := dao.db
	db.Model(&model.DNS{}).Where("domain=? and ip=?", domainsrc, ipsrc).Find(&src)
	for _, i := range src {
		i.Domain = domaindst
		i.IP = ipdst
		db.Save(&i)
	}
	return len(src)
}

//同时删除redis和mysql
func (dao *Dao)Delete(ctx context.Context, domain, ip string) (int) {
	if domain=="" || !IsIP(ip) {
		return 0
	}
	//redis
	r := dao.r
	err := r.SRem(ctx, domain, ip).Err()
	if Check(err, "redis.SRem error in dao.Delete") {	//redis删除出错，两个都不删
		return 0
	}

	//mysql
	var tmp []model.DNS
	db := dao.db
	db.Model(&model.DNS{}).Where("domain=? and ip=?", domain, ip).Find(&tmp)
	for _, i := range tmp {
		db.Delete(&i)
	}
	return len(tmp)
}