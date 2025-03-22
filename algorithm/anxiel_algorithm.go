package algorithm

import (
	m "anxiel_cube/models"
)

func SolveRC(puzzle *m.Rubik_cube) []string {

	_ = BottomCross(puzzle)

	return m.Commands
}
