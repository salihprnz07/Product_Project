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
	ProductName string  `json:"productName"`
	CategoryId  int     `json:categoryId`
	UnitPrice   float64 `json:unitPrice`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() {
	response, err := http.Get("http://localhost:3000/products") // GET ile okuma işlemlerini yapabiliriz.
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
	product := Product{Id: 9, ProductName: "Microphone2", CategoryId: 1, UnitPrice: 2000.00}
	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct)) //POST ile ekleme, düzenleme ve silme işlemlerini yapabiliriz.

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse Product

	json.Unmarshal(bodyBytes, &productResponse)
	fmt.Println("Kaydedildi", productResponse)
}
