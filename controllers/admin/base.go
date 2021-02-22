package admin

import (
	"admin/global"
	"admin/models"
	"admin/services"
	"encoding/json"
	"github.com/astaxie/beego"
)

const SessionKey = "adminUser"

type BaseController struct {
	beego.Controller
	controllerName string
	methodName     string
	currentUser    models.AdminUser
}

type returnDataStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (this *BaseController) Prepare() {
	path := this.Ctx.Request.URL.Path
	this.getAdminUserInfo()
	pathList := services.GetMenuPathByRoleId(this.currentUser.RoleId)
	if _, ok := pathList[path]; !ok {
		//this.json(500, "user has not permission", "")
	}
}

func (this *BaseController) getAdminUserInfo() {

	this.SetSession(SessionKey, models.AdminUser{
		Account:  "1",
		RoleId:   1,
		Name:     "",
		Position: "",
		Status:   1,
	})
	session := this.GetSession(SessionKey)
	if session != nil {
		this.currentUser = session.(models.AdminUser)
	} else {
		this.json(500, "user is not login", "")
	}

}

func (this *BaseController) setAdminUser2Session(userId int) error {
	am, e := services.GetAdminUserById(userId)
	if e != nil {
		return e
	}
	this.SetSession(SessionKey, am)
	return nil
}

func (this *BaseController) json(code int, msg string, data interface{}) {

	returnJson := global.Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ret, err := json.Marshal(returnJson)
	if err != nil {
		this.Data["json"] = err.Error()
		this.ServeJSON()
	}
	_, _ = this.Ctx.ResponseWriter.Write(ret)
	this.StopRun()
}
