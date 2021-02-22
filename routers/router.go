// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"admin/controllers/admin"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	beego.Get("/", func(ctx *context.Context) {
		_ = ctx.Output.Body([]byte("hello world"))
	})

	beego.Router("/admin/user", &admin.UserController{}, "*:Get")
	beego.Router("/admin/users", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/update", &admin.UserController{}, "*:Update")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")
	beego.Router("/admin/user/updatePwd", &admin.UserController{}, "*:UpdatePwd")
	beego.Router("/admin/user/status", &admin.UserController{}, "*:Status")

	beego.Router("/admin/menu", &admin.MenuController{}, "*:Get")
	beego.Router("/admin/menus", &admin.MenuController{}, "*:List")
	beego.Router("/admin/menu/add", &admin.MenuController{}, "*:Add")
	beego.Router("/admin/menu/update", &admin.MenuController{}, "*:Update")
	beego.Router("/admin/menu/delete", &admin.MenuController{}, "*:Delete")
	beego.Router("/admin/menu/status", &admin.MenuController{}, "*:Status")

	beego.Router("/admin/role", &admin.RoleController{}, "*:Get")
	beego.Router("/admin/roles", &admin.RoleController{}, "*:List")
	beego.Router("/admin/role/add", &admin.RoleController{}, "*:Add")
	beego.Router("/admin/role/update", &admin.RoleController{}, "*:Update")
	beego.Router("/admin/role/delete", &admin.RoleController{}, "*:Delete")
	beego.Router("/admin/role/roleAuth", &admin.RoleController{}, "*:RoleAuth")

	beego.Router("/admin/login", &admin.LoginController{}, "*:Login")
	beego.Router("/admin/info", &admin.LoginController{}, "*:Info")
	beego.Router("/admin/logout", &admin.LoginController{}, "*:LogOut")

	beego.Router("/admin/task", &admin.TaskController{}, "*:Get")
	beego.Router("/admin/upload", &admin.UploadController{}, "*:Upload")


}
