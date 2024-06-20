package main

import (
	"go_rest_app/server/app"
	"go_rest_app/server/repositories"
)

func main() {
	ap := app.App{}
	repo := repositories.PostgresRepo{}
	repo.InitDB()
	ap.StartServer(&repo)
}
