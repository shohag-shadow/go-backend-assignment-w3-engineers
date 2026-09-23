package services

import (
	"fmt"
	"reflect"
	"rental-property-api/models"
	"testing"
)

func TestGetPropertyByID(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		want    models.SourceProperty
		wantErr error
	}{
		{
			name:    "Valid id",
			id:      "HA-2000022",
			want:    wantHa2000022,
			wantErr: nil,
		},
		{
			name:    "Invalid id",
			id:      "1000000",
			want:    models.SourceProperty{},
			wantErr: fmt.Errorf("Id not found"),
		},
		{
			name:    "Empty string",
			id:      "  ",
			want:    models.SourceProperty{},
			wantErr: fmt.Errorf("Id not found"),
		},
		{
			name:    "String with unicode character",
			id:      "😃",
			want:    models.SourceProperty{},
			wantErr: fmt.Errorf("Id not found"),
		},
	}
	//calling getdata so that data file loads before the test starts
	GetData("../data/rental_properties.json")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetProertyByID(tc.id)

			if tc.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tc.wantErr)
				}
				if err.Error() != tc.wantErr.Error() {
					t.Fatalf("expected error %v, got %v", tc.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

var wantHa2000022 = models.SourceProperty{
	Id:                   "HA-2000022",
	Feed:                 12,
	Country:              "United States",
	CountryCode:          "US",
	State:                "Utah",
	StateAbbr:            "UT",
	City:                 "Salt Lake City",
	Display:              "Salt Lake City, United States",
	LocationID:           "7200022",
	PropertyName:         "Salt Lake City Harbor House",
	PropertySlug:         "salt-lake-city-harbor-house-0022",
	PropertyTypeCategory: "House",
	USDPrice:             91.5,
	Occupancy:            7,
	BedroomCount:         1,
	BathroomCount:        2,
	NumberOfReview:       275,
	ReviewScoreGeneral:   8.6,
	StarRating:           3,
	AmenityCategories: []string{
		"Child Friendly",
		"Laundry",
		"Parking",
		"Gym",
		"Internet",
	},
	LonLat: models.LongitudeLatitude{
		Coordinates: [2]float64{-111.891, 40.7608},
	},
	Categories: models.CategoryList{
		Items: []models.CategoryItem{
			{
				LocationID: "90",
				Name:       "United States",
				Type:       "country",
				Slug:       "united-states",
				Display:    []string{"united-states"},
			},
			{
				LocationID: "7050020",
				Name:       "Utah",
				Type:       "state",
				Slug:       "united-states/utah",
				Display:    []string{"united-states", "utah"},
			},
			{
				LocationID: "7200022",
				Name:       "Salt Lake City",
				Type:       "city",
				Slug:       "united-states/utah/salt-lake-city",
				Display:    []string{"united-states", "utah", "salt-lake-city"},
			},
		},
	},
	Published: false,
	Images: []string{
		"image-1.jpg",
		"image-2.jpg",
		"image-3.jpg",
		"image-4.jpg",
	},
}
