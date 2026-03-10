package app

import (
	. "battleship/internal/config"
	utils "battleship/internal/utils"
	"fmt"
)

func ReadUserMap(user *Data) {
	user.InitialMap = make([][]byte, 10)
	user.ShowMap = make([][]byte, 10)
	empty := ".........."

	for i := 0; i < 10; i++ {
		var s string
		fmt.Scanf("%s\n", &s)

		user.InitialMap[i] = []byte(s)
		user.ShowMap[i] = []byte(empty)
	}

	user.ShipCoordinates = make(map[[2]int]int)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if user.InitialMap[i][j] == '#' {
				user.ShipCount++
				utils.Run(user, i, j)
			}
		}
	}
}
