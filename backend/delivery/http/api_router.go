package http

import (
	"context"
	"log"
	"net/http"

	Handlers "social-network/delivery/http/handlers"
	mid "social-network/delivery/http/middlewares"
	ws "social-network/delivery/ws"
	internal "social-network/internal/app"
)

func Routes(hub *ws.Hub, app *internal.App) http.Handler {
	MW := mid.NewAuthMiddleware(app.SessionUsecase)
	Handlers := Handlers.Handler{
		UserUsecase:    app.UserUsecase,
		SessionUsecase: app.SessionUsecase,
		PostUsecase:    app.PostUsescase,
		MessageUsecase: app.MessageUsecase,
		FriendUsecase:  app.FriendUsecase,
		NotificationUsecase: app.NotificationUsecase,
		Hub:            hub,
	}
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../internal/assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /api/messages/", Handlers.HandleMessage) 

	mux.HandleFunc("POST /api/register", Handlers.HandleRegister)
	mux.HandleFunc("POST /api/signin", Handlers.HandleLogin)

	mux.HandleFunc("/api/friends", MW.RequireAuth(Handlers.HandleFriends))
	mux.HandleFunc("/api/notifications", MW.RequireAuth(Handlers.HandleNotifications))
	mux.HandleFunc("POST /api/createpost", MW.RequireAuth(Handlers.HandleCreatePost))
	mux.HandleFunc("GET /api/posts", MW.RequireAuth(Handlers.HandleGetPosts))
	mux.HandleFunc("GET /api/profile", MW.RequireAuth(Handlers.HandleProfil))
	mux.HandleFunc("GET /api/logout", MW.RequireAuth(Handlers.HandleLogout))
	mux.HandleFunc("GET /api/session", Handlers.HandleSession)
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("Token")
		if err != nil {
			log.Printf("❌ Cookie Token introuvable: %v", err)
			http.Error(w, "Token cookie required", http.StatusUnauthorized)
			return
		}
		UserId := func(ctx context.Context, token string) string {
			return Handlers.SessionUsecase.GetIdOnly(ctx, token)
		}
		ws.ServeWs(hub, w, r, UserId,Handlers.MessageUsecase)
	})

	return mid.CORSMiddleware(mid.SecureHeaders(mux))
}
