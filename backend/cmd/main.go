package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	routes "social-network/delivery/http"
	ws "social-network/delivery/ws"
	infrastructure "social-network/infrastructure/db"
	internal "social-network/internal/app"
)

var (
	srv *http.Server
	hub *ws.Hub
	db  *sql.DB
	app *internal.App
)

func init() {
	// Charger les variables d'environnement
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("❌ Erreur lors du chargement du fichier .env")
	}

	// Initialiser le hub WebSocket
	hub = ws.InitHub()

	go hub.Run()

	// Connexion à la base de données
	db, err = infrastructure.OpenDB()
	if err != nil {
		log.Fatal("❌ Erreur de connexion à la base de données :", err)
	}
	app = internal.NewApp(db)

	// Création du serveur HTTP
	srv = &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: routes.Routes(hub, app), // <-- Injection ici

	}
}

func main() {
	log.Println("🚀 Serveur lancé sur le port", srv.Addr)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// TLS("./tls/localhost.pem", "./tls/localhost-key.pem")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Erreur serveur : %v", err)
		}
	}()
	<-stop
	log.Println("Signal reçu, arrêt du serveur...")
	// Contexte avec timeout pour arrêt propre
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Shutdown du serveur HTTP
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Erreur lors du shutdown HTTP : %v", err)
	}
	infrastructure.Close(db)
	hub.Stop()
	log.Println("Serveur, base de données et WebSocket arrêtés proprement.")
}
