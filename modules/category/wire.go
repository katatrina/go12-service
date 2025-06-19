//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package categorymodule

import (
	"github.com/google/wire"
	"github.com/katatrina/go12-service/modules/category/infras/controller/controller"
	"gorm.io/gorm"
)

func InitializeCategoryController(db *gorm.DB) *controller.CategoryHTTPController {
	wire.Build(
		CategorySet,
	)
	
	return &controller.CategoryHTTPController{}
}
