package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	productName string  `json: "productName`
	categoryId  int     `json:"categoryId"`
	unitPrice   float64 `json: "unitPrice"`
}

type Category struct {
	id           int    `json:"id"`
	categoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	response, err := http.Get(" http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil
}

func AddProduct() (Product, error) {

	product := Product{Id: 4, productName: "Telephone", categoryId: 1, unitPrice: 6000}
	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		return Product{}, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse Product

	json.Unmarshal(bodyBytes, &productResponse)

	return productResponse, nil

}

// //
// I get this error when I write codes but my codes work properly.
// could not import gocourse/project (cannot find package "gocourse/project" in any of
// // 	C:\Program Files\Go\src\gocourse\project (from $GOROOT)
// // 	C:\Users\cemdo\go\src\gocourse\project (from $GOPATH))
