//Interfaces

package main

import "fmt"

type geometry interface{
	area() float64
	perim() float64
}

type rect struct{
	w,l float64
}

type squa struct{
	l float64
}

func (s squa) area() float64{
	return s.l*s.l
}

func( s squa) perim() float64{
	return 4*s.l
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	d1:=squa{l:9}
	measure(d1)
}