package router

import (
	"github.com/gorilla/mux"
)

func MakeRouter() *mux.Router {
	r := mux.NewRouter()
	//makeStaticRouter(r)
	//makeMusicRouter(r)
	makeWebRouter(r)
	return r
}

func makeWebRouter(mainRouter *mux.Router) {
	apiRouter := mainRouter.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(jsonResponseMiddleware)
	apiRouter.HandleFunc("/status", JSONHandler(statusHandler)).Methods("GET")
	apiRouter.HandleFunc("/login", JSONHandler(loginHandler)).Methods("POST")

	authRouter := apiRouter.PathPrefix("").Subrouter()
	authRouter.Use(authorizeByJWT)
	authRouter.HandleFunc("/temperature", JSONHandler(userTemperatureHandler)).Methods("GET")
	authRouter.HandleFunc("/pressure", JSONHandler(userPressureHandler)).Methods("GET")
	authRouter.HandleFunc("/humidity", JSONHandler(userHumidityHandler)).Methods("GET")

	adminRouter := authRouter.PathPrefix("/admin").Subrouter()
	//adminRouter.Use(authorizeAdmin)
	adminRouter.HandleFunc("/battery", JSONHandler(authHandler)).Methods("GET")
	adminRouter.HandleFunc("/info", JSONHandler(authHandler)).Methods("GET")
	adminRouter.HandleFunc("/work", JSONHandler(authHandler)).Methods("GET")

}
