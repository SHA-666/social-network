package handlers

import (
	"encoding/json"
	"net/http"

	"social-network/delivery/ws"
)

func (h *Handler) HandleSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Token")
	if err != nil || cookie.Value == "" {
		http.Error(w, "cookie not found", http.StatusUnauthorized)
		return
	}
	id := h.SessionUsecase.GetIdOnly(r.Context(), cookie.Value)
	Client := &ws.Client{UserId: id}
	err = cookie.Valid()
	if err != nil {
		h.SessionUsecase.DeleteSession(r.Context(), cookie.Value)
		h.Hub.Unregister <- Client
		http.Error(w, "cookie not valid", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"status": "connected",
		"id":     id,
		"token":  cookie.Value,
	})
}
func (h *Handler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Token")
	if err != nil {
		http.Error(w, "Erreur lors de la recherche de cookie", http.StatusInternalServerError)
		return
	}
	id := h.SessionUsecase.GetIdOnly(r.Context(), cookie.Value)
	h.Hub.DisconnectUser(id)

	if err := h.SessionUsecase.DeleteSession(r.Context(), cookie.Value); err != nil {
		http.Error(w, "Erreur de session", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "Token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Déconnexion réussie"})
}
