package handlers

import (
	"html/template"
	"io/fs"
	"net/http"
)

func AboutHandler(path fs.FS) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFS(path, "templates/about.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, "")
		if err != nil {
			http.Error(w, "Error rendering page", http.StatusInternalServerError)
		}
	})
}
