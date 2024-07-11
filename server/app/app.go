package app

import (
	"go_rest_app/server/handlers"
	"go_rest_app/server/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	routes *mux.Router
}

var repo repositories.Repository

func (app *App) StartServer(repository repositories.Repository) {
	repo = repository
	app.initRoutes()
	log.Fatal(http.ListenAndServe(":8000", app.routes))
}

func (app *App) initRoutes() {
	handlers.InitHandlers(repo)
	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/authentificate", handlers.AuthentificateUser).Methods("GET")
	router.HandleFunc("/add/record", handlers.AddRecord).Methods("POST")
	app.routes = router
}
