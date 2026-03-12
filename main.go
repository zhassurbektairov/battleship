package main

import (
	app "battleship/internal/app"
	. "battleship/internal/config"
	utils "battleship/internal/utils"
)

func main() {
	utils.ReadUserMap(&Player)
	State = true
	for State == true {
		app.Play(&Player)
	}
}
