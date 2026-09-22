package controllers

import (
	"rental-property-api/models"
	"rental-property-api/services"
)

type PropertyController struct {
	BaseController
}

func (o *PropertyController) DemoResponse() {
	data := models.GetData()

	o.Data["json"] = models.GetResponseFromSource(&data.Properties[0])
	o.ServeJSON()
}
func (c *PropertyController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		c.RespondError(404, "Please enter a valid id")
		return
	}
	prop, err := services.GetProertyByID(id)
	if err != nil {
		c.RespondError(404, err.Error())
		return
	}
	c.Data["json"] = prop
	c.ServeJSON()
}
func (c *PropertyController) GetAll() {

}
