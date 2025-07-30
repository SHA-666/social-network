package http

import (
	"net/http"

	usecases "social-network/application/usecases"
)

type AuthMiddleware struct {
	SessionUsecase usecases.SessionUsecase
}

func NewAuthMiddleware(su usecases.SessionUsecase) *AuthMiddleware {
	return &AuthMiddleware{SessionUsecase: su}
}

func (m *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized: No token", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
