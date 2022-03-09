// create an array
// inferred length
// get element
package main

import "fmt"

var array [3]int

func main() {
	cities := [5]string{"NY", "LA"}
	ids := [...]int{1, 2, 3, 4}
	array[0] = 1
	var matrix [2][2]int
	matrix[0][0] = 1

	fmt.Println(array[0])
	fmt.Println(cities[0])
	fmt.Println(ids[3])
	fmt.Println(matrix[0][0])

	items := [5]int{1, 2, 3, 4, 5}
	part := items[2:4] // 2,3
	fmt.Println("part", part)

	var numbers []int
	numbers = append(numbers, 1)
	fmt.Println("num:", numbers)
	numbers = append(numbers, 2)
	fmt.Println("num:", numbers)

	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

}
