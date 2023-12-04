package handlers

import (
	"fmt"
	"groupie/pkg/api"
	"groupie/pkg/models"
	"groupie/pkg/render"
	"net/http"
	"text/template"
)

var t *template.Template

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, errors[404])
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, errors[405])
		return
	}

	Artists, err := api.GetArtists()
	fmt.Print(Artists)
	if err != nil {
		ErrorHandler(w, r, errors[500])
	}

	Data := make(map[string]interface{})

	Data["Artists"] = Artists

	render.RenderTemplate(w, "index.page.html", &models.TemplateData{
		Data: Data,
	})
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		ErrorHandler(w, r, errors[404])
		return
	}
	
	if r.Method != "GET" {
		ErrorHandler(w, r, errors[405])
		return
	}
}
