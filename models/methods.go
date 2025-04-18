package models

import (
	"bufio"
	"os"
	"strings"
)

// ===( This is a function that will populate the file input into the virtual {3x3} rubik's cube )====
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
		Top_Layer: Top_Layer{
			Left_front:   Corner_cubit{Top_face: faces[4][6], Front_face: faces[0][0], Side_face: faces[3][2]},
			Left_mid:     Mid_cubit{Top_face: faces[4][3], Front_face: faces[3][1]},
			Left_back:    Corner_cubit{Top_face: faces[4][0], Front_face: faces[2][2], Side_face: faces[3][0]},
			Mid_front:    Mid_cubit{Top_face: faces[4][7], Front_face: faces[0][1]},
			Mid_back:     Mid_cubit{Top_face: faces[4][1], Front_face: faces[2][1]},
			Right_front:  Corner_cubit{Top_face: faces[4][8], Front_face: faces[0][2], Side_face: faces[1][0]},
			Right_mid:    Mid_cubit{Top_face: faces[4][5], Front_face: faces[1][1]},
			Right_back:   Corner_cubit{Top_face: faces[4][2], Front_face: faces[1][2], Side_face: faces[2][0]},
			Center_cubit: faces[4][4],
		},
		Middle_Layer: Middle_Layer{
			Left_front:         Mid_cubit{Front_face: faces[0][3], Top_face: faces[3][5]},
			Left_back:          Mid_cubit{Front_face: faces[2][5], Top_face: faces[3][3]},
			Right_front:        Mid_cubit{Front_face: faces[0][5], Top_face: faces[1][3]},
			Right_back:         Mid_cubit{Front_face: faces[1][5], Top_face: faces[2][3]},
			Center_front_cubit: faces[0][4],
			Center_back_cubit:  faces[2][4],
			Center_right_cubit: faces[1][4],
			Center_left_cubit:  faces[3][4],
		},
		Bottom_Layer: Bottom_Layer{
			Left_front:   Corner_cubit{Top_face: faces[5][0], Front_face: faces[0][6], Side_face: faces[3][8]},
			Left_mid:     Mid_cubit{Top_face: faces[5][3], Front_face: faces[3][7]},
			Left_back:    Corner_cubit{Top_face: faces[5][6], Front_face: faces[3][6], Side_face: faces[2][8]},
			Mid_front:    Mid_cubit{Top_face: faces[5][1], Front_face: faces[0][7]},
			Mid_back:     Mid_cubit{Top_face: faces[5][7], Front_face: faces[2][7]},
			Right_front:  Corner_cubit{Top_face: faces[5][2], Front_face: faces[0][8], Side_face: faces[1][6]},
			Right_mid:    Mid_cubit{Top_face: faces[5][5], Front_face: faces[1][7]},
			Right_back:   Corner_cubit{Top_face: faces[5][8], Front_face: faces[1][8], Side_face: faces[2][6]},
			Center_cubit: faces[5][4],
		},
	}

	return rubikCube, nil
}

//===( This function rotates the cube in the plane axis through a degree of angle rotation )====
func Rotation(c *Rubik_cube, plane string, angle int) *Rubik_cube {

	if plane == "x" {
		if angle == 90 || angle == -270 {
			c = c.U()
			c = c.Mx()
			c = c.D_p()
		} else if angle == 180 || angle == -180 {
			for range 2 {
				c = c.U()
				c = c.Mx()
				c = c.D_p()
			}
		} else if angle == 270 || angle == -90 {
			for range 3 {
				c = c.U()
				c = c.Mx()
				c = c.D_p()
			}
		}
	} else if plane == "y" {
		if angle == 90 || angle == -270 {
			c = c.R()
			c = c.My()
			c = c.L_p()
		} else if angle == 180 || angle == -180 {
			for range 2 {
				c = c.R()
				c = c.My()
				c = c.L_p()
			}
		} else if angle == 270 || angle == -90 {
			for range 3 {
				c = c.R()
				c = c.My()
				c = c.L_p()
			}
		}
	}
	return c
}
