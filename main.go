package main

import (
	"fmt"
	AssembledCar "main/assembledcar"
)

func main() {
	car := AssembledCar.NewCar()
	fmt.Printf("%+v\n", car)
}
