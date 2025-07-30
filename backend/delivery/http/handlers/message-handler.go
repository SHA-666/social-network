package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (h *Handler) HandleMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Token")
	if err != nil || c.Value == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Token cookie required"})
		return
	}
	userID := h.SessionUsecase.GetIdOnly(r.Context(), c.Value)
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user"})
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/messages/")
	if path == "" {
		http.Error(w, `{"error":"friend_id required"}`, http.StatusBadRequest)
		return
	}
	friendID := path
	fmt.Println("friend_id ))))==== ", friendID)

	res, err := h.MessageUsecase.LoadConv(r.Context(), userID, friendID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed while load conv"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"messages": res,
	})
}
