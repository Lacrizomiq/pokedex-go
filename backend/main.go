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

	r.GET("/api/pokemons", pokemonController.GetAllPokemon)
	r.GET("/api/pokemons/:id", pokemonController.GetPokemon)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error running server:", err)
	}

}
