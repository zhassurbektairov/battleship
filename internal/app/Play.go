package app

import (
	. "battleship/internal/config"
	utils "battleship/internal/utils"
	"fmt"
)

func Play(user *Data) {
	var input string
	fmt.Scanln(&input)

	x := int(input[0] - 'A')
	y := int(input[1]-'0') - 1

	switch user.InitialMap[x][y] {
	case 'V':
		user.ShowMap[x][y] = 'x'

		id := user.Ship.ID[x][y]
		user.Ship.Health[id]--

		if user.Ship.Health[id] == 0 {
			utils.RevealNeighbours(user, x, y)
			user.Ship.Count--
		}

		if user.Ship.Count == 0 {
			State = false
		}
	case '.':
		user.ShowMap[x][y] = '-'
	}

	utils.PrintMap(user)
}

/*
####..####
...#......
...#.....#
.#.......#
.#........
.#........
..,..#....
..........
.....#..#.
#.......#.

.##.......
....##....
..........
##........
..........
..........
..........
..........
..........
..........
*/
