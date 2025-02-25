package structembedding

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func StructEmbedding() {

	// b := base{num: 5}
	// fmt.Println(b.describe())

	co := container{
		base: base{
			num: 44,
		},
		str: "my name is riteshreg",
	}

	fmt.Println(co.describe())

}
