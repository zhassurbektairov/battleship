package utils

import (
	. "battleship/internal/config"
	msg "battleship/internal/msg"
	"fmt"
)

func PrintMap(a, b bool) {
	fmt.Println(" Your map    Computer's map")
	fmt.Println(" 1234567890  1234567890")
	for i := range 10 {
		map1, map2 := "", ""

		if a == true {
			map1 = string(Player.InitialMap[i])
		} else {
			map1 = string(Player.ShowMap[i])
		}

		if b == true {
			map2 = string(Pc.InitialMap[i])
		} else {
			map2 = string(Pc.ShowMap[i])
		}

		fmt.Println(string(Letters[i])+map1, string(Letters[i])+map2)
	}
	fmt.Println(msg.Icons)
}
