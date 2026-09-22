package services

import (
	"fmt"
	"rental-property-api/models"
	"slices"
	"strings"
)

func GetProertyByID(id string) (p models.SourceProperty, e error) {
	properties := models.GetData().Properties
	index, found := slices.BinarySearchFunc(properties, id, func(property models.SourceProperty, id string) int {
		return strings.Compare(property.Id, id)
	})
	if !found {
		e = fmt.Errorf("Id not found")
		return p, e
	}
	p = properties[index]
	e = nil
	return p, e
}
