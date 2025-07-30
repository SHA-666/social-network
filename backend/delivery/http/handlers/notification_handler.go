package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) HandleNotifications(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Token")
	if err != nil || c.Value == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Token cookie required"})
		return
	}
	// Vérifier l'utilisateur
	userID := h.SessionUsecase.GetIdOnly(r.Context(), c.Value)
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user"})
		return
	}
	switch r.Method {
	case "GET":
		notif, err := h.NotificationUsecase.GetAllNotif(r.Context(), userID)
		if err != nil {
			fmt.Println("error to check 'get notificztoun<<<<<<<<<<<<", err)
			return
		}
		response := map[string]any{
			"notification": notif,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	case "POST":
		var body struct {
			NotificationId string `json:"notif_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON body"})
			return
		}
		fmt.Println("notif body respose =", body)

		err := h.NotificationUsecase.RemoveNotif(r.Context(), body.NotificationId)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": " err remove notif"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Friend accepted"})
	}
}
