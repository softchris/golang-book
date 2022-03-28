package main

import "fmt"

func Advanced() {
	var no int
	names := ""
	var name string
	fmt.Println("Enter number of players:")
	fmt.Scan(&no)

	for i := 0; i < no; i++ {
		fmt.Printf("Enter Player %d name:", i+1)
		fmt.Scan(&name)
		names += name + " "
	}
	fmt.Println("Players are:", names)
}

func main() {
	// var response string
	fmt.Println("hi")

	// // one input
	// fmt.Scan(&response)
	// fmt.Println("User typed ", response)

	// var a1, a2 string

	// // multiple input
	// fmt.Scan(&a1, &a2)

	// // formatted string
	// str := fmt.Sprintf("a1: %s a2: %s", a1, a2)
	// fmt.Println(str)

	// var prefix string
	// var no int
	// // in110
	// fmt.Scanf("%3s%d", &prefix, &no)
	// fmt.Printf("prefix: %s, invoice no: %d", prefix, no)
	// fmt.Println("Enter player names (separated by space):")
	// var player1, player2 string
	// fmt.Scan(&player1, &player2)
	// fmt.Println("Players are: ", player1, player2)
	Advanced()
}
