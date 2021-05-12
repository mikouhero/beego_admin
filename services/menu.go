package services

import (
	"admin/global"
	"admin/models"
	"errors"
	"strconv"
	"strings"
)

/**
where 条件
 */
type AdminMenuWhere struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Status   int    `json:"status"`
	Level    int    `json:"level"`
	Pid      int    `json:"pid"`
	Name     string `json:"name"`
}

/**
添加菜单
 */
func CreateAdminMenu(am models.AdminMenu) (int, error) {
	var num int
	global.DB.
		Where(" name = ? ", am.Name).
		Where("pid = ? ", am.Pid).
		First(&models.AdminMenu{}).Count(&num)

	if num > 0 {
		return 0, errors.New("the name exists")
	}
	e := global.DB.Create(&am).Error
	if e != nil {
		return 0, e
	}
	return int(am.ID), nil
}

/**
根据id 获取菜单信息
 */
func GetAdminMenuById(id int) (am models.AdminMenu, e error) {
	e = global.DB.Where(" id = ? ", uint(id)).First(&am).Error
	if e != nil {
		return am, e
	}
	return am, nil
}

func GetAdminMenuList(where AdminMenuWhere, offset, limit int) (ams []models.AdminMenu, total int, e error) {
	db := global.DB

	if where.Status != -1 {
		db.Where("status = ?", where.Status)
	}

	if where.Pid != -1 {
		db.Where("pid = ?", where.Pid)
	}
	if where.Level != -1 {
		db.Where("level = ?", where.Level)
	}
	db.Count(&total)
	e = db.Offset(offset).Limit(limit).Find(&ams).Error
	return
}

func GetAdminMenuByRoleId(roleId int, selected bool) (ams []*models.AdminMenuTreeList, e error) {

	var pids []string
	if roleId != 0 {
		ar, err := GetAdminRoleById(roleId)
		if err != nil {
			return
		}
		pids = strings.Split(ar.Permission, ",")
	} else {
		ids := GetMenuAllIds()
		for _, v := range ids {
			pids = append(pids, v)
		}
	}

	var ids []int
	for _, v := range pids {
		id, _ := strconv.Atoi(v)
		ids = append(ids, id)
	}
	menu := new(models.AdminMenu)
	ams = menu.MenuList(ids, selected)
	return
}

func UpdateAdminMenu(am models.AdminMenu) error {
	return global.DB.Model(&am).Updates(am).Error
}

func DeleteAdminMenu(am models.AdminMenu) error {
	return global.DB.Delete(&am).Error
}

func StatusAdminMenu(status, id int) error {
	i := map[string]interface{}{
		"status": status,
	}
	return global.DB.Model(&models.AdminMenu{}).Where("id = ? ", id).Update(i).Error
}

func GetMenuPathByRoleId(roleId int) map[string]bool {
	pathList := map[string]bool{}
	ar, e := GetAdminRoleById(roleId)
	if e != nil {
		return pathList
	}
	pids := strings.Split(ar.Permission, ",")
	var ids []int
	for _, v := range pids {
		id, _ := strconv.Atoi(v)
		ids = append(ids, id)
	}
	var menuPath []struct {
		Urls string `json:"urls"`
	}
	global.DB.Model(&models.AdminMenu{}).Select("urls").Where("status = 1").Where("id in (?)", ids).Scan(&menuPath)
	for _, v := range menuPath {
		pathList[v.Urls] = true
	}
	return pathList
}

func GetMenuAllIds() map[string]string {
	idList := map[string]string{}
	global.DB.Model(&models.AdminMenu{}).Select("id").Where("status = 1").Scan(&idList)
	return idList
}
