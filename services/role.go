package services

import (
	"admin/global"
	"admin/models"
	"errors"
)

type AdminRoleWhere struct {
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
	Name       string `json:"name"`
	Permission string `json:"permission"`
	All        int    `json:"all"`
}

func CreateAdminRole(ar models.AdminRole) (int, error) {
	var num int

	global.DB.
		Where(" name = ? ", ar.Name).
		First(&models.AdminRole{}).Count(&num)

	if num > 0 {
		return 0, errors.New("the name exists")
	}
	e := global.DB.Create(&ar).Error
	if e != nil {
		return 0, e
	}
	return int(ar.ID), nil
}

func GetAdminRoleById(id int) (ar models.AdminRole, e error) {
	e = global.DB.Where(" id = ? ", uint(id)).First(&ar).Error
	if e != nil {
		return ar, e
	}
	return ar, nil
}

func GetAdminRoleList(where AdminRoleWhere, offset, limit int) (ars []models.ResAdminRole, total int, e error) {
	db := global.DB.Model(&models.AdminRole{})

	if where.All != 1 {
		if where.Name != "" {
			db = db.Where("admin_role.Name like ?", "%"+where.Name+"%")
		}
		db = db.Offset(offset).Limit(limit)
	}
	db.Count(&total)

	e = db.Scan(&ars).Error
	return
}

func UpdateAdminRole(ar models.AdminRole) error {

	return global.DB.Model(&ar).Updates(ar).Error
}

func DeleteAdminRole(ar models.AdminRole) error {
	return global.DB.Delete(&ar).Error
}
