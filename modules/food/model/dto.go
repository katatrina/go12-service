package foodmodel

type CreateFoodDTO struct {
	RestaurantID string  `json:"restaurant_id"`
	CategoryID   *string `json:"category_id"`
	Name         string  `json:"name"`
	Description  *string `json:"description"`
	Price        float64 `json:"price"`
	// Images       []FoodImageReference   `json:"images"` // TODO: Will be added later
}

func (dto *CreateFoodDTO) Validate() error {
	if dto.Name == "" {
		return ErrNameRequired
	}
	
	if dto.RestaurantID == "" {
		return ErrRestaurantRequired
	}
	
	if dto.Price <= 0 {
		return ErrPriceInvalid
	}
	
	return nil
}

type UpdateFoodDTO struct {
	CategoryID  *string  `json:"category_id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	// Images      *[]FoodImageReference `json:"images"` // TODO: Will be added later
}

func (dto *UpdateFoodDTO) Validate() error {
	if dto.Price != nil && *dto.Price <= 0 {
		return ErrPriceInvalid
	}
	
	return nil
}

type FoodListDTO struct {
	RestaurantID *string  `json:"restaurant_id,omitempty" form:"restaurant_id"`
	CategoryID   *string  `json:"category_id,omitempty" form:"category_id"`
	MinPrice     *float64 `json:"min_price,omitempty" form:"min_price"`
	MaxPrice     *float64 `json:"max_price,omitempty" form:"max_price"`
	Search       *string  `json:"search,omitempty" form:"search"`
	Page         int      `json:"page" form:"page"`
	Limit        int      `json:"limit" form:"limit"`
}

func (dto *FoodListDTO) Validate() error {
	if dto.Page <= 0 {
		dto.Page = 1
	}
	
	if dto.Limit <= 0 || dto.Limit > 100 {
		dto.Limit = 20
	}
	
	if dto.MinPrice != nil && dto.MaxPrice != nil {
		if *dto.MinPrice > *dto.MaxPrice {
			return ErrInvalidPriceRange
		}
	}
	
	return nil
}

// CategoryInfo represents category information from gRPC call
type CategoryInfo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// RestaurantInfo represents restaurant information from HTTP/gRPC call
type RestaurantInfo struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Address    string  `json:"address"`
	CategoryID *string `json:"category_id"`
	Status     string  `json:"status"`
}

type FoodResponseDTO struct {
	*Food
	Category   *CategoryInfo   `json:"category,omitempty"`
	Restaurant *RestaurantInfo `json:"restaurant,omitempty"`
}

func NewFoodResponseDTO(food *Food) *FoodResponseDTO {
	return &FoodResponseDTO{Food: food}
}

func (dto *FoodResponseDTO) WithCategory(category *CategoryInfo) *FoodResponseDTO {
	dto.Category = category
	return dto
}

func (dto *FoodResponseDTO) WithRestaurant(restaurant *RestaurantInfo) *FoodResponseDTO {
	dto.Restaurant = restaurant
	return dto
}

type FoodListResponseDTO struct {
	Data  []*FoodResponseDTO `json:"data"`
	Page  int                `json:"page"`
	Limit int                `json:"limit"`
	Total int64              `json:"total"`
}