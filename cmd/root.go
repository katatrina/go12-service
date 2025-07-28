package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	
	"github.com/katatrina/go12-service/middleware"
	categorymodule "github.com/katatrina/go12-service/modules/category"
	mediamodule "github.com/katatrina/go12-service/modules/media"
	restaurantmodule "github.com/katatrina/go12-service/modules/restaurant"
	restaurantlikemodule "github.com/katatrina/go12-service/modules/restaurantlike"
	usermodule "github.com/katatrina/go12-service/modules/user"
	sharedinfras "github.com/katatrina/go12-service/shared/infras"
	
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start rest api service",
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("PORT")
		
		if port == "" {
			port = "8080"
		}
		
		dsn := os.Getenv("DB_DSN")
		dbMaster, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		
		if err != nil {
			log.Fatal("failed to connect database", err)
		}
		
		db := dbMaster.Debug()
		
		fmt.Println("Connected to database", db)
		
		r := gin.Default()
		
		r.Use(middleware.RecoverMiddleware())
		
		r.GET("/ping", func(c *gin.Context) {
			
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
			
		})
		
		r.Static("/uploads", "./uploads")
		
		// CRUDL - Create Read Update Delete List
		// Version Prefix: /v1
		
		v1 := r.Group("/v1")
		
		appCtx := sharedinfras.NewAppContext(db)
		
		categorymodule.SetupCategoryModule(db, v1)
		restaurantmodule.SetupRestaurantModule(appCtx, v1)
		restaurantlikemodule.SetupRestaurantLikeModule(appCtx, v1)
		usermodule.SetupUserModule(appCtx, v1)
		mediamodule.SetupMediaModule(appCtx, v1)
		
		r.Run(fmt.Sprintf(":%s", port))
	},
}

func Execute() {
	setupConsumerCmd()
	
	rootCmd.AddCommand(consumerCmd)
	
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("failed to execute command", err)
	}
}
