package route

import (
	"ProzAlmoxarifado/src/model"
	"ProzAlmoxarifado/src/service"
	"fmt"
	"html/template"
	"net/http"
)

type TemplateName string
type PageResponses struct {
	ErrorLogin string
}

const (
	IndexTemplate    TemplateName = "index.html"
	LoginTemplate    TemplateName = "login.html"
	RegisterTemplate TemplateName = "registerproduct.html"
	ProductTemplate  TemplateName = "products.html"
)

var templates = map[TemplateName]*template.Template{
	IndexTemplate:    template.Must(template.ParseFiles("template/index.html")),
	LoginTemplate:    template.Must(template.ParseFiles("template/login.html")),
	RegisterTemplate: template.Must(template.ParseFiles("template/registerproduct.html")),
	ProductTemplate:  template.Must(template.ParseFiles("template/products.html")),
}

func GetTemplate(name TemplateName) *template.Template {
	return templates[name]
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := GetTemplate(IndexTemplate)
	tmpl.Execute(w, nil)
}

// Insere um produto
func RegisterProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		product := model.Product{
			NomeItem:      r.FormValue("nomeItem"),
			Quantidade:    r.FormValue("quantidade"),
			UnidadeMedida: r.FormValue("unidadeMedida"),
			Local:         r.FormValue("local"),
		}
		service.InsertProduct(product)
	}

	tmpl := GetTemplate(RegisterTemplate)
	tmpl.Execute(w, nil)

}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		products, err := service.SelectProducts()
		fmt.Println(products)
		if err != nil {
			http.Error(w, "Erro ao recuperar produtos", http.StatusInternalServerError)
			return
		}

		tmpl := GetTemplate(ProductTemplate)
		tmpl.Execute(w, products)
	}
}

// Realiza o Login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Caso o método seja POST, processa o login
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Dados para passar ao template de login
		data := PageResponses{}

		// Verifica as credenciais
		if !service.Login(username, password) {
			// Se a autenticação falhar, mostra a mensagem de erro
			data.ErrorLogin = "Falha ao logar"
			// Exibe o template de login com a mensagem de erro
			tmpl := GetTemplate(LoginTemplate)
			tmpl.Execute(w, data)
			return
		} else {
			// Se o login for bem-sucedido, redireciona para a página de índice
			http.Redirect(w, r, "/index", http.StatusSeeOther)
			return
		}
	}

	// Se não for um POST, apenas exibe o template de login
	tmpl := GetTemplate(LoginTemplate)
	tmpl.Execute(w, nil)
}
