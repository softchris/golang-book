package main

import (
	"fmt"
)

type Row struct {
	Title       string
	Description string
	Quantity    int
	UnitPrice   float32
}

func main() {
	row := Row{
		Title:       "LEGO set",
		Description: "4000 pieces",
		Quantity:    1,
		UnitPrice:   600,
	}
	row2 := Row{
		Title:       "Plushy",
		Description: "plush toy",
		Quantity:    3,
		UnitPrice:   5,
	}

	basket := make([]Row, 0)
	basket = append(basket, row)
	basket = append(basket, row2)

	var sum int = 0
	for i := 0; i < len(basket); i++ {
		current := basket[i]
		fmt.Println(current)
		sum += current.Quantity * int(current.UnitPrice)
	}
	fmt.Println("Total", sum)
}
