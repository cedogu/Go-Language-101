package project

import (
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
	response, err := http.Get("   http://localhost:3000/products")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	fmt.Println(products)
}
