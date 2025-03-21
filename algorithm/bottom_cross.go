package algorithm

import(
	m "anxiel_cube/models"
)

//====( This is a function to form the cross of starting face of the cube )=====
func BottomCross(cube *m.Rubik_cube) (string, *m.Rubik_cube) {
	
	referrenceFace := cube.Front_face
	var state string
	for i := 0; i < 4; i++ {
		if cube.Top_Layer.Mid_front.Front_face == referrenceFace {
			if i > 0 {
				for x := 0; x < i; x++ {
					m.Commands = append(m.Commands, "U")
				}
			}
			state = "[BOTTOM CROSS]: Success!"
			break
		}
		cube = cube.U()

	}

	
	// if m.Commands[len(m.Commands)-1] != "U" {
	// 	state = "[BOTTOM CROSS]: Fail!"
	// }
	return state, cube

}