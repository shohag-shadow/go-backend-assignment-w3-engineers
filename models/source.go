package models

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/beego/beego/v2/core/logs"
)

type CategoryItem struct {
	LocationID string   `json:"LocationID"`
	Name       string   `json:"Name"`
	Type       string   `json:"Type"`
	Slug       string   `json:"Slug"`
	Display    []string `json:"Display"`
}

type CategoryList struct {
	Items []CategoryItem `json:"categories"`
}

// custom unmershall as we need to unmarshall twice to get categorytem from the string.
func (c *CategoryList) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw == "" {
		c.Items = nil
		return nil
	}

	var items []CategoryItem
	if err := json.Unmarshal([]byte(raw), &items); err != nil {
		return err
	}
	c.Items = items
	return nil
}

type LongitudeLatitude struct {
	Coordinates [2]float64 `json:"coordinates"`
}

type Property struct {
	Id                   string            `json:"id"`
	Feed                 int               `json:"feed"`
	Country              string            `json:"country"`
	CountryCode          string            `json:"country_code"`
	State                string            `json:"state"`
	StateAbbr            string            `json:"state_abbr"`
	City                 string            `json:"city"`
	Display              string            `json:"display"`
	LocationID           string            `json:"location_id"`
	PropertyName         string            `json:"property_name"`
	PropertySlug         string            `json:"property_slug"`
	PropertyTypeCategory string            `json:"property_type_category"`
	USDPrice             float64           `json:"usd_price"`
	Occupancy            int               `json:"occupancy"`
	BedroomCount         int               `json:"bedroom_count"`
	BathroomCount        int               `json:"bathroom_count"`
	NumberOfReview       int               `json:"number_of_review"`
	ReviewScoreGeneral   float64           `json:"review_score_general"`
	StarRating           int               `json:"star_rating"`
	AmenityCategories    []string          `json:"amenity_categories"`
	LonLat               LongitudeLatitude `json:"lonlat"`
	Categories           CategoryList      `json:"categories"`
	Published            bool              `json:"published"`
	Images               []string          `json:"images"`
}
type Data struct {
	Properties []Property
}

var (
	instance *Data
	once     sync.Once
)

func GetData() *Data {
	once.Do(func() {
		instance = &Data{}
		instance.loadData()
	})
	return instance
}
func (d *Data) loadData() {
	logs.Informational("Reading data from disk")
	data, err := os.ReadFile("data/rental_properties.json")
	if err != nil {
		logs.Error("Error reading file:", err)
		return
	}
	var properties []Property
	err = json.Unmarshal(data, &properties)
	if err != nil {
		logs.Error("Error parsing JSON:", err)
		return
	}
	d.Properties = properties
}
