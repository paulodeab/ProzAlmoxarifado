package service

import (
	"ProzAlmoxarifado/src/model"
)

// acessa via nome do pacote/nome da função
func InsertProduct(product model.Product) int {

	return model.Insert(product)

}

func SelectProducts() ([]model.Product, error) {

	return model.SelectAllProducts()
}
