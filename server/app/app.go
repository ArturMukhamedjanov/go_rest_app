package app

import (
	"go_rest_app/server/repositories"

	"github.com/gorilla/mux"
)

type App struct {
	route *mux.Router
}

var repo repositories.Repository

func (app *App) StartServer(repository repositories.Repository) {
	repo = repository
	app.initRoutes()
}

func (app *App) initRoutes() {
	router := mux.NewRouter()
	
}


