package models

type Image struct {
	Count  int
	Images []string
}
type Counts struct {
	Bathroom  int
	Bedroom   int
	Reviews   int
	Occupancy int
}

type Property struct {
	Amenities    []string
	Name         string
	Slug         string
	PropertyType string
	Price        float64
	ReviewScore  float64
	StarRating   int
	Counts       Counts
	Image        Image
}

type GeoInfo struct {
	Breadcrumbs []CategoryItem
	City        string
	Country     string
	CountryCode string
	Name        string
	LocationID  string
	Lat         float64
	Lon         float64
	State       string
	StateAbbr   string
}

type Response struct {
	ID        string
	Feed      int
	GeoInfo   GeoInfo
	Property  Property
	Published bool
}

func GetResponseFromSource(s *SourceProperty) (r Response) {

	geoInfo := GeoInfo{
		Breadcrumbs: s.Categories.Items,
		City:        s.City,
		Country:     s.Country,
		CountryCode: s.CountryCode,
		Name:        s.Display,
		Lat:         s.LonLat.Coordinates[1],
		Lon:         s.LonLat.Coordinates[0],
		State:       s.State,
		StateAbbr:   s.StateAbbr,
	}
	counts := Counts{
		Bathroom:  s.BathroomCount,
		Bedroom:   s.BedroomCount,
		Reviews:   s.NumberOfReview,
		Occupancy: s.Occupancy,
	}
	image := Image{
		Count:  len(s.Images),
		Images: s.Images,
	}
	property := Property{
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
