package main

import (
	"fmt"
	"myapp/varmod"
	"myapp/callmod"
)

func main() {
	fmt.Println(varmod.See1)
	new(varmod.See2)
	see(varmod.See3, varmod.See4)
}

func new(name string) {
	fmt.Println(callmod.See5, name)
}


func see(name string ,name1 string){
	fmt.Println(name , name1)
}