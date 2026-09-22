package main

import (
	"rental-property-api/models"
	_ "rental-property-api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	models.GetData()
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
