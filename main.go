package main

import (
	"ProzAlmoxarifado/src/route"
	"fmt"
	"net/http"
	
)

func main() {
	fmt.Printf("SERVIDOR EM EXECUÇÃO em 8080.....")
	http.HandleFunc("/", route.LoginHandler)
	http.HandleFunc("/index", route.IndexHandler)
	http.HandleFunc("/registerproduct", route.RegisterProductHandler)
	http.ListenAndServe(":8080", nil)
}
