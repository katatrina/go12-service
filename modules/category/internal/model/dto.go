package model

import (
	"strings"
)

type CreateCategoryDTO struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type UpdateCategoryDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type FilterCategoryDTO struct {
	Status *string `json:"status" form:"status"`
}

func (dto *UpdateCategoryDTO) Validate() error {
	if dto.Name != nil {
		*dto.Name = strings.TrimSpace(*dto.Name)
		if *dto.Name == "" {
			return ErrNameRequired
		}

		if len(*dto.Name) > 100 {
			return ErrInvalidNameLength
		}
	}

	if dto.Description != nil {
		*dto.Description = strings.TrimSpace(*dto.Description)
	}

	// TODO: Add more validation rules if needed

	return nil
}
