package models

import (
	"bufio"
	"os"
	"strings"
)

func PopulateRubikCube(fileName string) (*Rubik_cube, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var faces [][]string
	scanner := bufio.NewScanner(file)
	var currentFace []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if len(currentFace) > 0 {
				faces = append(faces, currentFace)
				currentFace = nil
			}
			continue
		}
		currentFace = append(currentFace, strings.Fields(line)...)
	}
	if len(currentFace) > 0 {
		faces = append(faces, currentFace)
	}

	if len(faces) != 6 {
		return nil, os.ErrInvalid
	}

	rubikCube := &Rubik_cube{
		Front_face: faces[0][4],
		Top_face:   faces[4][4],
		Top_Layer: Top_Layer{
			Left_front: Corner_cubit{Top_face: faces[4][0], Front_face: faces[0][0], Right_face: faces[1][0]},
			Mid_front:  Mid_cubit{Top_face: faces[4][1], Front_face: faces[0][1]},
			Right_front: Corner_cubit{Top_face: faces[4][2], Front_face: faces[0][2], Right_face: faces[1][2]},
			Center_cubit: faces[4][4],
		},
		Middle_Layer: Middle_Layer{
			Left_front:  Mid_cubit{Front_face: faces[0][3], Top_face: faces[4][3]},
			Right_front: Mid_cubit{Front_face: faces[0][5], Top_face: faces[4][5]},
		},
		Bottom_Layer: Bottom_Layer{
			Left_front: Corner_cubit{Front_face: faces[0][6], Top_face: faces[5][0], Right_face: faces[1][6]},
			Mid_front:  Mid_cubit{Front_face: faces[0][7], Top_face: faces[5][1]},
			Right_front: Corner_cubit{Front_face: faces[0][8], Top_face: faces[5][2], Right_face: faces[1][8]},
			Center_cubit: faces[5][4],
		},
	}

	return rubikCube, nil
}
