package utils

import (
	. "battleship/internal/config"
	"fmt"
)

func PrintMap(user *Data) {
	for i := range 10 {
		fmt.Println(string(user.ShowMap[i]))
	}
}
