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
		utils.PrintMap(true, false)
		fmt.Print(msg.PlayerMove)
		for app.Destroy(&Pc) && State {
			Player.Moves++
			utils.PrintMap(true, false)
			fmt.Print(msg.PlayerMove)
		}
		if !State {
			break
		}
		fmt.Print(msg.PcMove)
		for app.Destroy(&Player) && State {
			Pc.Moves++
			fmt.Print(msg.PcMove)
		}
		fmt.Println(msg.Result)
	}
	utils.PrintMap(true, true)
}
