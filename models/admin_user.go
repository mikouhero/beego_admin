package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
	"io"
)

type AdminUser struct {
	gorm.Model
	Account  string `valid:"Required" json:"account" gorm:"comment:'用户名';type:varchar(30);index:idx_name;not null"`
	Password string `valid:"Required" json:"password" gorm:"comment:'密码';type:varchar(50);not null"`
	RoleId   int    `valid:"Required" json:"role_id" gorm:"comment:'角色id';not null;default:0"`
	Name     string `valid:"Required" json:"name" gorm:"comment:'真实姓名';type:varchar(30);not null;default:''"`
	Position string `valid:"Required" json:"position" gorm:"comment:'职位';type:varchar(30);not null;default:''"`
	Status   int    `json:"status" gorm:"type:tinyint(1);comment:'1开启 0禁用';default:0"`
	HeadImg  string `valid:"Required" json:"head_img" gorm:"comment:'头像';type:varchar(255);not null;default:''"`
}

type ResAdminUser struct {
	Id       int    `json:"id"`
	Account  string `json:"account" `
	RoleId   int    `json:"role_id" `
	Name     string `json:"name"`
	Position string `json:"position"`
	Status   int    `json:"status"`
	RoleName string `json:"role_name"`
	HeadImg  string `json:"head_img"`
}

func (AdminUser) TableName() string {
	return "admin_user"
}

func (au *AdminUser) Valid(v *validation.Validation) {

}

func (au *AdminUser) EncodePwd(password string) string {
	m := md5.New()
	io.WriteString(m, password)
	au.Password = hex.EncodeToString(m.Sum(nil))
	return au.Password
}

func (au *AdminUser) BeforeCreate() {
	au.ID = 0
	au.EncodePwd(au.Password)
}

func (au *AdminUser) AfterFind() () {
	au.Password = ""
}
