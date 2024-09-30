package handlers

import (
	"net/http"

	"github.com/Ayomided/prog.git/sqlite"
)

func HomeHandler(db *sqlite.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
