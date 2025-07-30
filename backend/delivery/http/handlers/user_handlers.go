package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	entities "social-network/core/entities"
	internal "social-network/internal"

	"github.com/google/uuid"
)

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	u := &entities.User{
		Username:    strings.TrimSpace(r.FormValue("username")),
		FirstName:   strings.TrimSpace(r.FormValue("first_name")),
		LastName:    strings.TrimSpace(r.FormValue("last_name")),
		Email:       strings.TrimSpace(r.FormValue("email")),
		Password:    r.FormValue("password"),
		DateOfBirth: r.FormValue("date_of_birth"),
		AboutMe:     strings.TrimSpace(r.FormValue("about_me")),
		City:        strings.TrimSpace(r.FormValue("city")),
	}
	file, header, _ := r.FormFile("profile_pic")
	avatarPath, err := internal.SavePics(file, header, "/default_profil_pic.jpeg")
	if err != nil {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Invalid form", http.StatusBadRequest)
			return
		}
		http.Error(w, "Error saving avatar", http.StatusInternalServerError)
		return
	}
	u.ProfilePic = avatarPath
	if err := u.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.UserUsecase.Register(r.Context(), u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Inscription réussie from delivery ;°)",
	})
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	u := &entities.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	id, err := h.UserUsecase.Login(r.Context(), u)
	if err != nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusUnauthorized)
		return
	}
	sessionID := uuid.NewString()
	session := &entities.Session{
		Token:     sessionID,
		UserID:    id, // ID
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}
	if err := h.SessionUsecase.CreateSession(r.Context(), session); err != nil {
		http.Error(w, "Erreur lors de la création de session", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "Token",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // true en production avec HTTPS
		SameSite: http.SameSiteLaxMode,
		Expires:  session.ExpiresAt,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "login successful",
		"token":   session.Token,
		"id":      strconv.Itoa(id),
	})
}

func (h *Handler) HandleProfil(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Token")
	if err != nil || cookie.Value == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Token cookie required"})
		return
	}

	userID := h.SessionUsecase.GetIdOnly(r.Context(), cookie.Value)
	if userID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid session"})
		return
	}

	user, found := h.UserUsecase.GetInfo(r.Context(), userID)
	if !found {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Erreur JSON", http.StatusInternalServerError)
	}
}
