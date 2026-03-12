package main

import (
	app "battleship/internal/app"
	. "battleship/internal/config"
	utils "battleship/internal/utils"
	"fmt"
)

func main() {
	fmt.Println("Your map:")
	Player.ID = 1
	Pc.ID = 2
	utils.ReadUserMap(&Player)
	utils.GenerateUserMap(&Pc)
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
