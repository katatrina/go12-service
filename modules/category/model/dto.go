package categorymodel

import (
	"strings"
	
	"github.com/katatrina/go12-service/shared/datatype"
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

func (dto *FilterCategoryDTO) Validate() error {
	if dto.Status != nil {
		*dto.Status = strings.TrimSpace(*dto.Status)
		status := datatype.Status(strings.ToLower(*dto.Status))
		if !status.Valid() {
			return ErrStatusInvalid
		}
	}
	
	return nil
}
