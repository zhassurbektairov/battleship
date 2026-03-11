package utils

import (
	. "battleship/internal/config"
	"fmt"
	"os"
)

func ValidateMap(user *Data) {
	ships := map[int]int{}

	for id := range user.ShipHealth {
		size := user.ShipHealth[id]
		if size > 4 {
			fmt.Println("Invalid ship size")
			os.Exit(0)
		}
		ships[size]++
	}

	if ships[4] != 1 || ships[3] != 2 || ships[2] != 3 || ships[1] != 4 {
		fmt.Println("Wrong ships configuration")
		os.Exit(0)
	}

	for i := range 10 {
		for j := range 10 {
			if user.InitialMap[i][j] == 'V' {
				for k := 4; k < 8; k++ {
					nx := i + dir[k][0]
					ny := j + dir[k][1]

					if !inBound(nx, ny) {
						continue
					}

					if user.InitialMap[nx][ny] == 'V' {
						fmt.Println("Ships touching diagonally")
						os.Exit(0)
					}
				}
			}
		}
	}
}
