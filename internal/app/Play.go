package app

import (
	. "battleship/internal/config"
	utils "battleship/internal/utils"
	"fmt"
)

func Play() {
	var input string
	fmt.Scanln(&input)

	x := int(input[0] - 'A')
	y := int(input[1]-'0') - 1

	cell := Player.InitialMap[x][y]

	if cell == 'V' {
		Player.ShowMap[x][y] = 'x'

		id := Player.ShipCoordinates[[2]int{x, y}]

		sunk := true
		for i, shipID := range Player.ShipCoordinates {
			if shipID == id {
				if Player.ShowMap[i[0]][i[1]] != 'x' {
					sunk = false
					break
				}
			}
		}

		if sunk {
			Player.ShipCount--
		}

		if Player.ShipCount == 0 {
			State = false
		}
	} else if cell == '.' {
		Player.ShowMap[x][y] = '-'
	}

	utils.PrintMap(Player)
}

/*
###.......
..........
..##......
..........
..........
......#...
......#...
......#...
..........
..........

.##.......
..........
..........
..........
..........
..........
..........
..........
..........
..........

*/
