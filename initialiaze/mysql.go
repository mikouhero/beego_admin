package initialiaze

import (
	"admin/global"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Mysql() {
	var err error
	host := beego.AppConfig.String("dbhost")
	port := beego.AppConfig.String("dbport")
	user := beego.AppConfig.String("dbuser")
	pwd := beego.AppConfig.String("dbpassword")
	name := beego.AppConfig.String("dbname")
	charset := beego.AppConfig.String("dbcharset")
	maxCoonns, _ := beego.AppConfig.Int("dbmaxCoonns")
	global.DB, err = gorm.Open("mysql", user+":"+pwd+"@("+host+":"+port+")/"+name+"?charset="+charset+"&parseTime=true")
	if err != nil {
		logs.Error(err)
		return
	}
	global.DB.DB().SetMaxIdleConns(maxCoonns)
	global.DB.LogMode(true)

}
