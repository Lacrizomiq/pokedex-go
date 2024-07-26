package controllers

import (
	"net/http"

	"github.com/Lacrizomiq/pokedex-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TypeController struct {
	DB *gorm.DB
}

func NewTypeController(db *gorm.DB) *TypeController {
	return &TypeController{DB: db}
}

func (tc *TypeController) GetAllTypes(c *gin.Context) {
	var types []models.Type
	result := tc.DB.Find(&types)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, types)
}

func (tc *TypeController) GetType(c *gin.Context) {
	id := c.Param("id")
	var type_ models.Type
	result := tc.DB.First(&type_, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Type not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, type_)
}
