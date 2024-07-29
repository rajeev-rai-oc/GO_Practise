package main

import "fmt"

func inc( a int ) func() int{
	i:=0
	return func() int {
		i++
		return (a + i)
	}

}

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr * int) {
    *iptr = 0
}

type person struct {
	name string
	age  int
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}


func main(){
	fmt.Println("Hello World ")
	ob:= inc(3)
	fmt.Println(ob())

	test:=1
	zeroval(test)
	fmt.Println(test)

	zeroptr(&test)
	fmt.Println(test)

	s:="hello"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Access struct fields with a dot.
	s1 := person{name: "Sean", age: 50}
	fmt.Println(s1)

	s2:=&s1
	s2.name="Seany"
	fmt.Println(s2)

	r:=rect { width:20, height:10}
	fmt.Println("Area:",r.area())
	fmt.Println("Perimeter:",r.perim())

}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found t")
	} else if r == 'h' {
		fmt.Println("h")
	}
}