package controllers

import (
	"fmt"
	"rental-property-api/models"
	"rental-property-api/services"
	"slices"
	"strconv"
	"strings"
)

type PropertyController struct {
	BaseController
}

func (o *PropertyController) DemoResponse() {
	data := services.GetData()

	o.Data["json"] = models.GetResponseFromSource(&data.Properties[0])
	o.ServeJSON()
}

func (c *PropertyController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		c.RespondError(400, "Please enter a valid id")
		return
	}
	prop, err := services.GetProertyByID(id)
	if err != nil {
		c.RespondError(404, err.Error())
		return
	}
	response := models.GetResponseFromSource(&prop)
	c.Data["json"] = response
	c.ServeJSON()
}

func (c *PropertyController) GetAll() {
	filters, err := c.parseFilters()
	if err != nil {
		return
	}
	data := services.GetData().Properties
	filteredData := services.FilterProperties(data, filters)
	resposeData := services.PrepareResponse(filteredData)
	c.Data["json"] = resposeData
	c.ServeJSON()
}

func (c *PropertyController) parseFilters() (models.PropertyFilters, error) {
	filters := models.PropertyFilters{}

	minPriceString := c.GetString("min_price")
	if minPriceString != "" {
		minPrice, minPriceError := strconv.ParseFloat(minPriceString, 64)
		if minPriceError != nil {
			c.RespondError(400, "min_price should be a number")
			return filters, fmt.Errorf("parsing error")
		}
		if minPrice < 0 {
			c.RespondError(400, "min_price cannot be negative")
			return filters, fmt.Errorf("parsing error")
		}
		filters.MinPrice = &minPrice
	}

	maxPriceString := c.GetString("max_price")
	if maxPriceString != "" {
		maxPrice, maxPriceError := strconv.ParseFloat(maxPriceString, 64)
		if maxPriceError != nil {
			c.RespondError(400, "max_price should be a number")
			return filters, fmt.Errorf("parsing error")
		}
		if maxPrice < 0 {
			c.RespondError(400, "max_price cannot be negative")
			return filters, fmt.Errorf("parsing error")
		}
		filters.MaxPrice = &maxPrice
	}

	minStarRatingString := c.GetString("min_star_rating")
	if minStarRatingString != "" {
		minStarRating, minStarRatingError := strconv.Atoi(minStarRatingString)
		if minStarRatingError != nil {
			c.RespondError(400, "min_star_rating should be a whole number")
			return filters, fmt.Errorf("parsing error")
		}
		if minStarRating < 0 {
			c.RespondError(400, "min_star_rating cannot be negative")
			return filters, fmt.Errorf("parsing error")
		}
		filters.MinStarRating = &minStarRating
	}

	minReviewScoreString := c.GetString("min_review_score")
	if minReviewScoreString != "" {
		minReviewScore, minReviewScoreError := strconv.ParseFloat(minReviewScoreString, 64)
		if minReviewScoreError != nil {
			c.RespondError(400, "min_review_score should be a number")
			return filters, fmt.Errorf("parsing error")
		}
		if minReviewScore < 0 {
			c.RespondError(400, "min_review_score cannot be negative")
			return filters, fmt.Errorf("parsing error")
		}
		filters.MinReviewScore = &minReviewScore
	}

	minReviewsString := c.GetString("min_reviews")
	if minReviewsString != "" {
		minReviews, minReviewsError := strconv.Atoi(minReviewsString)
		if minReviewsError != nil {
			c.RespondError(400, "min_reviews should be a whole number")
			return filters, fmt.Errorf("parsing error")
		}
		if minReviews < 0 {
			c.RespondError(400, "min_reviews cannot be negative")
			return filters, fmt.Errorf("parsing error")
		}
		filters.MinReviews = &minReviews
	}

	publishedString := c.GetString("published")
	if publishedString != "" {
		published, publishedError := strconv.ParseBool(publishedString)
		if publishedError != nil {
			c.RespondError(400, "published should be either true or false")
			return filters, fmt.Errorf("parsing error")
		}
		filters.Published = &published
	}

	propertyTypeCategory := c.GetString("property_type")
	if propertyTypeCategory != "" {
		validTypes := []string{"Hotel", "House", "Apartment", "Villa", "Resort", "Hostel"}
		if !slices.Contains(validTypes, propertyTypeCategory) {
			c.RespondError(400, "property_type should be one of Hotel, House, Apartment, Villa, Resort, Hostel")
			return filters, fmt.Errorf("parsing error")
		}
		filters.PropertyType = propertyTypeCategory
	}

	feedString := c.GetString("feed")
	if feedString != "" {
		feed, feedError := strconv.Atoi(feedString)
		if feedError != nil {
			c.RespondError(400, "feed should be a whole number")
			return filters, fmt.Errorf("parsing error")
		}
		validFeeds := []int{11, 12, 22, 24}
		if !slices.Contains(validFeeds, feed) {
			c.RespondError(400, "feed should be one of 11, 12, 22, 24")
			return filters, fmt.Errorf("parsing error")
		}
		filters.Feed = &feed
	}

	minBedroomString := c.GetString("min_bedroom")
	if minBedroomString != "" {
		minBedroom, minBedroomError := strconv.Atoi(minBedroomString)
		if minBedroomError != nil {
			c.RespondError(400, "min_bedroom should be a whole number")
			return filters, fmt.Errorf("parsing error")
		}
		if minBedroom <= 0 {
			c.RespondError(400, "min_bedroom should be at least 1")
			return filters, fmt.Errorf("parsing error")
		}
		filters.MinBedroom = &minBedroom
	}

	limitString := c.GetString("limit")
	if limitString != "" {
		limit, limitError := strconv.Atoi(limitString)
		if limitError != nil {
			c.RespondError(400, "limit should be a whole number")
			return filters, fmt.Errorf("parsing error")
		}
		if limit <= 0 {
			c.RespondError(400, "limit should be at least 1")
			return filters, fmt.Errorf("parsing error")
		}
		filters.Limit = &limit
	}
	amenities := c.GetString("amenities")
	if amenities != "" {
		values := strings.Split(amenities, ",")
		filters.Amenities = values
	}
	return filters, nil
}
