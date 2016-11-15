package main

import "fmt"

func main() {
	var sum int = 0
	const limit int = 1000
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
