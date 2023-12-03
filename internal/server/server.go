package server

import (
	"fmt"
	"groupie/internal/config"
	"groupie/internal/handlers"
	"groupie/internal/render"
	"net/http"
)

func RunServer() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		fmt.Println("ERROR CREATING CACHE")
		return
	}

	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplate(&app)
	
	mux := http.NewServeMux()

	fmt.Println("Listening on http://localhost:8080/ ... ")


	mux.HandleFunc("/", handlers.IndexHandler)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Could not start sever on port http://localhost:8080/")
		return
	}
}
