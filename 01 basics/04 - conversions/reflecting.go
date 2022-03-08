package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var no int = 100
	fmt.Println(reflect.TypeOf(no))

	var intStr string = "100"
	fourBaseEightBitInt, _ := strconv.ParseInt(intStr, 4, 8)    // becomes no 16 and 8bit
	tenBaseSixteenBitInt, _ := strconv.ParseInt(intStr, 10, 16) // no 100,  16 bit
	fmt.Println(fourBaseEightBitInt)
	fmt.Println(tenBaseSixteenBitInt)
	fmt.Println(reflect.TypeOf(fourBaseEightBitInt))
	fmt.Println(reflect.TypeOf(tenBaseSixteenBitInt))
}
