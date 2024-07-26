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

	typeController := controllers.NewTypeController(db)
	r.GET("/api/types", typeController.GetAllTypes)
	r.GET("/api/types/:id", typeController.GetType)

	teamController := controllers.NewTeamController(db)
	r.GET("/api/teams", teamController.GetAllTeams)
	r.GET("/api/teams/:id", teamController.GetTeam)
	r.POST("/api/teams", teamController.CreateTeam)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error running server:", err)
	}
}
