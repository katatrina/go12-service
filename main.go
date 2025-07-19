package main

import (
	"fmt"
	"log"
	"os"
	
	"github.com/gin-gonic/gin"
	"github.com/katatrina/go12-service/middleware"
	categorymodule "github.com/katatrina/go12-service/modules/category"
	mediamodule "github.com/katatrina/go12-service/modules/media"
	restaurantmodule "github.com/katatrina/go12-service/modules/restaurant"
	usermodule "github.com/katatrina/go12-service/modules/user"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	dsn := os.Getenv("DB_DSN")
	dbMaster, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to init db session", err)
	}
	
	sqlDB, err := dbMaster.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB from gorm session", err)
	}
	
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("failed to connect to database", err)
	}
	
	db := dbMaster.Debug()
	
	fmt.Println("Connected to database", db)
	
	r := gin.Default()
	gin.ForceConsoleColor()
	
	// CRUDL - Create Read Update Delete List
	// Version Prefix: /v1
	
	r.Use(middleware.RecoverMiddleware())
	
	v1 := r.Group("/v1")
	
	appCtx := sharedinfras.NewAppContext(db)
	{
		categorymodule.SetupCategoryModule(db, v1)
		restaurantmodule.SetupRestaurantModule(appCtx, v1)
		usermodule.SetupUserModule(appCtx, v1)
		mediamodule.SetupMediaModule(appCtx, v1)
	}
	
	r.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
