package main

import (
	"fmt"
	"os"
	u "anxiel_cube/utils"
)


func main() {
	fmt.Printf("Your first argument id %s\n", os.Args[1])
	fmt.Println("Hello Raymond! We successfully setup golang!")
	fmt.Println(u.Process_input(os.Args[1]))
}

