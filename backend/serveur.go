package backend

import (
	"log"
	"net/http"
)

func Run() {
	serveur := &http.Server{
		Addr:    ":8080",
		Handler: Routes(),
	}

	log.Println("http://localhost:8080")

	log.Fatal(http.ListenAndServe(serveur.Addr, serveur.Handler))
}
