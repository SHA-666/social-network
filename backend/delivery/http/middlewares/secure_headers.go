package http

import (
	"net/http"
)

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy",
			"default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")

		w.Header().Set("X-XSS-Protection", "1; mode=block") // 1=active // 0=pas active
		w.Header().Set("Server", "")                        // no info to get from my headers

		next.ServeHTTP(w, r)
	})
}
