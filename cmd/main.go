package main

import (
	"fmt"

	"github.com/abnt713/testtest/dude"
)

func main() {
	x := dude.SayMyName{Name: "Dude"}
	fmt.Println(x.SayIt())
}
