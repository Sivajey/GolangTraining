package main

import (
	"fmt"
	"github.com/goestoeleven/GolangTraining/05_scope/13_capitalization/characters"
)

func main() {
	fmt.Println(characters.Sherrif())
	fmt.Println(characters.SherrifName) // valid - can access the variable
}