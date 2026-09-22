package controllers

import (
	"rental-property-api/models"
	"rental-property-api/services"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	beego.Controller
}

func (o *PropertyController) DemoResponse() {
	data := models.GetData()

	o.Data["json"] = models.GetResponseFromSource(&data.Properties[0])
	o.ServeJSON()
}
func (c *PropertyController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	prop, _ := services.GetProertyById(id)
	c.Data["json"] = prop
	c.ServeJSON()
}
