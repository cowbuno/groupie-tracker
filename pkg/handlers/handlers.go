package handlers

import (
	"groupie/pkg/api"
	"groupie/pkg/models"
	"groupie/pkg/render"
	"net/http"
	"strconv"
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

	Artists, err := api.GetAllArtists()
	if err != nil {
		ErrorHandler(w, r, errors[500])
	}

	Data := make(map[string]interface{})

	Data["Artists"] = Artists

	render.RenderTemplate(w, "index.page.html", &models.TemplateData{
		Data: Data,
	})
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		ErrorHandler(w, r, errors[404])
		return
	}

	if r.Method != "GET" {
		ErrorHandler(w, r, errors[405])
		return
	}

	artistID := r.URL.Query().Get("id")

	_, err := strconv.Atoi(artistID)

	if err != nil {
		ErrorHandler(w, r, errors[400])
		return
	}

	Artist, err := api.GetArtistByID(artistID)
	if err != nil {
		ErrorHandler(w, r, errors[404])
		return
	}

	Data := make(map[string]interface{})

	Data["Artist"] = Artist

	render.RenderTemplate(w, "artist.page.html", &models.TemplateData{
		Data: Data,
	})

}
