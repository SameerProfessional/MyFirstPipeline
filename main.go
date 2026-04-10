package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered from panic:", r)
		}
	}()

	fmt.Println("This is a simple Go program that demonstrates panic and recover.")
}