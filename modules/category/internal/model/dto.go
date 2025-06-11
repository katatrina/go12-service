package categorymodel

type UpdateCategoryDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

type FilterCategoryDTO struct {
	Status *string `json:"status" form:"status"`
}
