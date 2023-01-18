package main

import (
	"log"

	"github.com/duel-me-next-level/api-consumer/internal/app/web"
	"github.com/duel-me-next-level/api-consumer/internal/infra/server"
)

func main() {
	//Inicializar o servidor
	srv := server.New()

	//Inicializar o router
	router := web.NewRouter()

	//Inicializar handlers
	web.SetupHandlers(router)

	//Iniciar servidor HTTP
	log.Fatal(srv.ListenAndServe(":8080", router))
}
