package controllers

import (
	"rental-property-api/models"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	beego.Controller
}

func (o *PropertyController) DemoResponse() {
	o.Data["json"] = models.Response{
		Status:  "succed",
		Messege: "this is a demo messege",
	}
	o.ServeJSON()
}
