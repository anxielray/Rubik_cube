package algorithm

import (
	m "anxiel_cube/models"
	"fmt"
)

// ====( This is a function to form the cross of starting face of the cube )=====
func BottomCross(cube *m.Rubik_cube) *m.Rubik_cube {

	referrenceFace := cube.Front_face
	for i := range 4 {
		if cube.Top_Layer.Mid_front.Front_face == referrenceFace {
			if i > 0 {
				if i > 1 {
					m.Commands = append(m.Commands, fmt.Sprintf("%dU", i))
				} else {
					m.Commands = append(m.Commands, "U")
				}
			}
			break
		}
		cube = cube.U()
	}

	for i := range 4 {
		if cube.Middle_Layer.Right_front.Front_face == referrenceFace {
			if i > 0 {
				if i > 1 {
					m.Commands = append(m.Commands, fmt.Sprintf("%dR", i))
				} else {
					m.Commands = append(m.Commands, "R")
				}
				cube = BottomCross(cube)
			}
			break
		}
		cube = cube.R()
	}

	for i := range 4 {
		if cube.Middle_Layer.Left_front.Front_face == referrenceFace {
			if i > 0 {
				if i > 1 {
					m.Commands = append(m.Commands, fmt.Sprintf("%dL", i))
				} else {
					m.Commands = append(m.Commands, "L")
				}
				cube = BottomCross(cube)
			}
			break
		}
		cube = cube.L()

	}

	for i := range 4 {
		if cube.Bottom_Layer.Mid_front.Front_face == referrenceFace {
			if i > 0 {
				if i > 1 {
					m.Commands = append(m.Commands, fmt.Sprintf("%dD", i))
				} else {
					m.Commands = append(m.Commands, "D")
				}
				cube = BottomCross(cube)
			}
			break
		}
		cube = cube.D()

	}
	// if m.Commands[len(m.Commands)-1] != "U" {

	// }
	return cube

}

func checkMidCompatability(cube *m.Rubik_cube) {
	//check if the top middle face is the base color and check if the color is compatible to the corresponding center cube territory
	for cube.Top_Layer.Mid_front.Front_face == cube.Front_face {
		if cube.Top_Layer.Mid_front.Top_face == cube.Top_Layer.Center_cubit {
			break
		} else {

		}
	}
}
