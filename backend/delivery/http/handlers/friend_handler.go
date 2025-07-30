package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) HandleFriends(w http.ResponseWriter, r *http.Request) {
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
	switch r.Method {
	case "GET":
		peopleToAdd, err := h.FriendUsecase.LoadFriend(r.Context(), userID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed while get peopl"})
			return
		}
			fmt.Println("geeeeeeeeeet",peopleToAdd)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{
			"friends": peopleToAdd,
		})

	case "POST":
		var body struct {
			Action     string `json:"action"`
			ReceiverID string `json:"user_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON body"})
			return
		}
		fmt.Println("user = ", userID, " // body = ", body)

		switch body.Action {
		case "add":
			receiverId := body.ReceiverID

			err = h.FriendUsecase.Addfriend(r.Context(), userID, receiverId)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed while adding Friends"})
				return
			}
			username := h.UserUsecase.UserName(r.Context(), userID)
			h.NotificationUsecase.FriendRequest(r.Context(), userID, receiverId, username)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Friend added successfully"})
		case "reject":
			receiverId := body.ReceiverID
			err = h.FriendUsecase.DeleteFriend(r.Context(), userID, receiverId)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{"message": "Friend probleme delete"})
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "reject friend accepted"})
		case "accepte":
			receiverId := body.ReceiverID
			err = h.FriendUsecase.Acceptefriend(r.Context(), userID, receiverId)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{"message": "Friend accepte probeme"})
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Friend accepted"})
		case "get":
			Friends, err := h.FriendUsecase.GetFriend(r.Context(), userID)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Failed while get peopl"})
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{"friends": Friends,})
		}

	}
}
