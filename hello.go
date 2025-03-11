package main

import "fmt"

func main() {
	fmt.Println(Hello("chris"))
}

func Hello(name string) string {
	return "Hello, " + name
}
