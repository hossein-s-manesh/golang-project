package main

import (
	"fmt"
	"myapp/data"
)

func main() {
	fmt.Println("are you redy")
	say(data.Hide)
}

func say(name string) {
	fmt.Println("are you redy", name)
}