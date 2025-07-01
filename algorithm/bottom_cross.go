package algorithm

import (
	m "anxiel_cube/models"
	"fmt"
	"os"
)

// var counter int
var markedMap = map[string]string{
	"U": "!",
	"R": "!",
	"D": "!",
	"L": "!",
}

//	Bottom cross function takes a virtual cube prolem,
//
// sorts the bottom face to have the bottom cross
// returns the sorted cube
func BottomCross(cube *m.Rubik_cube) *m.Rubik_cube {

	var referrenceFace = cube.Front_face

	if CheckCross(cube, referrenceFace) {
		return alignTheCube(cube)
	}

	// check for the top section of the cross
	for i := range 4 {
		if cube.Top_Layer.Mid_front.Front_face == referrenceFace {
			if i > 0 {
				// record the command used in the array of commands
				if i > 1 {
					m.Commands = append(m.Commands, fmt.Sprintf("%dU", i))
				} else {
					m.Commands = append(m.Commands, "U")
				}

				// if the cross is made at this point, return from the function
				if CheckCross(cube, referrenceFace) {
					return alignTheCube(cube)
				}
			}
			markedMap["U"] = "x"
			break
		}
		cube = cube.U()
	}

	// check for the right section of the cross
	for i := range 4 {

		if cube.Middle_Layer.Right_front.Front_face == referrenceFace {
			if i > 0 {
				if i > 1 {
					m.Commands = append(m.Commands, fmt.Sprintf("%dR", i))
				} else {
					m.Commands = append(m.Commands, "R")
				}
				if CheckCross(cube, referrenceFace) {
					return alignTheCube(cube)
				}
				if !checkTLCross(cube, referrenceFace) {
					cube = BottomCross(cube)
				}
			}
			markedMap["R"] = "x"
			break
		}
		cube = cube.R()
	}

	// check for the left section of the cross
	for i := range 4 {
		if cube.Middle_Layer.Left_front.Front_face == referrenceFace {
			if i > 0 {
				if i > 1 {
					m.Commands = append(m.Commands, fmt.Sprintf("%dL", i))
				} else {
					m.Commands = append(m.Commands, "L")
				}
				if CheckCross(cube, referrenceFace) {
					return alignTheCube(cube)
				}
			}
			markedMap["L"] = "x"
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
				if CheckCross(cube, referrenceFace) {
					return alignTheCube(cube)
				}
			}
			markedMap["D"] = "x"
			break
		}
		cube = cube.D()

	}

	if markedMap["U"] == "!" {
		var ok bool
		ok, cube = InvertedTopCross(cube, referrenceFace)
		if ok {
			m.Commands = append(m.Commands, "2U")
			m.Commands = append(m.Commands, "B")
			m.Commands = append(m.Commands, "L")
			m.Commands = append(m.Commands, "U'")
			m.Commands = append(m.Commands, "L'")
			markedMap["U"] = "x"
		}
	}
	return alignTheCube(cube)
}

func InvertedTopCross(c *m.Rubik_cube, refColor string) (bool, *m.Rubik_cube) {
	// 2U B L U'
	var c_cp m.Rubik_cube
	c_cp = *c
	c_cp = *c_cp.U()
	c_cp = *c_cp.U()
	c_cp = *c_cp.B()
	c_cp = *c_cp.L()
	c_cp = *c_cp.U_p()
	c_cp = *c_cp.L_p()
	var decicion = c_cp.Top_Layer.Mid_front.Front_face == refColor
	c = &c_cp
	return decicion, c
}

// function to align the solved bottom layer to the corresponding color face on the sides
func alignTheCube(cube *m.Rubik_cube) *m.Rubik_cube {
	c := cube
	for i := range 4 {
		if c.Top_Layer.Mid_front.Top_face == c.Top_Layer.Center_cubit {
			fmt.Println(":hey")

			// after checking the top, the rest should be alligned
			if c.Middle_Layer.Right_front.Top_face == c.Middle_Layer.Center_right_cubit {

				// check if the left
				if c.Middle_Layer.Left_front.Top_face == c.Middle_Layer.Center_left_cubit {
					// check if the bottom
					if c.Bottom_Layer.Mid_front.Top_face == c.Bottom_Layer.Center_cubit {
						if i > 0 {
							if i > 1 {
								m.Commands = append(m.Commands, fmt.Sprintf("%dF", i))
							} else {
								m.Commands = append(m.Commands, "F")
							}
						}
						fmt.Println("all the faces alligned")
						break
					} else {
						println("the bottom face and the cross do not allign")
						os.Exit(0)
					}
				} else {
					println("the left face and the cross do not allign")
					os.Exit(0)
				}
			} else {
				println("the right face and the cross do not allign")
				os.Exit(0)
			}
		}
		c = c.F()
	}
	cube = c
	return cube
}
