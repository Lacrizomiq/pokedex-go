package controllers

import (
	"net/http"

	"github.com/Lacrizomiq/pokedex-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeamController struct {
	DB *gorm.DB
}

func NewTeamController(db *gorm.DB) *TeamController {
	return &TeamController{DB: db}
}

func (tc *TeamController) GetAllTeams(c *gin.Context) {
	var teams []models.Team
	result := tc.DB.Preload("Pokemon.Types").Find(&teams)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

func (tc *TeamController) GetTeam(c *gin.Context) {
	id := c.Param("id")
	var team models.Team
	result := tc.DB.Preload("Pokemon.Types").First(&team, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, team)
}

func (tc *TeamController) CreateTeam(c *gin.Context) {
	var team models.Team
	err := c.ShouldBindJSON(&team)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := tc.DB.Create(&team)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, team)
}
