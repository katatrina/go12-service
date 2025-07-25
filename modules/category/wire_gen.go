// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package categorymodule

import (
	"github.com/katatrina/go12-service/modules/category/infras/controller/http"
	"github.com/katatrina/go12-service/modules/category/infras/repository/mysql"
	"github.com/katatrina/go12-service/modules/category/service"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeCategoryController(db *gorm.DB) *httpcontroller.CategoryController {
	categoryRepository := mysqlrepository.NewCategoryRepository(db)
	createCommandHandler := categoryservice.NewCreateCommandHandler(categoryRepository)
	getByIDQueryHandler := categoryservice.NewGetDetailQueryHandler(categoryRepository)
	listCategoriesQueryHandler := categoryservice.NewListCategoriesQueryHandler(categoryRepository)
	updateByIDCommandHandler := categoryservice.NewUpdateByIDCommandHandler(categoryRepository)
	deleteByIDCommandHandler := categoryservice.NewDeleteByIDCommandHandler(categoryRepository)
	categoryController := httpcontroller.NewCategoryController(createCommandHandler, getByIDQueryHandler, listCategoriesQueryHandler, updateByIDCommandHandler, deleteByIDCommandHandler, categoryRepository)
	return categoryController
}
