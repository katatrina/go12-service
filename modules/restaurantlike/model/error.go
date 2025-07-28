package restaurantlikemodel

import "errors"

var (
	ErrRestaurantLikeNotFound = errors.New("restaurant like not found")
	ErrRestaurantLikeExists   = errors.New("restaurant already liked")
)
