package backend

import "net/http"

func Run() {
	serveur := &http.Server{
		Addr:    ":8080",
		Handler: Routes(),
	}

	http.ListenAndServe(serveur.Addr, serveur.Handler)
}

