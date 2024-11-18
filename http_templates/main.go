package http_templates

import (
	"embed"
	"html/template"
)

//go:embed *.html
var templateFS embed.FS

var Template_collection = template.Must(template.ParseFS(templateFS, "*.html"))