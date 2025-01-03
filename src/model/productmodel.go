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


const INSERT_PRODUCT string = "INSERT INTO produto(nomeItem, quantidade, unidadeMedida, local) VALUES (?, ?, ?, ?)";
const SELECT_ALL_PRODUCT string = "SELECT id, nomeItem, quantidade, unidadeMedida, local FROM produto";

// Função para inserir um produto na tabela
func Insert(product Product) int {
    db := Connect();
  
    //Prepared stament para inserir os dados
    stmt, err := db.Prepare(INSERT_PRODUCT)
    if err != nil {
        fmt.Println("Insert command error!", err);
        panic(err);
    }
    defer stmt.Close() //fecha as conexões e as queries quando tudo acabar

    response, err := stmt.Exec(product.NomeItem, product.Quantidade, product.UnidadeMedida, product.Local);
    if err != nil {
        fmt.Println("Erro ao inserir:", err);
        return 500;
    }

    lastId, err := response.LastInsertId()
    if err != nil {
        fmt.Println("Erro ao obter o ID do último registro inserido:", err);
        return 500;
    }

	defer db.Close();
    fmt.Println("Produto inserido com sucesso! ID:", lastId);
    return 201;
}

func SelectAllProducts() ([]Product, error ){

    db := Connect();
    defer db.Close();

    //Para select se usa isso
    rows, err := db.Query(SELECT_ALL_PRODUCT);

    if err != nil{
        fmt.Println("Select Command Error!", err);
        return nil, err;
    }
    defer rows.Close();

    //Criando um tipo ou objeto;
    var products []Product;

    for rows.Next(){
        var product Product;
        err := rows.Scan(&product.ID, &product.NomeItem, &product.Quantidade, &product.UnidadeMedida, &product.Local);
        if err != nil {
            fmt.Println("Erro ao escanear o produto: ", err);
            return nil, err;
        }
        products = append(products, product);
    }


    return products, nil;

}