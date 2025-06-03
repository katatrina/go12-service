package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Id          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

func (Category) TableName() string {
	return "categories"
}

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusDeleted  = "deleted"
)

func (c *Category) Validate() error {
	c.Name = strings.TrimSpace(c.Name)
	
	if c.Name == "" {
		return errors.New("name is required")
	}
	
	if c.Status != StatusActive && c.Status != StatusInactive && c.Status != StatusDeleted {
		return errors.New("status must be active, inactive or deleted")
	}
	
	return nil
}

// DTO = Data Transfer Object
type CategoryUpdateDTO struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Status      *int    `json:"status"`
}

type PagingDTO struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total"`
}

func (p *PagingDTO) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	
	if p.Limit <= 0 {
		p.Limit = 10
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	
	dsn := os.Getenv("DB_DSN")
	dbMaster, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		log.Fatal("failed to connect database", err)
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
		{
			categories.POST("", CreateAPI(db))
			categories.GET("", ListAPI(db))
			categories.GET("/:id", GetAPI(db))
			categories.PATCH("/:id", UpdateAPI(db))
			categories.DELETE("/:id", DeleteAPI(db))
		}
	}
	
	r.Run(fmt.Sprintf(":%s", port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
