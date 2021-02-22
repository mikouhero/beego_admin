package admin

import (
	"admin/models"
	"admin/services"
	"encoding/json"
	"github.com/astaxie/beego/validation"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) Get() {
	i, err := this.GetInt("role_id")
	if err != nil {
		this.json(500, err.Error(), "")
	}

	au, e := services.GetAdminRoleById(i)
	if e != nil {
		this.json(500, e.Error(), "")
	}
	this.json(0, "ok", au)
}

func (this *RoleController) List() {
	where := services.AdminRoleWhere{}
	where.PageSize, _ = this.GetInt("page_size", 10)
	where.Page, _ = this.GetInt("page", 1)
	where.Name = this.GetString("name")
	where.All, _ = this.GetInt("all", 0)

	offset := (where.Page - 1) * where.PageSize
	limit := where.PageSize

	aus, total, e := services.GetAdminRoleList(where, offset, limit)
	if e != nil {
		this.json(500, e.Error(), "")
	}
	this.json(0, "ok", map[string]interface{}{
		"list":  aus,
		"total": total,
	})
}

func (this *RoleController) Add() {
	var AdminRole models.AdminRole
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &AdminRole)
	if err != nil {
		this.json(500, err.Error(), "")

	}
	valid := validation.Validation{}
	b, err := valid.Valid(&AdminRole)
	if err != nil {
		this.json(500, err.Error(), "")
	}
	if !b {
		msg := ""
		for _, err := range valid.Errors {
			msg += err.Message + ";"
		}
		this.json(500, msg, "")
	}

	i, e := services.CreateAdminRole(AdminRole)
	if e != nil {
		this.json(500, e.Error(), "")

	}
	data := map[string]int{
		"id": i,
	}
	this.json(0, "ok", data)
}

func (this *RoleController) Update() {
	var AdminRole models.AdminRole
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &AdminRole)
	if err != nil {
		this.json(500, err.Error(), "")

	}
	valid := validation.Validation{}
	valid.Required(AdminRole.ID, "id").Message("id is required")
	if valid.HasErrors() {
		this.json(500, valid.Errors[0].Message, "")
	}

	err = services.UpdateAdminRole(AdminRole)

	if err != nil {
		this.json(500, err.Error(), "")
	}

	this.json(0, "ok", "")
}

func (this *RoleController) Delete() {
	var AdminRole models.AdminRole
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &AdminRole)
	if err != nil {
		this.json(500, err.Error(), "")
	}
	valid := validation.Validation{}
	valid.Required(AdminRole.ID, "id").Message("id is required")
	if valid.HasErrors() {
		this.json(500, valid.Errors[0].Message, "")
	}
	err = services.DeleteAdminRole(AdminRole)
	if err != nil {
		this.json(500, err.Error(), "")
	}

	this.json(0, "ok", "")
}

func (this *RoleController) RoleAuth() {

	ams, e := services.GetAdminMenuByRoleId(0, true)
	if e != nil {
		this.json(500, e.Error(), "")
	}

	this.json(0, "ok", map[string]interface{}{"menuList": ams})

}
