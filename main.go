package main

import (
	"fmt"
	"myapp/data"
)

func main() {
	fmt.Println("are you redy")
	say(data.Hide)
	Vark("how are you" , "are you go ing")
}

func say(name string) {
	fmt.Println("are you redy", name)
}

// start learn Var

func Vark(Name string , sod string) {
	var less string = "hossein"
	var less1 int = 34
	var less2 bool = true
	less3 := "ali"
	var ali, hossein, nikoo = 18, 18, 21
//2
	var(
		max=1
		nike string="icon"
		dan = true
		frank = 4.8
		
	)
	//	var ali , hossein , nik oo  = "good boy" , 18 , true
	fmt.Println("lees:",less)
	fmt.Println("lees1:",less1)
	fmt.Println("lee2:",less2)
	fmt.Println("lee3:",less3)
	fmt.Println("ali:",ali)
	fmt.Println("hossein:",hossein)
	fmt.Println("nikoo:",nikoo)
	//2
	fmt.Println("max: ",max)
	fmt.Println("dan: ",dan)
	fmt.Println("nike: ",nike , sod)
	fmt.Println("frank: ",frank ,Name)



}
