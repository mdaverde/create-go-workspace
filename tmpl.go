package main

func envrc() string {
	return "layout go"
}

func mainGo() string {
	return `package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}`
}