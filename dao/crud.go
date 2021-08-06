package dao

import (
	. "MiniDNS2/library"
	"MiniDNS2/model"
)

///所有操作均不接受*值
//插入之前先查表，不会重复插入
func Insert(domain, ip string) (string) {
	db := OpenTheDB()
	var tmp []model.DNS
	db.Model(&model.DNS{}).Where("domain=? and ip=?", domain, ip).Find(&tmp)
	if len(tmp) != 0 {
		return "此条目已存在"
	}
	dns := model.DNS{Domain: domain, IP: ip}
	db.Create(&dns)
	return "插入成功"
}

//查询
func GetIP(domain string) (ips []string) {
	db := OpenTheDB()
	var tmp []model.DNS
	db.Model(&model.DNS{}).Where("domain=?", domain).Find(&tmp)
	for _, i := range tmp {
		ips = append(ips, i.IP)
	}
	return
}

//更新，允许重复
func Update(domainsrc, ipsrc, domaindst, ipdst string) (int) {
	var src []model.DNS
	db := OpenTheDB()
	db.Model(&model.DNS{}).Where("domain=? and ip=?", domainsrc, ipsrc).Find(&src)
	for _, i := range src {
		i.Domain = domaindst
		i.IP = ipdst
		db.Save(&i)
	}
	return len(src)
}

//删除
func Delete(domain, ip string) (int) {
	var tmp []model.DNS
	db := OpenTheDB()
	db.Model(&model.DNS{}).Where("domain=? and ip=?", domain, ip).Find(&tmp)
	for _, i := range tmp {
		db.Delete(&i)
	}
	return len(tmp)
}