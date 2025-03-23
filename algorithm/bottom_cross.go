package algorithm

import (
	m "anxiel_cube/models"
	"fmt"
)

var counter int
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
	/*
		There are 14 insances where missing values can occur
		Start with the planB to B{4loop{U}} this will enable the program to venture into the back middle-cubits
		If planBU is not sufficient planCU 2{U B R' U' R}, {U'B R' U' R}, {2U, B R' U' R}
		And if planC is not sufficient, then finally planDU 2{U B L U' L'}, {U'B L U' L'}, {2U, B L U' L'}
		Performing these operations at each interval before performing the 4loop{U} will ensure that
		all middle cubits present are tested without necessarily interfering with the other
		arranged middle-cubits of the other layers. Alll the layers should have been covered by
		now meaning that the referrence mid-cubt must have been found already
		============================================================
		To avoid writing too many commands, mind writing the function to  turn the cube in different planes to reuse the same functions
		and methods
	*/

	if markedMap["U"] == "!" {
		if counter > 0 {
			counter += 1
		}
		if counter == 0 {
			
			for i := range 4 {
				for x := range 4 {
					if !checkTLCross(cube, referrenceFace) {
						cube = cube.B()
						fmt.Println(cube.Top_Layer)
					} else if checkTLCross(cube, referrenceFace) {
						if i > 0 {
							if i > 1 {
								m.Commands = append(m.Commands, fmt.Sprintf("%dB", i))
								if x > 0 {
									if x > 1 {
										m.Commands = append(m.Commands, fmt.Sprintf("%dU", x))
									} else {
										m.Commands = append(m.Commands, "U")
									}
								}
							} else {
								m.Commands = append(m.Commands, "B")
								if x > 0 {
									if x > 1 {
										m.Commands = append(m.Commands, fmt.Sprintf("%dU", x))
									} else {
										m.Commands = append(m.Commands, "U")
									}
								}
							}
						}
						markedMap["U"] = "x"
						break
					}
					cube = cube.U()
				}
			}
			counter = 1
		}
		if !checkTLCross(cube, referrenceFace) {

			planBU := func(cube *m.Rubik_cube, counter int) *m.Rubik_cube {
				if counter == 1 {
					cube = cube.U()
					cube = cube.B()
					cube = cube.R_p()
					cube = cube.U_p()
					cube = cube.R()
				} else if counter == 2 {
					for range counter {
						cube = cube.U()
					}
					cube = cube.B()
					cube = cube.R_p()
					cube = cube.U_p()
					cube = cube.R()
				} else if counter == 3 {
					cube = cube.U_p()
					cube = cube.B()
					cube = cube.R_p()
					cube = cube.U_p()
					cube = cube.R()
				} else if counter == 4 {
					cube = cube.U()
					cube = cube.U()
					cube = cube.B()
					cube = cube.R_p()
					cube = cube.R()
				}
				return cube
			}
			if counter > 0 && counter < 5 {
				cube = BottomCross(planBU(cube, counter))
			}

			planCU := func(c *m.Rubik_cube, counter int) *m.Rubik_cube {
				if counter == 5 { // 2{U B L U' L'}
					cube = cube.U()
					cube = cube.B()
					cube = cube.L()
					cube = cube.U_p()
					cube = cube.L_p()
				} else if counter == 6 {
					cube = cube.U()
					cube = cube.U()
					cube = cube.B()
					cube = cube.L()
					cube = cube.U_p()
					cube = cube.L_p()
				} else if counter == 7 {
					cube = cube.U_p()
					cube = cube.B()
					cube = cube.L()
					cube = cube.U_p()
					cube = cube.L_p()
				} else if counter == 8 {
					cube = cube.U_p()
					cube = cube.B()
					cube = cube.L()
					cube = cube.U_p()
					cube = cube.L_p()
				}
				return cube
			}
			if !checkTLCross(cube, referrenceFace) {
				cube = BottomCross(planCU(cube, counter))
			}
		}
	}

	if !checkBLCross(cube, referrenceFace) {
		cube = BottomCross(cube)
	}
	fmt.Println(markedMap)
	return cube
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
