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
	hi()
	io()
	num() // numeric expressions
	main1() // pointers
	map1() // data structures >> map
	map2() // data structures >> map2

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


//praction constant

func hi(){
	const day int= 4
	var day1 int= 4
	const(
		month=5
		sun=7
	)

	const n ,b,c =1,2,3
	
	//  day = day +1
	day1 = day1+1

	fmt.Println(day ,month , sun , n, b, c)
}



// iota

func io(){

	
	const(
		b= iota+1
		c=iota 
		d=iota 
		k=iota
	)
	h:=b
	fmt.Println(b,c,d,k,h)
}



// umeric expressions
func num(){
	const(
		c=19.0
		a=19
		b=18
	)
	e:= c/2
	p:= a/2
	n:=a+c
	fmt.Println("p >" ,p)
	fmt.Println("c >" ,e)
	fmt.Println("n >", n)
	fmt.Println("n >", n)
	
}




//pointers

func main1() {
	i:=34
	var p =&i
	e:=*p
	
	var y =&i	
	fmt.Println("p=  >", p)
	fmt.Println("e=  >", e)
	fmt.Println("i=  >", i)
	fmt.Println("y=  >",y)
	fmt.Println("y=  >",*y)



	n:=433
	m:=&n
	fmt.Println("m=  >", m)
	fmt.Println("*M=  >",*m)




}


//data structures Map
func map1(){
	var m map[string]int
	m=map[string]int{
		"h1":34,
		"h2":34,
		"h3":24,
	}
	m["k1"]=3
	m["g2"]=323
	n:=	m["g2"]
	m["d8"]=34
//	delete(m , "g2")
//	k2,ok:=m["g2"]
	_,ok:=m["k3"]
	fmt.Println("m  >", m)
	fmt.Println("m  >", len(m))
	fmt.Println("g2  >",n )
//	fmt.Println("k1,ok  >",ok,k2)
	fmt.Println("k1,ok  >",ok)

	




		
}
func map2(){
	m:=make(map[string]string)
	m["b1"]="hossein"
	m["b2"]="ali"
	m["b3"]="nikoo"
	c:=	m["b3"]
	delete(m,"b1")
	_,ok:=m["b1"]
	fmt.Println("map=  >",m)
	fmt.Println("b3=  >",c)
	fmt.Println("ok  >",ok)

	

	






}