package services

import (
	"fmt"
	"path/filepath"
	"reflect"
	"rental-property-api/models"
	"runtime"
	"slices"
	"testing"

	"github.com/beego/beego/v2/server/web"
)

var testFilePath string

func setfilepath() {
	_, thisFile, _, _ := runtime.Caller(0)
	projectRoot, err := filepath.Abs(filepath.Join(filepath.Dir(thisFile), ".."))
	if err != nil {
		panic(err)
	}

	confPath := filepath.Join(projectRoot, "conf", "app.conf")
	if err := web.LoadAppConfig("ini", confPath); err != nil {
		return
	}

	testFilePath, _ = web.AppConfig.String("testfilepath")

}
func TestGetData(t *testing.T) {
	setfilepath()
	tests := []struct {
		name string
		want int
	}{
		{"Test GetData", 100},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetData(testFilePath).Properties
			if len(got) != tc.want {
				t.Errorf("got %v values,\n want %v values", len(got), tc.want)
			}
		})
	}
}
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
	GetData(testFilePath)
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
				t.Errorf("got %+v,\n want %+v", got, tc.want)
			}
		})
	}
}

func f64(v float64) *float64 { return &v }
func i(v int) *int           { return &v }

func TestFilterProperties(t *testing.T) {
	tests := []struct {
		name    string
		filters models.PropertyFilters
		want    int
	}{
		{
			name: "Filter with and including limit",
			filters: models.PropertyFilters{
				MinPrice:     f64(100),
				PropertyType: "Resort",
				Limit:        i(2),
			},
			want: 2,
		},
		{
			name: "Filter with and without limit",
			filters: models.PropertyFilters{
				MinPrice:     f64(100),
				PropertyType: "Resort",
			},
			want: 15,
		},
		{
			name: "Filter for no value",
			filters: models.PropertyFilters{
				MinPrice:     f64(10000),
				PropertyType: "Resort",
			},
			want: 0,
		},
	}

	GetData(testFilePath)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := len(FilterProperties(tc.filters))

			if got != tc.want {
				t.Errorf("got %v values, want %v", got, tc.want)
			}
		})
	}
}
func TestFilterPropertiesForAmenities(t *testing.T) {
	tests := []struct {
		name    string
		filters models.PropertyFilters
	}{
		{
			name: "Filter with and including limit",
			filters: models.PropertyFilters{
				Amenities: []string{"Internet", "Parking"},
			},
		},
	}

	GetData(testFilePath)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := FilterProperties(tc.filters)
			doesSatisfyAmenities(tc.filters.Amenities, got, t)
		})
	}
}
func doesSatisfyAmenities(amenities []string, properties []models.SourceProperty, t *testing.T) {
	for _, property := range properties {
		found := false
		for _, gotAmenity := range property.AmenityCategories {
			if slices.Contains(amenities, gotAmenity) {
				found = true
				break
			}
			if found {
				break
			}
		}
		if found == false {
			t.Errorf("Amenities does not match got %v wanted from %v", property.AmenityCategories, amenities)
			return
		}
	}
}
func TestFilterCombined(t *testing.T) {
	tests := []struct {
		name    string
		filters models.PropertyFilters
		want    int
	}{
		{
			name: "Filter with and including limit",
			filters: models.PropertyFilters{
				MinPrice:     f64(100),
				PropertyType: "Resort",
				Amenities:    []string{"Internet", "Parking"},
			},
			want: 9,
		},
	}

	GetData(testFilePath)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := FilterProperties(tc.filters)
			doesSatisfyAmenities(tc.filters.Amenities, got, t)
			if len(got) != tc.want {
				t.Errorf("got %v values, want %v", len(got), tc.want)
			}
		})
	}
}
func TestGetResponseFromSource(t *testing.T) {
	tests := []struct {
		name     string
		source   models.SourceProperty
		response models.Response
	}{
		{"Test sourse with response", wantHa2000022, responseHA2000022},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetResponseFromSource(&tc.source)
			if !reflect.DeepEqual(got, tc.response) {
				t.Errorf("got %+v,\n want %+v", got, tc.response)
			}
		})
	}
}

func TestPrepareResponse(t *testing.T) {

	tests := []struct {
		name   string
		source []models.SourceProperty
		want   int
	}{
		{"Test Response structure", GetData(testFilePath).Properties, 100},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := PrepareResponse(tc.source)
			if len(got.Result.Items) != got.Result.Count || got.Result.Count != tc.want {
				t.Errorf("got %v values,\n want %v values", got.Result.Count, tc.want)
			}
		})
	}
}

var responseHA2000022 = models.Response{
	ID:   "HA-2000022",
	Feed: 12,
	GeoInfo: models.GeoInfo{
		Breadcrumbs: []models.ResponseCategoryItem{
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
		City:        "Salt Lake City",
		Country:     "United States",
		CountryCode: "US",
		Name:        "Salt Lake City, United States",
		LocationID:  "7200022",
		Lat:         40.7608,
		Lon:         -111.891,
		State:       "Utah",
		StateAbbr:   "UT",
	},
	Property: models.Property{
		Amenities:    []string{"Child Friendly", "Laundry", "Parking", "Gym", "Internet"},
		Name:         "Salt Lake City Harbor House",
		Slug:         "salt-lake-city-harbor-house-0022",
		PropertyType: "House",
		Price:        91.5,
		ReviewScore:  8.6,
		StarRating:   3,
		Counts: models.Counts{
			Bathroom:  2,
			Bedroom:   1,
			Reviews:   275,
			Occupancy: 7,
		},
		Image: models.Image{
			Count:  4,
			Images: []string{"image-1.jpg", "image-2.jpg", "image-3.jpg", "image-4.jpg"},
		},
	},
	Published: false,
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
