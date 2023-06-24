package render

import (
	"go_forms/models"
	"html/template"
	"net/http"
)



func RenderTemplate(w http.ResponseWriter, url string, td *models.TemplateData) {
	tmpl, _ := template.ParseFiles("./templates/"+url,)

	tmpl.Execute(w, td)
}
