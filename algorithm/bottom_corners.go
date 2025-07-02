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
			position = 0
		case 3:
			c = c.U().R_p().U_p().R()
			m.Commands = append(m.Commands, "B'")
			c = sortCorner(c)
			position = 0
		case 4:
			c = c.B()
			m.Commands = append(m.Commands, "B'")
			c = sortCorner(c)
			position = 0
		case 5: // rotate 90, do case 1 without the B' in one rotation, rotate -90 perform case 1 recursively including B'
			c = c.B_p()
			m.Commands = append(m.Commands, "B'")
			c = sortCorner(c)
			position = 0
		case 6:
			c = c.B_p()
			m.Commands = append(m.Commands, "B'")
			c = sortCorner(c)
			position = 0
		case 7: // rotate -180, perform case 1 without B' in one rotation; perform case 8
			c = c.B_p()
			m.Commands = append(m.Commands, "B'")
			c = sortCorner(c)
			position = 0
		case 8: // rotate  2B then perform case 1 recursively without B'

			c = c.B_p()
			m.Commands = append(m.Commands, "B'")
			c = sortCorner(c)
			position = 0
		}
	}
	return c
}
