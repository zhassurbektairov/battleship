package utils

import (
	. "battleship/internal/config"
	"fmt"
)

func Initialize(user *Data) {
	user.InitialMap = make([][]byte, 10)
	user.ShowMap = make([][]byte, 10)
	user.Ship.Health = make(map[int]int)

	empty := ".........."

	for i := range 10 {
		if user.ID == 1 {
			var s string
			fmt.Scanf("%s\n", &s)

			user.InitialMap[i] = []byte(s)
		} else {
			user.InitialMap[i] = []byte(empty)
		}
		user.ShowMap[i] = []byte(empty)
	}
}

func ReadUserMap(user *Data) {
	Initialize(user)
	CountShip(user)
	ValidateMap(user)
}
