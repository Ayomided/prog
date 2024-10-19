package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
)

type OG struct {
	Title       string
	Image       string
	Url         string
	Description string
	Type        string
	SiteName    string
}

func getBaseSVG() string {
	var baseSVG string
	content, err := os.ReadFile("internal/utils/base.svg")
	if err != nil {
		baseSVG = ""
	} else {
		baseSVG = string(content)
	}
	return baseSVG
}

func NewMetaOg(title, image, url, description, ogType, siteName string) (*OG, error) {
	if title == "" {
		return nil, fmt.Errorf("Title is required")
	}
	if url == "" {
		return nil, fmt.Errorf("Url is required")
	}
	if image == "" {
		return nil, fmt.Errorf("Image is required")
	}
	return &OG{
		Title:       title,
		Image:       image,
		Url:         url,
		Description: description,
		Type:        ogType,
		SiteName:    siteName,
	}, nil
}

func (og *OG) GenerateMetaOg() ([]byte, error) {
	tmpl := `
	<meta property="og:title" content="{{.Title}}">
	<meta property="og:image" content="{{.Image}}">
	<meta property="og:url" content="{{.Url}}">
	<meta property="og:image" content="/static/og-images/{{ .Url }}.png" />
	<meta property="og:description" content="{{.Description}}">
	<meta property="og:type" content="{{.Type}}">
	<meta property="og:site_name" content="{{.SiteName}}">
	`
	t, err := template.New("og").Parse(tmpl)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, og)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func GenerateOGImage(title, description, date string) (string, error) {
	baseSVG := getBaseSVG()
	if len(title) > 15 {
		title = title[:12] + "..."
	}
	if len(date) > 10 {
		date = date[:10]
	}
	if len(description) > 20 {
		description = description[:17] + "..."
	}

	svgContent := strings.TrimSpace(fmt.Sprintf(baseSVG, title, description, date))
	return svgContent, nil
}
