package project

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func GetAllProducts() {
	response, err := http.Get(" http://localhost:3000/products")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	fmt.Println(products)
}

func AddProduct() {

	product := Product{Id: 4, productName: "Telephone", categoryId: 1, unitPrice: 6000}
	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	fmt.Println("Saved")

}

// //
// I get this error when I write codes but my codes work properly.
// could not import gocourse/project (cannot find package "gocourse/project" in any of
// // 	C:\Program Files\Go\src\gocourse\project (from $GOROOT)
// // 	C:\Users\cemdo\go\src\gocourse\project (from $GOPATH))
