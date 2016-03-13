package tomato

import (
	_ "github.com/lfq7413/tomato/rest"
	_ "github.com/lfq7413/tomato/routers"

	"github.com/astaxie/beego"
	"github.com/lfq7413/tomato/orm"
)

// Run ...
func Run() {

	orm.OpenDB()
	defer orm.CloseDB()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
