package handlers

import (
	"groupie/internal/api"
	"groupie/internal/models"
	"groupie/internal/render"
	"net/http"
	"text/template"
)

var t *template.Template

// func init() {
// 	t = template.Must(template.ParseGlob("/templates/*.html"))
// }

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, errors[404])
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, errors[405])
	}

	Artists, err := api.GetArtists()

	if err != nil {
		ErrorHandler(w, r, errors[500])
	}

	Data := make(map[string]interface{})

	Data["Artists"] = Artists

	render.RenderTemplate(w, "index.page.html", &models.TemplateData{
		Data: Data,
	})

}
