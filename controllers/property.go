package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"rental-property-api/models"

	beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	beego.Controller
}

func (o *PropertyController) DemoResponse() {
	data, err := os.ReadFile("data/rental_properties.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	var properties []models.Property
	err = json.Unmarshal(data, &properties)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	o.Data["json"] = properties
	// models.Response{
	// 	Status:  "succed",
	// 	Messege: "this is a demo messege",
	// }
	o.ServeJSON()
}
