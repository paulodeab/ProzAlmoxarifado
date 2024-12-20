package route

import (
	"fmt"
	"html/template"
	"net/http"
)

type TemplateName string

const (
	IndexTemplate TemplateName = "index.html"
	LoginTemplate TemplateName = "login.html"
	RegisterTemplate TemplateName = "registerproduct.html"
)

var templates = map[TemplateName]*template.Template{
	IndexTemplate: template.Must(template.ParseFiles("template/index.html")),
	LoginTemplate: template.Must(template.ParseFiles("template/login.html")),
	RegisterTemplate: template.Must(template.ParseFiles("template/registerproduct.html")),
}

func GetTemplate(name TemplateName) *template.Template {
	return templates[name]
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := GetTemplate(IndexTemplate)
	tmpl.Execute(w, nil)
}


func RegisterProductHandler(w http.ResponseWriter, r *http.Request){
	tmpl := GetTemplate(RegisterTemplate);
	tmpl.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		username := r.FormValue("username");
		password := r.FormValue("password");

		fmt.Print(username);
		fmt.Print(password);

		if (username == "admin" && password == "1234") {
			http.Redirect(w, r, "/index", http.StatusSeeOther)
			return
		}else{
			fmt.Print("Falha ao logar");
		}

	}
	tmpl := GetTemplate(LoginTemplate)
	tmpl.Execute(w, nil)
}
