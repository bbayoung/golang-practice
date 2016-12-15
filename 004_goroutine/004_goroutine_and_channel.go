package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	persons := []string{"Anna", "Bob", "Cavin", "Dave", "Eva"}
	watch := make(chan string)

	for _, name := range persons {
		wg.Add(1)
		go Seek(name, watch, &wg)
	}

	wg.Wait()

	select {
	case lastPerson := <- watch:
		fmt.Printf("%s noone\n", lastPerson)
	default:
		//
	}
}

func Seek(name string, watch chan string, wg *sync.WaitGroup) {
	select {
	case peer := <- watch:
		fmt.Printf("%s / %s\n", peer, name)
	case watch <- name:
		fmt.Printf("%s WATCH!!\n", name)
	}
	wg.Done()
}
