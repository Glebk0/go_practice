package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	for i, number := range numbers {
		if numbers[i] % 2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}