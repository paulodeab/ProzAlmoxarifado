package model

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

// Definição da struct Product
type Product struct {
    ID            int
    NomeItem      string
    Quantidade    string
    UnidadeMedida string
    Local         string
}

var insertProduct string = "INSERT INTO produto(nomeItem, quantidade, unidadeMedida, local) VALUES (?, ?, ?, ?)"

// Função para inserir um produto na tabela
func Insert(product Product) int {
    db := Connect()
  

    stmt, err := db.Prepare(insertProduct)
    if err != nil {
        fmt.Println("Insert command error!", err)
        panic(err)
    }
    defer stmt.Close() //fecha as conexões e as queries quando tudo acabar

    response, err := stmt.Exec(product.NomeItem, product.Quantidade, product.UnidadeMedida, product.Local)
    if err != nil {
        fmt.Println("Erro ao inserir:", err)
        panic(err) // interrompe e faz uma limpeza da pilha de execução
    }

    lastId, err := response.LastInsertId()
    if err != nil {
        fmt.Println("Erro ao obter o ID do último registro inserido:", err)
        panic(err)
    }
	defer db.Close()
    fmt.Println("Produto inserido com sucesso! ID:", lastId)
    return 201
}