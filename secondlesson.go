package main

import (
	"fmt"
)

func arrReverse(arr [10]int) int {
	var x int
	for i := 9; i >= 0; i-- {
		fmt.Printf("%v ", arr[i])
		x += arr[i]
	}
	return x

}

func main() {

	var num [10]int

	for i := 0; i < 10; i++ {
		fmt.Scan(&num[i])
	}

	arrReverse(num)
	fmt.Println("Конец программы")
}
