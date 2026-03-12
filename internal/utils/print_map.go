package utils

import (
	. "battleship/internal/config"
	"fmt"
)

func PrintMap() {
	fmt.Println(" 1234567890  1234567890")
	for i := range 10 {
		fmt.Println(string(Letters[i])+string(Player.ShowMap[i]), string(Letters[i])+string(Pc.ShowMap[i]))
	}
}
