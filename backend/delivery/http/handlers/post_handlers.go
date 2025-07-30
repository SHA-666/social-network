package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	entities "social-network/core/entities"
	internal "social-network/internal"
)

func (h *Handler) HandleGetPosts(w http.ResponseWriter, r *http.Request) {
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

	posts, userPic, err := h.PostUsecase.LoadPost(r.Context(), userID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to load posts"})
		return
	}
	response := map[string]any{
		"posts":    posts,
		"user_pic": userPic,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) HandleCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	c, err := r.Cookie("Token")
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userID := h.SessionUsecase.GetIdOnly(r.Context(), c.Value)
	if userID == "" {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	content := strings.TrimSpace(r.FormValue("content"))
	if content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	post := &entities.Post{
		UserId:  userID,
		Title:   r.FormValue("title"),
		Content: content,
		Privacy: r.FormValue("privacy"),
	}

	// Gestion de l'image optionnelle
	file, header, err := r.FormFile("image")
	if err == nil && file != nil {
		imagePath, err := internal.SavePost(file, header, "")
		if err != nil {
			http.Error(w, "Error saving image", http.StatusInternalServerError)
			return
		}
		post.Image = imagePath
	}

	err = h.PostUsecase.CreatePost(r.Context(), post)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

