package services

import (
	"admin/global"
	"admin/models"
	"errors"
	"fmt"
)

type AdminUserWhere struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Status    int    `json:"status"`
	Account   string `json:"account"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	CreatedAt string `json:"created_at"`
}

func CreateAdminUser(au models.AdminUser) (int, error) {
	var num int

	global.DB.Where("account=?", au.Account).First(&models.AdminUser{}).Count(&num)

	if num > 0 {
		return 0, errors.New("the account exists")
	}
	e := global.DB.Create(&au).Error
	if e != nil {
		return 0, e
	}
	return int(au.ID), nil
}

func GetAdminUserById(id int) (au models.ResAdminUser, e error) {
	e = global.DB.Model(&models.AdminUser{}).
		Select("admin_user.*,admin_role.name as role_name").
		Joins("left join admin_role on admin_role.id = admin_user.role_id").
		Where(" admin_user.id = ? ", uint(id)).Scan(&au).Error
	if e != nil {
		return au, e
	}
	return au, nil
}

func GetAdminUserList(where AdminUserWhere, offset, limit int) (aul []models.ResAdminUser, total int, e error) {
	db := global.DB.Model(&models.AdminUser{}).
		Select("admin_user.id,admin_user.account,admin_user.head_img,admin_user.name,admin_user.role_id,admin_user.position,admin_user.status,admin_role.name as role_name")

	if where.Status != -1 {
		db.Where("admin_user.status = ?", where.Status)
	}

	if where.Account != "" {
		db = db.Where("admin_user.account like ?", "%"+where.Account+"%")
	}
	if where.Name != "" {
		db = db.Where("admin_user.Name like ?", "%"+where.Name+"%")
	}
	if where.CreatedAt != "" {
		db = db.Where("admin_user.created_at > ?", where.CreatedAt)
	}
	db.Count(&total)
	db = db.Offset(offset).
		Joins("left join admin_role  on admin_role.id = admin_user.role_id ")
	if where.Sort == 1 {
		db = db.Order("admin_user.id asc")
	} else {
		db = db.Order("admin_user.id desc")
	}
	e = db.Limit(limit).Scan(&aul).Error
	return
}

func GetAdminUserByAccountAndPassword(account, password string) (au models.AdminUser, e error) {
	e = global.DB.Where(" account = ? ", account).Where("password = ?", au.EncodePwd(password)).First(&au).Error
	if e != nil {
		return au, e
	}
	return au, nil
}

func UpdateAdminUser(au models.AdminUser) error {
	if au.Password != "" {
		au.EncodePwd(au.Password)
	}
	return global.DB.Model(&au).Updates(au).Error
}

func DeleteAdminUser(au models.AdminUser) error {
	return global.DB.Delete(&au).Error
}

func StatusAdminUser(status, id int) error {
	fmt.Println(1)
	i := map[string]interface{}{
		"status": status,
	}
	return global.DB.Model(&models.AdminUser{}).Where("id = ? ", id).Update(i).Error
}
