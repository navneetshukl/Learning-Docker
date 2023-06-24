package routes

import (
	"fmt"
	"go_forms/models"
	"go_forms/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "form.page.tmpl", &models.TemplateData{})

}
func Data(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	email := r.FormValue("email")
	fmt.Println("Name is : " +name)
	fmt.Println("Email is : "+email)
	w.Write([]byte("Name is " + name+"\n"))
	w.Write([]byte("Email is " + email))

}
