package utils

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Ayomided/prog/internal/config"
)

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod"`
	ChangeFreq string  `xml:"changefreq"`
	Priority   float32 `xml:"priority"`
}

func GenerateSitemap(cfg *config.Config) error {
	baseURL := "https://adediiji.uk"
	postsDir := cfg.PostsPath // Adjust this to your posts directory
	staticPages := []string{"/", "/about"}

	sitemap := Urlset{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}

	// Add static pages
	for _, page := range staticPages {
		sitemap.Urls = append(sitemap.Urls, Url{
			Loc:        baseURL + page,
			LastMod:    time.Now().Format("2006-01-02"),
			ChangeFreq: "weekly",
			Priority:   0.8,
		})
	}

	// Add posts
	err := filepath.Walk(postsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			// Remove the file extension and convert to URL path
			urlPath := strings.TrimSuffix(filepath.ToSlash(strings.TrimPrefix(path, postsDir)), ".md")
			sitemap.Urls = append(sitemap.Urls, Url{
				Loc:        baseURL + "/posts" + urlPath,
				LastMod:    info.ModTime().Format("2006-01-02"),
				ChangeFreq: "monthly",
				Priority:   0.5,
			})
		}
		return nil
	})

	if err != nil {
		return err
	}

	// Create sitemap.xml in the static folder
	file, err := os.Create("static/sitemap.xml")
	if err != nil {
		return err
	}
	defer file.Close()

	enc := xml.NewEncoder(file)
	enc.Indent("", "  ")
	if err := enc.Encode(sitemap); err != nil {
		return err
	}

	return nil
}
