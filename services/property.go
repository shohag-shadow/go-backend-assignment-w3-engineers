package services

import (
	"fmt"
	"rental-property-api/models"
	"slices"
	"strings"
	"sync"
)

var (
	instance *models.Data
	once     sync.Once
)

func GetData() *models.Data {
	once.Do(func() {
		instance = &models.Data{}
		instance.LoadData()
	})
	return instance
}
func GetProertyByID(id string) (p models.SourceProperty, e error) {
	properties := GetData().Properties
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

func FilterProperties(properties []models.SourceProperty, filter models.PropertyFilters) []models.SourceProperty {
	out := make([]models.SourceProperty, 0)

	for _, property := range properties {
		if filter.MinPrice != nil && property.USDPrice < *filter.MinPrice {
			continue
		}
		if filter.MaxPrice != nil && property.USDPrice > *filter.MaxPrice {
			continue
		}
		if filter.MinStarRating != nil && property.StarRating < *filter.MinStarRating {
			continue
		}
		if filter.MinReviewScore != nil && property.ReviewScoreGeneral < *filter.MinReviewScore {
			continue
		}
		if filter.MinReviews != nil && property.NumberOfReview < *filter.MinReviews {
			continue
		}
		if filter.Published != nil && property.Published != *filter.Published {
			continue
		}
		if filter.PropertyType != "" && property.PropertyTypeCategory != filter.PropertyType {
			continue
		}
		if filter.Feed != nil && property.Feed != *filter.Feed {
			continue
		}
		if filter.MinBedroom != nil && property.BedroomCount < *filter.MinBedroom {
			continue
		}
		if filter.Amenities != nil {
			if !(slices.ContainsFunc(property.AmenityCategories, func(amenity string) bool {
				return slices.Contains(filter.Amenities, amenity)
			})) {
				continue
			}
		}
		out = append(out, property)

		if filter.Limit != nil && *filter.Limit > 0 && len(out) >= *filter.Limit {
			break
		}
	}
	return out
}

func PrepareResponse(s []models.SourceProperty) models.SuccessResponse {
	r := []models.Response{}
	for _, property := range s {
		r = append(r, models.GetResponseFromSource(&property))
	}
	successResult := models.SuccessResult{
		Count: len(r),
		Items: r,
	}
	return models.SuccessResponse{
		Result: successResult,
	}
}
