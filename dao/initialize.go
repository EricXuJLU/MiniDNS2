package dao

import (
	. "MiniDNS2/library"
	"MiniDNS2/model"
)

func init() {
	db := OpenTheDB()
	db.AutoMigrate(&model.DNS{})
}
