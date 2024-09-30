package utils

import (
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/Ayomided/prog/article"

	"github.com/adrg/frontmatter"
)

type Post struct {
	Title   string `toml:"title"`
	Slug    string `toml:"slug"`
	Content template.HTML
	Date    time.Time `yaml:"date"`
	Author  Author    `toml:"author"`
}

type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

func parseMarkdown(filename string) (Post, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return Post{}, err
	}

	var post Post
	part, err := frontmatter.Parse(strings.NewReader(string(content)), &post)
	if err != nil {
		return Post{}, err
	}

	parser := article.NewParser()
	out, err := parser.ParseBlog(part)
	if err != nil {
		return Post{}, err
	}
	post.Content = template.HTML(out)

	return post, nil
}

func GetAllArticles(dir string) ([]Post, error) {
	var posts []Post

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".md" {
			post, err := parseMarkdown(path)
			if err != nil {
				return err
			}
			posts = append(posts, post)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	return posts, nil
}
