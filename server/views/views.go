package views

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed html/*.html
var fs embed.FS
var templates = template.Must(template.ParseFS(fs, "html/*.html"))

func Render(w http.ResponseWriter, r *http.Request, name string, data interface{}) error {
	return templates.ExecuteTemplate(w, name, data)
}
