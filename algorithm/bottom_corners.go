package algorithm

import (
	m "anxiel_cube/models"
	"fmt"
)

// Bottom corners takes a cube
// sorts the corners of the front face
// returns a sorted cube(one face sorted)
func BottomCorners(c *m.Rubik_cube) *m.Rubik_cube {
	c = alignTheCube(c)
	// locate the target corner cubit
	for !cornersAllign(c) {
		position := locateCorner(c)
		fmt.Println(position)
		switch position {
		case 1, 2:
			c = sortCorner(c)
		case 3:
			c = c.U().R_p().U_p().R()
			m.Commands = append(m.Commands, "U")
			m.Commands = append(m.Commands, "R'")
			m.Commands = append(m.Commands, "U'")
			m.Commands = append(m.Commands, "R")
			c = sortCorner(c)
		case 4:
			c = c.B()
			m.Commands = append(m.Commands, "B")
			c = sortCorner(c)
		case 5:
			c = m.Rotation(c, "y", -90)
			c = m.Rotation(c, "x", 90)
			c = m.Rotation(c, "y", 90)
			c = c.L().U_p().L_p().U()
			c = m.Rotation(c, "y", -90)
			c = m.Rotation(c, "x", -90)
			c = m.Rotation(c, "y", 90)

			m.Commands = append(m.Commands, "L")
			m.Commands = append(m.Commands, "U'")
			m.Commands = append(m.Commands, "L'")
			m.Commands = append(m.Commands, "U")
			c = c.B_p()
			m.Commands = append(m.Commands, "B'")
			c = sortCorner(c)
		case 6:
			c = c.B_p()
			m.Commands = append(m.Commands, "B'")
			c = sortCorner(c)
		case 7:
			c = m.Rotation(c, "y", 90)
			c = m.Rotation(c, "x", 90)
			c = m.Rotation(c, "y", -90)
			c = c.U().R_p().U_p().R()
			c = m.Rotation(c, "y", 90)
			c = m.Rotation(c, "x", -90)
			c = m.Rotation(c, "y", -90)
			c = c.B().B()
			m.Commands = append(m.Commands, "U")
			m.Commands = append(m.Commands, "R'")
			m.Commands = append(m.Commands, "U'")
			m.Commands = append(m.Commands, "R")
			m.Commands = append(m.Commands, "2B")
			c = sortCorner(c)
		case 8:
			c = c.B().B()
			m.Commands = append(m.Commands, "2B")
			c = sortCorner(c)
		}
		if cornersAllign(c) {
			break
		}
	}
	return alignTheCube(c)
}
