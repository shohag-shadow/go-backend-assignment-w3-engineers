package main

import (
	_ "rental-property-api/routers"
	"rental-property-api/services"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	services.GetData("")
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
