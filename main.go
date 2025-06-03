package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
	
	"github.com/gin-gonic/gin"
)

type Category struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      int        `json:"status"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

func (c *Category) Validate() error {
	c.Name = strings.TrimSpace(c.Name)
	
	if c.Name == "" {
		return errors.New("name is required")
	}
	
	if c.Status <= 0 {
		return errors.New("status must be greater than zero")
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
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
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
	categories := make([]Category, 0)
	// var categories []Category // categories = nil
	latestId := 0
	
	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context) {
		
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		
	})
	
	// CRUDL - Create Read Update Delete List
	// Version Prefix: /v1
	
	group := r.Group("/v1")
	// API Create
	
	group.POST("/categories", func(c *gin.Context) {
		var requestBodyData Category
		
		if err := c.ShouldBindJSON(&requestBodyData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		if err := requestBodyData.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		latestId++
		requestBodyData.Id = latestId
		
		now := time.Now().UTC()
		
		requestBodyData.CreatedAt = &now
		requestBodyData.UpdatedAt = &now
		
		categories = append(categories, requestBodyData)
		
		c.JSON(http.StatusCreated, gin.H{"data": requestBodyData.Id})
	})
	
	// Listing
	group.GET("/categories", func(c *gin.Context) {
		var paging PagingDTO
		
		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		paging.Process()
		
		c.JSON(http.StatusOK, gin.H{"data": categories, "paging": paging})
	})
	
	// Get
	group.GET("/categories/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		for _, category := range categories {
			if category.Id == id {
				c.JSON(http.StatusOK, gin.H{"data": category})
				return
			}
		}
		
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
	})
	
	// Delete
	group.DELETE("/categories/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		for i, category := range categories {
			if category.Id == id {
				
				// Explain:
				// [1,2,3], i = 1
				// categories[:i] = [1]
				// categories[i+1:] = [3]
				// [1,3]
				categories = append(categories[:i], categories[i+1:]...)
				
				c.JSON(http.StatusOK, gin.H{"message": true})
				return
			}
		}
		
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
	})
	
	// Update
	group.PATCH("/categories/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		var requestBodyData CategoryUpdateDTO
		
		if err := c.ShouldBindJSON(&requestBodyData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		for i, category := range categories {
			if category.Id == id {
				if requestBodyData.Name != nil {
					categories[i].Name = *requestBodyData.Name
				}
				
				if requestBodyData.Description != nil {
					categories[i].Description = *requestBodyData.Description
				}
				
				if requestBodyData.Status != nil {
					categories[i].Status = *requestBodyData.Status
				}
				
				now := time.Now().UTC()
				categories[i].UpdatedAt = &now
				
				c.JSON(http.StatusOK, gin.H{"data": true})
				return
			}
		}
		
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
	})
	
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
