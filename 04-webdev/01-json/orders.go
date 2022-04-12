package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type OrderItem struct {
	Id       int     `json: "id"`
	Quantity int     `json: "quantity"`
	Total    float32 `json: "total"`
}

type Order struct {
	Id    int         `json: "id"`
	Items []OrderItem `json: items`
}

type Response struct {
	Orders []Order `json: orders`
}

func main() {
	file, _ := ioutil.ReadFile("orders.json")

	data := Response{}

	_ = json.Unmarshal([]byte(file), &data)
	fmt.Println(data)
	for i := 0; i < len(data.Orders); i++ {
		fmt.Println("Order Id: ", data.Orders[i].Id)

		for j := 0; j < len(data.Orders[i].Items); j++ {
			item := data.Orders[i].Items[j]
			fmt.Println("Item id", item.Id)
			fmt.Println("Item quantity", item.Quantity)
			fmt.Println("Item total", item.Total)
		}
	}
}
