package models

import (
	"admin/global"
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
)

type AdminMenu struct {
	gorm.Model
	Level  int    `valid:"Required" json:"level" gorm:"type:tinyint(1);default:0;comment:'级别'"`
	Pid    int    `json:"pid" gorm:"type:int(11);default:0;comment:'父级id'"`
	Name   string `valid:"Required" json:"name" gorm:"type:varchar(30);comment:'权限名';not null"`
	Icon   string `json:"icon" gorm:"type:varchar(30);default:'';comment:'icon'"`
	Urls   string `json:"urls" gorm:"type:varchar(50);default:'';comment:'菜单地址'"`
	Remark string `json:"remark" gorm:"type:varchar(50);default:'';comment:'备注'"`
	Sort   int    `valid:"Required" json:"sort" gorm:"type:int(11);default:0;comment:'排序'"`
	Status int    `valid:"Required" json:"status" gorm:"type:tinyint(1);default:1;comment:'1开启 0禁用'"`
}

type AdminMenuTreeList struct {
	Id       int				`json:"id"`
	Level    int                  `json:"level"`
	Pid      int                  `json:"pid"`
	Name     string               `json:"name"`
	Icon     string               `json:"icon"`
	Urls     string               `json:"urls"`
	Remark   string               `json:"remark"`
	Sort     int                  `json:"sort"`
	Status   int                  `json:"status"`
	Child    []*AdminMenuTreeList `json:"child"`
	Selected bool                 `json:"selected"`
}

func (AdminMenu) TableName() string {
	return "admin_menu"
}

func (au *AdminMenu) Valid(v *validation.Validation) {

}

func (am *AdminMenu) BeforeCreate() {
	am.ID = 0
}

func (am *AdminMenu) MenuList(ids []int,selected bool) []*AdminMenuTreeList {
	return am.getMenu(0, ids, selected)
}

func (am *AdminMenu) getMenu(pid int, ids []int, selected bool) []*AdminMenuTreeList {
	var menu []AdminMenu
	db := global.DB.Model(AdminMenu{}).Where("pid = ? ", pid).Where("status =1")
	if len(ids) > 0 && selected == false {
		db = db.Where(ids)
	}
	db.Scan(&menu)
	treeList := []*AdminMenuTreeList{}
	for _, v := range menu {
		child := v.getMenu(int(v.ID), ids, selected)
		node := &AdminMenuTreeList{
			Id:     int(v.ID),
			Level:  v.Level,
			Pid:    v.Pid,
			Name:   v.Name,
			Icon:   v.Icon,
			Urls:   v.Urls,
			Remark: v.Remark,
			Sort:   v.Sort,
			Status: v.Status,
		}
		if InArray(ids, int(v.ID)) {
			node.Selected = true
		}
		node.Child = child
		treeList = append(treeList, node)
	}
	return treeList

}
func InArray(obj []int, target int) bool {
	for _, v := range obj {
		if v == target {
			return true
		}
	}
	return false
}

