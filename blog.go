package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Ayomided/prog.git/article"
	"github.com/adrg/frontmatter"
)

type SlugReader interface {
	Read(slug string) (string, error)
}

type FileReader struct{}

type Post struct {
	Title   string `toml:"title"`
	Slug    string `toml:"slug"`
	Content template.HTML
	Author  Author `toml:"author"`
}

type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

func (fsr FileReader) Read(slug string) (string, error) {
	f, err := os.Open(slug + ".md")
	if err != nil {
		return "", err
	}
	defer f.Close()
	log.Println(f.Name())
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func PostHandler(sl SlugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		post.Slug = r.PathValue("slug")
		markdownText, err := sl.Read(post.Slug)
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
		tpl, err := template.ParseFiles("template.gohtml")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, post)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /posts/{slug}", PostHandler(FileReader{}))

	err := http.ListenAndServe(":3030", mux)
	log.Println("Listening on :3030")
	if err != nil {
		log.Fatal(err)
	}
}
