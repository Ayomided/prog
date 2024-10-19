package handlers

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/Ayomided/prog/article"
	"github.com/Ayomided/prog/internal/utils"
	"github.com/adrg/frontmatter"
)

func OGImageHandler(sl SlugReader, posts fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		slug := r.PathValue("path")
		// path := path.Join("/posts/", slug)
		markdownText, err := sl.Read(posts, slug)
		if err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		rest, err := frontmatter.Parse(strings.NewReader(markdownText), &post)
		if err != nil {
			http.Error(w, "Error parsing frontmatter", http.StatusInternalServerError)
			return
		}

		parser := article.NewParser()
		_, err = parser.ParseBlog(rest)
		if err != nil {
			http.Error(w, "Error converting markdown", http.StatusInternalServerError)
		}

		svgContent, err := utils.GenerateOGImage(post.Title, post.Description, post.Date.String())
		if err != nil {
			http.Error(w, "Failed to generate SVG", http.StatusInternalServerError)
			return
		}

		fmt.Println(svgContent)

		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write([]byte(svgContent))
	}
}
