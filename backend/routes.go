package backend

import (
	"net/http"

	controllers "sn/backend/controllers"
	middlewares "sn/backend/middlewares"
)

func Routes() http.Handler {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", controllers.Welcome)
	mux.HandleFunc("/sign-in", controllers.SignIn)
	mux.HandleFunc("/register", controllers.Register)

	return middlewares.CORSMiddleware(mux)
}
