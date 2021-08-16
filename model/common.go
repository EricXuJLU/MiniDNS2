// Package model 不允许依赖别的包！！！
package model

const (
	Port1    = ":10086"
	Port2    = ":10010"
	Port3    = ":3985"
	Local    = "localhost"
	Address  = "localhost"
	Redis    = "localhost:6379"
	Database = "root:root@/minidns2?charset=utf8mb4&parseTime=True&loc=Local"
	GinIndex = "Welcome!\n" +
		"GetIP:\n" +
		"Insert:\n" +
		"Update:\n" +
		"Delete:\n"
)
