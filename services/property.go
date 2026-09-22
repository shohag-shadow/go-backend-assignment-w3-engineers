package services

import (
	"fmt"
	"rental-property-api/models"
	"slices"
	"strings"
)

func GetProertyById(id string) (p models.SourceProperty, e error) {
	properties := models.GetData().Properties
	sorted := slices.IsSortedFunc(properties, func(a, b models.SourceProperty) int {
		return strings.Compare(a.Id, b.Id)
	})
	fmt.Println("Is sorted by Id:", sorted)
	p = properties[0]
	e = nil
	return p, e
}
