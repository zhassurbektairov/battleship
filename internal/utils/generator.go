package utils

import (
	. "battleship/internal/config"
)

var ships = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}

func GenerateMap(user *Data) {
	user.InitialMap = make([][]byte, 10)
	user.ShowMap = make([][]byte, 10)
	user.Ship.Health = make(map[int]int)

	empty := ".........."

	for i := range 10 {
		user.InitialMap[i] = []byte(empty)
		user.ShowMap[i] = []byte(empty)
	}
}
