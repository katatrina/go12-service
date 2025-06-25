//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

package restaurantmodule

import (
	"github.com/google/wire"
	controller "github.com/katatrina/go12-service/modules/restaurant/infras/controller/http"
	"gorm.io/gorm"
)

func InitializeRestaurantController(db *gorm.DB) *controller.RestaurantController {
	wire.Build(
		RestaurantSet,
	)
	return &controller.RestaurantController{}
}
