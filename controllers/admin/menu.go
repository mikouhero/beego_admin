package admin

import (
	"admin/models"
	"admin/services"
	"encoding/json"
	"github.com/astaxie/beego/validation"
)

type MenuController struct {
	BaseController
}

func (this *MenuController) Get() {
	i, err := this.GetInt("role_id")
	if err != nil {
		this.json(500, err.Error(), "")
	}

	au, e := services.GetAdminMenuById(i)
	if e != nil {
		this.json(500, e.Error(), "")
	}
	this.json(0, "ok", au)
}

func (this *MenuController) List() {
	where := services.AdminMenuWhere{}
	where.PageSize = 10
	where.Page = 1
	where.Status = -1
	where.Level = -1
	where.Pid = -1
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &where)
	if err != nil {
		this.json(500,err.Error(),"")
	}

	offset := (where.Page - 1) * where.PageSize
	limit := where.PageSize

	aus, total, e := services.GetAdminMenuList(where, offset, limit)
	if e != nil {
		this.json(500, e.Error(), "")
	}
	this.json(0, "ok", map[string]interface{}{"list": aus, "total": total})
}

func (this *MenuController) Add() {
	var AdminMenu models.AdminMenu
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &AdminMenu)
	if err != nil {
		this.json(500,err.Error(),"")

	}
	valid := validation.Validation{}
	b, err := valid.Valid(&AdminMenu)
	if err != nil {
		this.json(500,err.Error(),"")

	}
	if !b {
		msg := ""
		for _, err := range valid.Errors {
			msg += err.Message + ";"
		}
		this.json(500,err.Error(),"")

	}

	i, e := services.CreateAdminMenu(AdminMenu)
	if e != nil {
		this.json(500,err.Error(),"")

	}
	data := map[string]int{
		"menu_id": i,
	}
	this.json(0, "ok", data)
}

func (this *MenuController) Update() {
	var AdminMenu models.AdminMenu
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &AdminMenu)
	if err != nil {
		this.json(500, err.Error(), "")
	}
	valid := validation.Validation{}
	valid.Required(AdminMenu.ID, "id").Message("id is required")
	if valid.HasErrors() {
		this.json(500, valid.Errors[0].Message, "")
	}

	err = services.UpdateAdminMenu(AdminMenu)

	if err != nil {
		this.json(500, err.Error(), "")
	}

	this.json(0, "ok", "")
}

func (this *MenuController) Delete() {
	var AdminMenu models.AdminMenu
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &AdminMenu)
	if err != nil {
		this.json(500, err.Error(), "")

	}
	valid := validation.Validation{}
	valid.Required(AdminMenu.ID, "id").Message("id is required")
	if valid.HasErrors() {
		this.json(500, valid.Errors[0].Message, "")
	}
	err = services.DeleteAdminMenu(AdminMenu)
	if err != nil {
		this.json(500, err.Error(), "")
	}

	this.json(0, "ok", "")
}

func (this *MenuController) Status() {

	data := struct {
		Id     int `json:"id"`
		Status int `json:"status"`
	}{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &data)
	if err != nil {
		this.json(500, err.Error(), "")

	}
	valid := validation.Validation{}
	valid.Required(data.Id, "id").Message("id is required")
	valid.Min(data.Status, 0, "status")
	valid.Max(data.Status, 1, "status")

	if valid.HasErrors() {
		this.json(500, valid.Errors[0].Message, "")
	}
	err = services.StatusAdminMenu(data.Status, data.Id)
	if err != nil {
		this.json(500, err.Error(), "")
	}
	this.json(0, "ok", "")
}

