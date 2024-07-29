package main
import "fmt"
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base //embeding name without field
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// We can access the base's fields directly on `co`,
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Alternatively, we can spell out the full path using the embedded type name.
	fmt.Println("also num:", co.base.num)

	// Since `container` embeds `base`, the methods of
	// `base` also become methods of a `container`
	fmt.Println("describe:", co.describe())
}