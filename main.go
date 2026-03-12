package main

import (
	app "battleship/internal/app"
	. "battleship/internal/config"
	utils "battleship/internal/utils"
	"fmt"
)

func main() {
	fmt.Println("Your map:")
	utils.ReadUserMap(&Player)
	Player.ID = 1
	utils.GenerateMap(&Pc)
	Pc.ID = 2
	State = true
	for State == true {
		utils.PrintMap()
		fmt.Print("PLAYER GO: ")
		app.Destroy(&Pc)
		fmt.Print("PC GO: ")
		app.Destroy(&Player)
		fmt.Println("RESULT:")
	}
	utils.PrintMap()
}
