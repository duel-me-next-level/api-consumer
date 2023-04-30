package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	/* tournamentRepository := repository.NewTournamentRepository()
	tournamentService := gamesrv.New(tournamentRepository) */
	/* gamesHandler := gamehdl.NewHTTPHandler(gamesService) */

	router := gin.New()
	/* router.GET("/tournament/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create) */

	router.Run(":8080")
}
