package admin

import (
	"admin/models"
	"admin/services"
	"encoding/json"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	BaseController
}

// @Title Get
// @Description get user by uid
// @Param       uid             path    string  true            "The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (this *UserController) Get() {
	i, err := this.GetInt("user_id")
	if err != nil {
		this.json(500, err.Error(), "")
	}

	au, e := services.GetAdminUserById(i)
	if e != nil {
		this.json(500, e.Error(), "")
	}
	this.json(0, "ok", au)
}

func (this *UserController) List() {
	where := services.AdminUserWhere{}
	//where.PageSize = 10
	//where.Page = 1
	//where.Status = -1
	where.PageSize, _ = this.GetInt("page_size", 10)
	where.Page, _ = this.GetInt("page", 1)
	where.Status, _ = this.GetInt("status", -1)
	where.Sort, _ = this.GetInt("sort", 1)

	where.Account = this.GetString("account", "")
	where.Name = this.GetString("name", "")

	offset := (where.Page - 1) * where.PageSize
	limit := where.PageSize

	aus, total, e := services.GetAdminUserList(where, offset, limit)
	if e != nil {
		this.json(500, e.Error(), "")
	}
	this.json(0, "ok", map[string]interface{}{"list": aus, "total": total})
}

func (this *UserController) Add() {
	var adminUser models.AdminUser
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &adminUser)
	if err != nil {
		this.json(500, err.Error(), "")

	}
	valid := validation.Validation{}
	b, err := valid.Valid(&adminUser)
	if err != nil {
		this.json(500, err.Error(), "")
	}
	if !b {
		msg := ""
		for _, err := range valid.Errors {
			msg += err.Message + ";"
		}
		this.json(500, msg, "")

		return
	}

	i, e := services.CreateAdminUser(adminUser)
	if e != nil {
		this.json(500, e.Error(), "")

	}
	data := map[string]int{
		"id": i,
	}
	this.json(0, "ok", data)
}

func (this *UserController) Update() {
	var adminUser models.AdminUser
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &adminUser)
	if err != nil {
		this.json(500, err.Error(), "")

	}
	valid := validation.Validation{}
	valid.Required(adminUser.ID, "id").Message("id is required")
	if valid.HasErrors() {
		this.json(500, valid.Errors[0].Message, "")
	}
	adminUser.Account = ""
	err = services.UpdateAdminUser(adminUser)

	if err != nil {
		this.json(500, err.Error(), "")
	}

	this.json(0, "ok", "")
}

func (this *UserController) Delete() {
	var adminUser models.AdminUser
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &adminUser)
	if err != nil {
		this.json(500, err.Error(), "")

	}
	valid := validation.Validation{}
	valid.Required(adminUser.ID, "id").Message("id is required")
	if valid.HasErrors() {
		this.json(500, valid.Errors[0].Message, "")
	}
	err = services.DeleteAdminUser(adminUser)
	if err != nil {
		this.json(500, err.Error(), "")
	}

	this.json(0, "ok", "")
}

func (this *UserController) UpdatePwd() {

	newPwd := struct {
		Account     string `json:"account"`
		Password    string `json:"password"`
		NewPassword string `json:"new_password"`
	}{}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &newPwd)
	if err != nil {
		this.json(500, err.Error(), "")

	}
	valid := validation.Validation{}
	valid.Required(newPwd.Account, "account").Message("account is required")
	valid.Required(newPwd.Password, "password").Message("password is required")
	valid.Required(newPwd.NewPassword, "new_password").Message("newPassword is required")

	if valid.HasErrors() {
		this.json(500, valid.Errors[0].Message, "")
	}
	au, e := services.GetAdminUserByAccountAndPassword(newPwd.Account, newPwd.Password)
	if e != nil {
		this.json(500, e.Error(), "")
	}
	au.Password = newPwd.NewPassword
	err = services.UpdateAdminUser(au)
	if err != nil {
		this.json(500, err.Error(), "")
	}
	this.json(0, "ok", "")

}

func (this *UserController) Status() {

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
	err = services.StatusAdminUser(data.Status, data.Id)
	if err != nil {
		this.json(500, err.Error(), "")
	}
	this.json(0, "ok", "")
}
