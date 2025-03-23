package main

import (
	"fmt"
	"os"
	// u "anxiel_cube/utils"
	m "anxiel_cube/models"
	algo "anxiel_cube/algorithm"
)


func main() {
	fmt.Println("========This is the Solution======")
	pzzle, _ := m.PopulateRubikCube(os.Args[1])
	fmt.Println(algo.SolveRC(pzzle))
}

