package app

import (
	. "battleship/internal/config"
	utils "battleship/internal/utils"
	"fmt"
	"math/rand"
)

func Destroy(user *Data) {
	x := rand.Intn(10)
	y := rand.Intn(10)

	if user.ID == 2 {
		var input string
		fmt.Scanln(&input)

		x = int(input[0] - 'A')
		y = int(input[1]-'0') - 1

		if len(input) == 3 {
			y = 9
		}
	} else {

		fmt.Println(string(Letters[x]), y+1)
	}

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
			if user.ID == 1 {
				fmt.Println("user lost")
			} else {
				fmt.Println("pc lost")
			}
			State = false
		}
	case '.':
		user.ShowMap[x][y] = '-'
	}
}

/*
###...####
..........
...##....#
.#.......#
.#...#....
.#........
.....#....
..........
.....#..#.
#.......#.
*/
