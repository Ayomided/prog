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
	RelativeUrl string
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
	slug := url[strings.LastIndex(url, "/")+1:]
	return &OG{
		Title:       title,
		Image:       image,
		Url:         url,
		RelativeUrl: slug,
		Description: description,
		Type:        ogType,
		SiteName:    siteName,
	}, nil
}

func (og *OG) GenerateMetaOg() ([]byte, error) {
	var tmpl string
	if og.Image != "" {
		tmpl = `
	<meta property="og:title" content="{{.Title}}">
	<meta property="og:url" content="{{.Url}}">
	<meta property="og:image" content="{{ .Image }}" />
	<meta property="og:description" content="{{.Description}}">
	<meta property="og:type" content="{{.Type}}">
	<meta property="og:site_name" content="{{.SiteName}}">
	<meta property="og:logo" content="https://adediiji.uk/static/og-images/logo.png">
	`
	} else {
		tmpl = `
	<meta property="og:title" content="{{.Title}}">
	<meta property="og:url" content="{{.Url}}">
	<meta property="og:image" content="/static/og-images/{{ .RelativeUrl }}.png" />
	<meta property="og:description" content="{{.Description}}">
	<meta property="og:type" content="{{.Type}}">
	<meta property="og:site_name" content="{{.SiteName}}">
	<meta property="og:logo" content="https://adediiji.uk/static/og-images/logo.png">
	`
	}

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

func splitTextIntoLines(title string, maxLineLength int) []string {
	var lines []string
	words := strings.Fields(title)
	currentLine := ""

	for _, word := range words {
		if len(currentLine)+len(word)+1 > maxLineLength { // +1 for space
			lines = append(lines, currentLine)
			currentLine = word // Start a new line with the current word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}

	if currentLine != "" {
		lines = append(lines, currentLine) // Add the last line
	}

	return lines
}

func GenerateOGImage(title, description, date string) (string, error) {
	baseSVG := getBaseSVG()
	if len(title) > 50 {
		title = title[:47] + "..."
	}
	if len(date) > 10 {
		date = date[:10]
	}
	if len(description) > 120 {
		description = description[:117] + "..."
	}

	titleLines := splitTextIntoLines(title, 25)
	descriptionLines := splitTextIntoLines(description, 20)

	var titleSVG string
	var descriptionSVG string

	yOffset := 379.409
	for _, line := range titleLines {
		titleSVG += fmt.Sprintf(`<tspan x="309" y="%.3f">%s</tspan>`, yOffset, line)
		yOffset += 28.0
	}
	y2Offset := 368.227
	for _, line := range descriptionLines {
		descriptionSVG += fmt.Sprintf(`<tspan x="707" y="%.3f">%s</tspan>`, y2Offset, line)
		y2Offset += 36.0
	}

	svgContent := strings.TrimSpace(fmt.Sprintf(baseSVG, titleSVG, descriptionSVG, date))
	return svgContent, nil
}
