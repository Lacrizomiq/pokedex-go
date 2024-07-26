package controllers

import (
	"net/http"

	"github.com/Lacrizomiq/pokedex-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PokemonController struct {
	DB *gorm.DB
}

func NewPokemonController(db *gorm.DB) *PokemonController {
	return &PokemonController{DB: db}
}

func (pc *PokemonController) GetAllPokemon(c *gin.Context) {
	var pokemons []models.Pokemon
	result := pc.DB.Preload("Types").Find(&pokemons)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, pokemons)
}

func (pc *PokemonController) GetPokemon(c *gin.Context) {
	id := c.Param("id")
	var pokemon models.Pokemon
	result := pc.DB.Preload("Types").First(&pokemon, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, pokemon)
}
