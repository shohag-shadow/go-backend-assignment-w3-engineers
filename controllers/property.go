package controllers

import (
	"rental-property-api/models"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	beego.Controller
}

func (o *PropertyController) DemoResponse() {
	data := models.GetData()
	o.Data["json"] = data
	o.ServeJSON()
}
