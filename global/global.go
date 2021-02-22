package global

import "github.com/jinzhu/gorm"

// 定义全局变量
var (
	DB *gorm.DB
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
