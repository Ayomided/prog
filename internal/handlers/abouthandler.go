package handlers

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/Ayomided/prog/internal/utils"
)

type about struct {
	OGMeta template.HTML
}

func AboutHandler(templatesFS fs.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var about about
		tpl, err := template.ParseFS(templatesFS, "about.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		uri := getFullURL(r)
		og, err := utils.NewMetaOg("David Adediji", "/static/og-images/home.png", uri, "David's Pika Pika", "website", "David Adediji")
		if err != nil {
			http.Error(w, "Error generating Open Graph tags", http.StatusInternalServerError)
			return
		}
		metaTags, err := og.GenerateMetaOg()
		if err != nil {
			http.Error(w, "Error generating Open Graph tags", http.StatusInternalServerError)
			return
		}
		about.OGMeta = template.HTML(metaTags)
		err = tpl.Execute(w, about)
		if err != nil {
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
		}
	})
}
