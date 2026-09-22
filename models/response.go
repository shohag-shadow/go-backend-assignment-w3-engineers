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
type ResponseCategoryItem struct {
	LocationID string
	Name       string
	Type       string
	Slug       string
	Display    []string
}
type GeoInfo struct {
	Breadcrumbs []ResponseCategoryItem
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

type ErrorResponse struct {
	Error string
}
type SuccessResult struct {
	Count int
	Items []Response
}
type SuccessResponse struct {
	Result SuccessResult
}
