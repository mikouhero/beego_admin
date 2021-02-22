package initialiaze

import (
	"admin/global"
)

func init() {
	Mysql()
	Migrate()
	Session()
}

func Defer() {
	 global.DB.Close()
}
