package handlers

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/Ayomided/prog/internal/utils"
)

type home struct {
	OGMeta   template.HTML
	Articles []utils.Post
}

func HomeHandler(postFS, templatesFS fs.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var home home
		articles, err := utils.GetAllArticles(postFS)
		if err != nil {
			http.Error(w, "Error getting posts", http.StatusInternalServerError)
			return
		}
		home.Articles = articles
		tpl, err := template.ParseFS(templatesFS, "index.html")
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
		home.OGMeta = template.HTML(metaTags)

		err = tpl.Execute(w, home)
		if err != nil {
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
		}
	})
}
