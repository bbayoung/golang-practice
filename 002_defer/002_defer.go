package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
	for i := 0; i < 2; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}

	fmt.Println("Main Done")
}
