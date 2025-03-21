package algorithm

import(
	m "anxiel_cube/models"
	"strings"
)

func SolveRC(puzzle *m.Rubik_cube) []string {

	var bottom_state string
	bottom_state, puzzle = BottomCross(puzzle)
	
	if strings.Contains(bottom_state, "Success") {
		return m.Commands
	}else {
		return []string{"[FAIL]: No Solution could be found at this time!"}
	}
	return nil
}