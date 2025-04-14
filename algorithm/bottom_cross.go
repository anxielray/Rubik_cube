package algorithm

import (
	m "anxiel_cube/models"
	"fmt"
)

// var counter int
var markedMap = map[string]string{
	"U": "!",
	"R": "!",
	"D": "!",
	"L": "!",
}

// ====( This is a function to form the cross of starting face of the cube )=====
func BottomCross(cube *m.Rubik_cube) *m.Rubik_cube {

	var referrenceFace = cube.Front_face
	for i := range 4 {
		if cube.Top_Layer.Mid_front.Front_face == referrenceFace {
			if i > 0 {
				if i > 1 {
					m.Commands = append(m.Commands, fmt.Sprintf("%dU", i))
				} else {
					m.Commands = append(m.Commands, "U")
				}
			}
			markedMap["U"] = "x"
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
				if !checkTLCross(cube, referrenceFace) {
					cube = BottomCross(cube)
				}
			}
			markedMap["R"] = "x"
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
				if !checkTLCross(cube, referrenceFace) || !checkRMLCross(cube, referrenceFace) {
					cube = BottomCross(cube)
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
				if !checkTLCross(cube, referrenceFace) || !checkRMLCross(cube, referrenceFace) || !checkLMLCross(cube, referrenceFace) {
					cube = BottomCross(cube)
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

	// if !checkBLCross(cube, referrenceFace) {
	// 	cube = BottomCross(cube)
	// }
	fmt.Println(markedMap)
	/*if CheckCross(cube, referrenceFace) {*/return cube//}
	// return nil
}

func checkTLCross(c *m.Rubik_cube, refColor string) bool {
	return c.Top_Layer.Mid_front.Front_face == refColor
}

func checkRMLCross(c *m.Rubik_cube, refColor string) bool {
	return c.Middle_Layer.Right_front.Front_face == refColor
}

func checkLMLCross(c *m.Rubik_cube, refColor string) bool {
	return c.Bottom_Layer.Mid_front.Front_face == refColor
}

func checkBLCross(c *m.Rubik_cube, refColor string) bool {
	return c.Middle_Layer.Left_front.Front_face == refColor
}

func CheckCross(c *m.Rubik_cube, refColor string) bool {
	return checkTLCross(c, refColor) && checkRMLCross(c , refColor) && checkLMLCross(c, refColor) && checkBLCross(c, refColor)
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