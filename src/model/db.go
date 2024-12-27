package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)


//Conexao com o banco
func Connect() *sql.DB {

	dsn := "root:admin@tcp(127.0.0.1:3306)/db_proz";

	var err error;

	db, err := sql.Open("mysql", dsn);
	if err != nil{
		fmt.Print("ERRO no banco", err);
		panic(err);
	}

	fmt.Print("Conexão bem sucedida!!");
	return db;
}

