// Package model 不允许依赖别的包！！！
package model
import (
	"gorm.io/gorm"
)
type DNS struct {
	gorm.Model
	Domain string
	IP string
}
