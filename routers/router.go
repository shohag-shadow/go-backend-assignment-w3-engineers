// @APIVersion 1.0.0
// @Title Renatal Property API
// @Description API to browse rental properties. Supports filtering by price,star rating, review score, number of reviews, published status,property type, feed, bedroom count, and amenities.
// @Contact mdshohagshowdagor001@gmail.com
package routers

import (
	"rental-property-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/properties",
			beego.NSRouter("", &controllers.PropertyController{}, "get:GetAll"),
			beego.NSRouter("/:id", &controllers.PropertyController{}, "get:GetOne"),
		),
	)
	beego.AddNamespace(ns)
}
