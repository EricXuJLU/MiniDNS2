// Package library 只允许依赖 Package model
package library

import (
	"MiniDNS2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"regexp"
)

//检查一个错误，若发生错误则返回true
func Check(err error, descriptions ...interface{}) (fail bool) {
	if err != nil {
		log.Println(err, descriptions)
		return true
	}
	return false
}
func OpenTheDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(model.Database), &gorm.Config{})
	Check(err, "library.OpenTherDB")
	return db
}
func IsIP(ip string) (m bool) {
	m, _ = regexp.MatchString("^((2(5[0-5]|[0-4]\\d))|1\\d{2}|[1-9]?\\d)(\\.((2(5[0-5]|[0-4]\\d))|1\\d{2}|[1-9]?\\d)){3}$", ip)
	return
}
