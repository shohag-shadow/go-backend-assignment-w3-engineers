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

func GetData(filepath string) *models.Data {
	once.Do(func() {
		instance = &models.Data{}
		instance.LoadData(filepath)
	})
	return instance
}
func GetProertyByID(id string) (p models.SourceProperty, e error) {
	properties := GetData("").Properties
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

func FilterProperties(filter models.PropertyFilters) []models.SourceProperty {
	properties := GetData("").Properties
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
func GetResponseFromSource(s *models.SourceProperty) (r models.Response) {
	breadcurmbs := []models.ResponseCategoryItem{}
	for _, item := range s.Categories.Items {
		breadcurmbs = append(breadcurmbs, models.ResponseCategoryItem{
			LocationID: item.LocationID,
			Name:       item.Name,
			Type:       item.Type,
			Slug:       item.Slug,
			Display:    item.Display,
		})
	}
	geoInfo := models.GeoInfo{
		Breadcrumbs: breadcurmbs,
		City:        s.City,
		Country:     s.Country,
		CountryCode: s.CountryCode,
		Name:        s.Display,
		Lat:         s.LonLat.Coordinates[1],
		Lon:         s.LonLat.Coordinates[0],
		State:       s.State,
		StateAbbr:   s.StateAbbr,
	}
	counts := models.Counts{
		Bathroom:  s.BathroomCount,
		Bedroom:   s.BedroomCount,
		Reviews:   s.NumberOfReview,
		Occupancy: s.Occupancy,
	}
	image := models.Image{
		Count:  len(s.Images),
		Images: s.Images,
	}
	property := models.Property{
		Amenities:    s.AmenityCategories,
		Name:         s.PropertyName,
		Slug:         s.PropertySlug,
		PropertyType: s.PropertyTypeCategory,
		Price:        s.USDPrice,
		ReviewScore:  s.ReviewScoreGeneral,
		StarRating:   s.StarRating,
		Counts:       counts,
		Image:        image,
	}
	r.ID = s.Id
	r.Feed = s.Feed
	r.GeoInfo = geoInfo
	r.Property = property
	r.Published = s.Published
	return r
}

func PrepareResponse(s []models.SourceProperty) models.SuccessResponse {
	r := []models.Response{}
	for _, property := range s {
		r = append(r, GetResponseFromSource(&property))
	}
	successResult := models.SuccessResult{
		Count: len(r),
		Items: r,
	}
	return models.SuccessResponse{
		Result: successResult,
	}
}
