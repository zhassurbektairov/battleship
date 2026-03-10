package utils

import (
	. "battleship/internal/config"
	"fmt"
)

func PrintMap(user Data) {
	for i := 0; i < 10; i++ {
		fmt.Println(string(user.ShowMap[i]))
	}
}
