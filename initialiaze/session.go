package initialiaze

import (
	"github.com/astaxie/beego"
)


func Session() {
	beego.BConfig.WebConfig.Session.SessionOn = true
}
