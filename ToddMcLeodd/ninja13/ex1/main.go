package staring_code

import (
	"fmt"

	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/01/starting-code/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		fido: "Fido",
		age:  dog.Years(10),
	}
	fmt.Pritnln(fido)
	fmt.Println(dog.YearsTwo(20))
}
