package main

import (
	"fmt"
	"myapp/data"
	"myapp/orented"
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
	Array1()// data structures >> Array and slice >> Array
	slice1()//data structures >> Array and slice >> Slice
	struct1()//data structures >> Array and slice >> struct
	man()//loop for rage
	ifo()//for and els
	stat(3)//switch statcment
	orented.Ory()

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




//lern slice and Array
//Array

func Array1(){

	var b [3]int
	b[2]=3
	d:=b[2]
	j:= [10]int{1,2,23,3,4,4,5,}
	o:=len(j)
	fmt.Println("b=  >",b)
	fmt.Println("d=  >",d)
	fmt.Println("j= >>",j)
	fmt.Println("o= >>",o)

	







}



//Slice

func slice1(){
	e:=make([]int,3)
	e[0]=23
	e[2]=34
	e[1]=43
	n:=[]int{1,2,3,4,5,5,6,3,3,3,3,3}
	n[2]=4
	h:=n[2]
	k:=len(e)
	l:=len(n)
	i:=n[:]
	q:=append(n[0:5],n[5+1:]...)
	u:=cap(n)
	m:=len(n)
	v:=cap(e)
	x:=len(e)
	n=append(n,400)
	n=append(n, 400,300)

	for key,value :=range n{
		fmt.Println(key,value)
	}


	for key,value :=range n{
		fmt.Println(key,value)
	}

	fmt.Print("===================================")

	for i:=0 ;i<len(n);i++{
		fmt.Println(i,n[i])
	}

	for i:=0;i<len(n);i++{
		fmt.Println(i,n[i])
	}

	fmt.Print("===================================")

	for key,value := range n{
		fmt.Println(key,value)
	}

	fmt.Print("===================================")

		for i:=0 ;i<len(n);i++{
			fmt.Println(i,n[i])
		}

		fmt.Print("===================================")



	fmt.Println("e=  >",e)
	fmt.Println("n=  >",n)
	fmt.Println("n=  >",n)
	fmt.Println("h=  >",h)
	fmt.Println("k,l=  >",k,l)
	fmt.Println("i,q=  >",i,q)
	fmt.Println("u,m=  >",u,m)
	fmt.Println("v,x=  >",v,x)


}


//struct

type person struct{
	name string
	last_name string
	age int
}

func struct1()  {
	p:=person{age: 14,name: "hose",last_name: "dlsjf"}
	p.age=30
	p.last_name="soltan"
	p.name="hossein"
	fmt.Println(p)



	
}

func man(){
	n:=[]int{1,2,2,3,3,4,45,5,6,7,88,}


	for key,value :=range n{
		fmt.Println(key,value)
	}


	for i:=0 ; i<len(n);i++{
		fmt.Println(i,n[i])
	}


	for i:=0;i<len(n);i++{
		fmt.Println(i,n[i])
	}


	for i:=0 ; i<len(n) ; i++{
		fmt.Println(i,n[i])
	}


}

func ifo(){
	n:=45
	s:=2
	if
	 n<s || s!=n{
		fmt.Print("hello")
	}else
	 {fmt.Println("good morinig")
	}


	s=32
	m:=1
	if m!=s || m==m{
		fmt.Println("\niam ali")
	}else if s==2{
		fmt.Println("are you happy")
	}

	m=5
	if s%m==0{
		fmt.Println("ok")
	}else if s%m==1{
		fmt.Println("no1")
	}else if s%m==2{
		fmt.Println("no2")
	}else if s%m==3{
		fmt.Println("no3")
	}else if s%m==4{
		fmt.Println("no4")
	}else if s%m==5{
		fmt.Println("no5")
	}else if s%m==6{
		fmt.Println("no6")
	}else if s%m==7{
		fmt.Println("no7")
	}else if s%m==8{
		fmt.Println("no8")
	}else if s%m==9{
		fmt.Println("no9")
	}
	

	
}


//switch statcment

func stat(day int){
	switch day{
	case 1:fmt.Println("sunday")
	case 2:fmt.Println("monday")
	case 3:fmt.Println("Tuesday")
	case 4:fmt.Println("Wednesday")
	case 5:fmt.Println("Thursday")
	case 6:fmt.Println("Friday")
	default:fmt.Println("not day")

	}



	
	switch day{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("not int")
	}



	y:=12
	u:=9
	b:=y&u
	switch b{
	case 0:
		fmt.Println("Playable")
	case 1,2,3,4,5,6,7,8,9:
		fmt.Println("not Playable")
	}




}


//object orented


