package categorymodel

type CreateCategoryDTO struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type UpdateCategoryDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

type FilterCategoryDTO struct {
	Status *string `json:"status" form:"status"`
}
