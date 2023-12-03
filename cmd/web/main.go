package main

import (
	"fmt"
	"groupie/pkg/config"
	"groupie/pkg/handlers"
	"groupie/pkg/render"
	"net/http"
)

func main() {
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
