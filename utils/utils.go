package utils


import (
//	m "anxiel_cube/models"
	"fmt"
	"os"
	"bufio"
)


//==( This is a function to process the input file content )===
func Process_input(file_name string) (string, error) {

	f, err := os.Open(file_name)
	if err != nil {
		return "", fmt.Errorf("processing file input: {%v}", err)
	}
	var combined_line string
	
	new_scanner := bufio.NewScanner(f)
	for new_scanner.Scan() {
		line := new_scanner.Text()
		combined_line += line+"\t"
	}

	return combined_line, nil
}


//==( This is a function to load the input into the Rubik's cube object
//func Populate_puzzle(input string) *m.Rubik_cube {

//	var tmp_r  m.Rubik_cube

//	tmp_r
//	return *tmp_r
//}
