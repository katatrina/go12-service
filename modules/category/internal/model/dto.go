package categorymodel

type CategoryUpdateDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *int    `json:"status"`
}

type FilterCategoryDTO struct {
	Status *string `json:"status" form:"status"`
}
