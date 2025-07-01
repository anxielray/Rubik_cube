package algorithm

import (
	m "anxiel_cube/models"
)

// solve Rubic cube invokes the algorithm to solve the rubic cube and returns the commands used to solve the problm.
func SolveRC(puzzle *m.Rubik_cube) []string {

	_ = BottomCross(puzzle)

	return m.Commands
}
