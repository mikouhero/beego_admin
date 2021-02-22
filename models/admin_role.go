package models

import (
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
)

type AdminRole struct {
	gorm.Model
	Name       string `valid:"Required" json:"name" gorm:"type:varchar(20);comment:'角色名';default:''"`
	Permission string `valid:"Required" json:"permission" gorm:"type:text;comment:'权限Id集合'"`
	Remark     string `json:"remark" gorm:"type:varchar(50);comment:'说明'"`
}

type ResAdminRole struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
	Remark     string `json:"remark"`
}

func (AdminRole) TableName() string {
	return "admin_role"
}
func (au *AdminRole) Valid(v *validation.Validation) {

}

func (ar *AdminRole) BeforeCreate() {
	ar.ID = 0
}
