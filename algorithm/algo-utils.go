package algorithm

import m "anxiel_cube/models"

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
