package utils

import (
	. "battleship/internal/config"
	msg "battleship/internal/msg"
	"fmt"
	"os"
)

func ValidateMap(user *Data) {
	ships := map[int]int{}

	for row, val := range user.InitialMap {
		if len(val) != 10 {
			fmt.Println(msg.ErrorInputLen, row+1)
			os.Exit(0)
		}
		for col, char := range val {
			if char != '.' && char != 'V' {
				fmt.Printf("%s row %d, column %d", msg.ErrorChar, row+1, col+1)
				os.Exit(0)
			}
		}
	}

	for id := range user.Ship.Health {
		size := user.Ship.Health[id]
		if size > 4 {
			fmt.Println(msg.ErrorShipSize)
			os.Exit(0)
		}
		ships[size]++
	}

	if ships[4] != 1 || ships[3] != 2 || ships[2] != 3 || ships[1] != 4 {
		fmt.Println(msg.ErrorShipNum)
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
						fmt.Println(msg.ErrorShipTouch)
						os.Exit(0)
					}
				}
			}
		}
	}
}
