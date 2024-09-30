package handlers

import (
	"html/template"
	"net/http"

	"github.com/Ayomided/prog/internal/utils"
)

func HomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articles, err := utils.GetAllArticles("./posts")
		if err != nil {
			http.Error(w, "Error getting posts", http.StatusInternalServerError)
			return
		}
		tpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, articles)
		if err != nil {
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
		}
	})
}
