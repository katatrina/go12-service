package foodmodel

import (
	"errors"
)

var (
	ErrFoodNotFound       = errors.New("food not found")
	ErrNameRequired       = errors.New("food name is required")
	ErrPriceInvalid       = errors.New("food price must be greater than 0")
	ErrRestaurantRequired = errors.New("restaurant_id is required")
	ErrFoodAlreadyDeleted = errors.New("food is already deleted")
	ErrInvalidPriceRange  = errors.New("min_price cannot be greater than max_price")
	// Image-related errors - TODO: Will be added later
	// ErrInvalidMediaID        = errors.New("media_id cannot be empty")
	// ErrMultipleDefaultImages = errors.New("only one default image is allowed")
)