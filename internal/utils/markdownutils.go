package utils

import (
	"html/template"
	"io/fs"
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

func parseMarkdown(postsFS fs.FS, filename string) (Post, error) {
	content, err := fs.ReadFile(postsFS, filename)
	if err != nil {
		return Post{}, err
	}

	var post Post
	part, err := frontmatter.Parse(strings.NewReader(string(content)), &post)
	if err != nil {
		return Post{}, err
	}

	post.Slug = filepath.Base(filename[:len(filename)-len(filepath.Ext(filename))])

	parser := article.NewParser()
	out, err := parser.ParseBlog(part)
	if err != nil {
		return Post{}, err
	}
	post.Content = template.HTML(out)

	return post, nil
}

func GetAllArticles(postsFS fs.FS) ([]Post, error) {
	var posts []Post

	err := fs.WalkDir(postsFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".md" {
			post, err := parseMarkdown(postsFS, path)
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
