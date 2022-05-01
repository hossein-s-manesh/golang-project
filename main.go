package main

import (
	"fmt"
	"myapp/data"
)

func main() {
	fmt.Println("are you redy")
	say(data.Hide)
	Vark("how are you" , "are you go ing")
	conver()
	dat()

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

// tipe conversion

func conver(){
	var(
		son =50
		day=66.9
		month=2
	)

	hos:=int(son)*int(month)*int(day)
	had := int(son)+int(month)+int(day)
	haz := float64(son)+float64(month)+float64(day)
	hor := float32(day)
	hoz := float64(day)
	hog := int(day)
	fmt.Println(hos)
	fmt.Println(had)
	fmt.Println(haz)
	fmt.Println(hor, hoz , hog)

}


// type aliaing

type int float64
//type int float32

func dat(){
	var(
		c=25.5
		d=58.9
		
	)

	u := float64(c)
	v := float64(d)
	y := int(c)
	k:= int(d)
	p :=float32(d)


	fmt.Println(u,v,y,k,p)
}