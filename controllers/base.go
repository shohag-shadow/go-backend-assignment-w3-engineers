package controllers

import (
	"rental-property-api/models"

	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

// func (c *BaseController) RespondSuccess(data interface{}) {
// 	c.Data["json"] = models.SuccessResponse{Status: "success", Data: data}
// 	c.ServeJSON()
// }

func (c *BaseController) RespondError(statusCode int, message string) {
	c.Ctx.Output.SetStatus(statusCode)
	c.Data["json"] = models.ErrorResponse{Error: message}
	c.ServeJSON()
}
