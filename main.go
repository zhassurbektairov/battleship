package main

import (
	app "battleship/internal/app"
	. "battleship/internal/config"
	msg "battleship/internal/msg"
	utils "battleship/internal/utils"
	"fmt"
)

func main() {
	fmt.Println(msg.AskInput)
	Player.ID = 1
	Pc.ID = 2
	utils.ReadUserMap(&Player)
	utils.GenerateUserMap(&Pc)
	State = true
	for State == true {
		utils.PrintMap(false, false)
		fmt.Print(msg.PlayerMove)
		app.Destroy(&Pc)
		fmt.Print(msg.PcMove)
		app.Destroy(&Player)
		fmt.Println(msg.Result)
	}
	utils.PrintMap(true, true)
}
