package utils

import (
	. "battleship/internal/config"
	msg "battleship/internal/msg"
	"fmt"
)

func PrintMap(a, b bool) {
	fmt.Println(msg.TopInfo)
	for i := range 10 {
		map1 := []byte(Player.ShowMap[i])
		map2 := []byte(Pc.ShowMap[i])

		if a == true {
			for j, c := range Player.InitialMap[i] {
				if map1[j] == '.' && c == 'V' {
					map1[j] = '#'
				}
			}
		}

		if b == true {
			for j, c := range Pc.InitialMap[i] {
				if map2[j] == '.' && c == 'V' {
					map2[j] = '#'
				}
			}
		}

		fmt.Println(string(Letters[i])+string(map1), string(Letters[i])+string(map2))
	}
	fmt.Println(msg.Icons)
}
