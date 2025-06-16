//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package categorymodule

import (
	"github.com/google/wire"
	categoryhttpgin "github.com/katatrina/go12-service/modules/category/infras/controller/http-gin"
	"gorm.io/gorm"
)

func InitializeCategoryController(db *gorm.DB) *categoryhttpgin.CategoryHTTPController {
	wire.Build(
		CategorySet,
	)
	
	return &categoryhttpgin.CategoryHTTPController{}
}
