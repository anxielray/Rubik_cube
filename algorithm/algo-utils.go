package algorithm

import (
	m "anxiel_cube/models"
)

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

// CheckCross function takes a rubic cube
// chceks if the bottom layer has a bottom cross
func CheckCross(c *m.Rubik_cube, refColor string) bool {
	return checkTLCross(c, refColor) && checkRMLCross(c, refColor) && checkLMLCross(c, refColor) && checkBLCross(c, refColor)
}

// locate corner function takes a cube
// locates the corner cubit and alligns the cubit correlclty
// returns a number as detailed in the notes.txt file
func bottomLocateCorner(c *m.Rubik_cube) int {
	var (
		p1 = c.Top_Layer.Left_front
		p2 = c.Top_Layer.Left_back
		p3 = c.Top_Layer.Right_front
		p4 = c.Top_Layer.Right_back
		p5 = c.Bottom_Layer.Left_front
		p6 = c.Bottom_Layer.Left_back
		p7 = c.Bottom_Layer.Right_front
		p8 = c.Bottom_Layer.Right_back
	)

	if comboB(c.Top_Layer.Mid_front.Top_face, c.Top_Layer.Mid_front.Front_face, c.Middle_Layer.Left_front.Top_face, p1) {
		return 1
	} else if comboB(c.Top_Layer.Mid_front.Top_face, c.Top_Layer.Mid_front.Front_face, c.Middle_Layer.Left_front.Top_face, p2) {
		return 2
	} else if comboB(c.Top_Layer.Mid_front.Top_face, c.Top_Layer.Mid_front.Front_face, c.Middle_Layer.Left_front.Top_face, p3) {
		return 3
	} else if comboB(c.Top_Layer.Mid_front.Top_face, c.Top_Layer.Mid_front.Front_face, c.Middle_Layer.Left_front.Top_face, p4) {
		return 4
	} else if comboB(c.Top_Layer.Mid_front.Top_face, c.Top_Layer.Mid_front.Front_face, c.Middle_Layer.Left_front.Top_face, p5) {
		return 5
	} else if comboB(c.Top_Layer.Mid_front.Top_face, c.Top_Layer.Mid_front.Front_face, c.Middle_Layer.Left_front.Top_face, p6) {
		return 6
	} else if comboB(c.Top_Layer.Mid_front.Top_face, c.Top_Layer.Mid_front.Front_face, c.Middle_Layer.Left_front.Top_face, p7) {
		return 7
	} else if comboB(c.Top_Layer.Mid_front.Top_face, c.Top_Layer.Mid_front.Front_face, c.Middle_Layer.Left_front.Top_face, p8) {
		return 8
	}
	return 0
}

// correct combo fucntion is a function that confirms if the correct combination of corner colors have been met for the corner
func comboB(ct, cf, cs string, cubit m.Corner_cubit) bool {
	var (
		T = cubit.Top_face
		F = cubit.Front_face
		S = cubit.Side_face
	)

	if T == ct && F == cf && S == cs {
		return true
	}

	if T == ct && S == cf && F == cs {
		return true
	}

	if S == ct && T == cf && F == cs {
		return true
	}

	if S == ct && F == cf && T == cs {
		return true
	}

	if F == ct && T == cf && S == cs {
		return true
	}

	if F == ct && S == cf && T == cs {
		return true
	}

	return false
}

// corner allign is a functiont that checks if the corner cubit is in it's correct position
func cornerAllign(c *m.Rubik_cube) bool {
	return (c.Top_Layer.Left_front.Top_face == c.Top_Layer.Mid_front.Top_face) && c.Top_Layer.Left_front.Side_face == c.Middle_Layer.Left_front.Top_face && c.Top_Layer.Left_front.Front_face == c.Middle_Layer.Center_front_cubit
}

