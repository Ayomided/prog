package handlers

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/Ayomided/prog/article"
	"github.com/Ayomided/prog/internal/utils"
	"github.com/adrg/frontmatter"
)

type Post struct {
	Title       string `toml:"title"`
	Slug        string `toml:"slug"`
	Content     template.HTML
	OGMeta      template.HTML
	Date        time.Time `toml:"date"`
	Author      Author    `toml:"author"`
	Description string    `toml:"description"`
}

type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

type SlugReader interface {
	Read(posts fs.FS, slug string) (string, error)
}

type FileReader struct{}

func (fsr FileReader) Read(posts fs.FS, slug string) (string, error) {
	f, err := posts.Open(slug + ".md")
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func getFullURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s%s", scheme, r.Host, r.URL.RequestURI())
}

func PostHandler(sl SlugReader, posts, templatesFS fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		post.Slug = r.PathValue("slug")
		markdownText, err := sl.Read(posts, post.Slug)
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
		out, err := parser.ParseBlog(rest)
		if err != nil {
			http.Error(w, "Error converting markdown", http.StatusInternalServerError)
		}
		post.Content = template.HTML(out)
		tpl, err := template.ParseFS(templatesFS, "post.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		uri := getFullURL(r)
		og, err := utils.NewMetaOg(post.Title, "", uri, post.Description, "article", "David Adediji")
		if err != nil {
			http.Error(w, "Error creating open graph tags", http.StatusInternalServerError)
			return
		}

		metaTags, err := og.GenerateMetaOg()
		if err != nil {
			http.Error(w, "Error generating Open Graph tags", http.StatusInternalServerError)
			return
		}
		post.OGMeta = template.HTML(metaTags)

		err = tpl.Execute(w, post)
	}
}
