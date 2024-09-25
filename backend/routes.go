package backend

import (
	"net/http"

	controllers "sn/backend/controllers"
	middlewares "sn/backend/middlewares"

)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /myapi/sign-in", controllers.SignIn)
	mux.HandleFunc("POST /myapi/register", controllers.Register)

	return middlewares.CORSMiddleware(mux)
}


