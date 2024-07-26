package main

import (
	"log"

	"github.com/Lacrizomiq/pokedex-go/config"
	"github.com/Lacrizomiq/pokedex-go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}

	r := gin.Default()

	pokemonController := controllers.NewPokemonController(db)

	r.GET("/pokemons", pokemonController.GetAllPokemon)
	r.GET("/pokemons/:id", pokemonController.GetPokemon)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error running server:", err)
	}

}
