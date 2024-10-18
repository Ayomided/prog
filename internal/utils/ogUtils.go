package utils

import (
	"bytes"
	"fmt"
	"html/template"
)

type OG struct {
	Title       string
	Image       string
	Url         string
	Description string
	Type        string
	SiteName    string
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
