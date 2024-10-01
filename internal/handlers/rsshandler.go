package handlers

import (
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/Ayomided/prog/internal/utils"
	"github.com/gorilla/feeds"
)

func RssHandler(postFS fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		feed := &feeds.Feed{
			Title:       "David Adediji | blog",
			Link:        &feeds.Link{Href: "https://prog.fly.dev"},
			Description: "Hello! I am David, I share my thoughts here",
			Author:      &feeds.Author{Name: "David Adediji", Email: "idavid.adediji@gmail.com"},
			Created:     time.Now(),
		}

		articles, err := utils.GetAllArticles(postFS)
		if err != nil {
			http.Error(w, "Error getting posts", http.StatusInternalServerError)
			return
		}

		var feedItems []*feeds.Item
		for _, article := range articles {
			feedItems = append(feedItems,
				&feeds.Item{
					Id:      fmt.Sprintf("tag: %v, %v:/articles/%v", time.Now().Year(), article.Date, article.Slug),
					Title:   article.Title,
					Link:    &feeds.Link{Href: fmt.Sprintf("https://prog.fly.dev/articles/%v", article.Slug)},
					Created: article.Date,
				})
		}
		feed.Items = feedItems
		rss, err := feed.ToRss()
		if err != nil {
			http.Error(w, "Failed to generate RSS", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/rss+xml")

		_, err = w.Write([]byte(rss))
		if err != nil {
			http.Error(w, "Failed to write RSS to response", http.StatusInternalServerError)
			return
		}
	}
}
