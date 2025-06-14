package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	
	"github.com/gin-gonic/gin"
	categorymodule "github.com/katatrina/go12-service/modules/categorymodule"
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
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	
	// CRUDL - Create Read Update Delete List
	// Version Prefix: /v1
	
	v1 := r.Group("/v1")
	
	{
		categories := v1.Group("/categories")
		categorymodule.SetupCategoryModule(db, categories)
	}
	
	r.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