func cornersAllign(c *m.Rubik_cube) bool {
	if c.Top_Layer.Left_front.Top_face == c.Top_Layer.Mid_back.Top_face && c.Top_Layer.Left_front.Side_face == c.Middle_Layer.Left_front.Top_face && c.Top_Layer.Left_front.Front_face == c.Top_Layer.Mid_front.Front_face {
		if c.Top_Layer.Right_front.Top_face == c.Top_Layer.Mid_front.Top_face && c.Top_Layer.Right_front.Side_face == c.Middle_Layer.Right_front.Top_face && c.Top_Layer.Right_front.Front_face == c.Middle_Layer.Right_front.Front_face {
			if c.Bottom_Layer.Right_front.Top_face == c.Bottom_Layer.Mid_front.Top_face && c.Bottom_Layer.Right_front.Front_face == c.Bottom_Layer.Mid_front.Front_face && c.Bottom_Layer.Right_front.Side_face == c.Middle_Layer.Right_front.Top_face {
				if c.Bottom_Layer.Left_front.Top_face == c.Bottom_Layer.Mid_front.Top_face && c.Bottom_Layer.Left_front.Front_face == c.Bottom_Layer.Mid_front.Front_face && c.Bottom_Layer.Left_front.Side_face == c.Middle_Layer.Left_front.Top_face {
					return true
				}
			}
		}
	}
	return false
}

// sort corners will run some commands to rotate the corner cubit trying to allign it properly with its supposed position
func sortCorner(c *m.Rubik_cube) *m.Rubik_cube {
	for !cornerAllign(c) {
		c = c.L().U_p().L_p().U()
		m.Commands = append(m.Commands, "L")
		m.Commands = append(m.Commands, "U'")
		m.Commands = append(m.Commands, "L'")
		m.Commands = append(m.Commands, "U")
	}
	c = c.F()
	m.Commands = append(m.Commands, "F")
	return c
}

// locate corner function takes a cube
// locates the corner cubit and alligns the cubit correlclty
// returns a number as detailed in the notes.txt file
func midLocateEdge(c *m.Rubik_cube) int {
	var (
		p1 = c.Middle_Layer.Left_front
		p2 = c.Middle_Layer.Right_front
		p3 = c.Middle_Layer.Right_back
		p4 = c.Middle_Layer.Left_back
		p5 = c.Top_Layer.Mid_front
		p6 = c.Top_Layer.Right_mid
		p7 = c.Top_Layer.Mid_back
		p8 = c.Top_Layer.Left_mid
	)

	if comboM(c.Middle_Layer.Center_left_cubit, c.Middle_Layer.Center_front_cubit, p1) {
		return 1
	} else if comboM(c.Middle_Layer.Center_left_cubit, c.Middle_Layer.Center_front_cubit, p2) {
		return 2
	} else if comboM(c.Middle_Layer.Center_left_cubit, c.Middle_Layer.Center_front_cubit, p3) {
		return 3
	} else if comboM(c.Middle_Layer.Center_left_cubit, c.Middle_Layer.Center_front_cubit, p4) {
		return 4
	} else if comboM(c.Middle_Layer.Center_left_cubit, c.Middle_Layer.Center_front_cubit, p5) {
		return 5
	} else if comboM(c.Middle_Layer.Center_left_cubit, c.Middle_Layer.Center_front_cubit, p6) {
		return 6
	} else if comboM(c.Middle_Layer.Center_left_cubit, c.Middle_Layer.Center_front_cubit, p7) {
		return 7
	} else if comboM(c.Middle_Layer.Center_left_cubit, c.Middle_Layer.Center_front_cubit, p8) {
		return 8
	}
	return 0
}

// function to find a  matching combination of the edge cubits
func comboM(et, ef string, cubit m.Mid_cubit) bool {
	var (
		T = cubit.Top_face
		F = cubit.Front_face
	)

	if T == et && F == ef {
		return true
	}

	if F == et && T == ef {
		return true
	}

	return false
}