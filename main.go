package main

import (
	app "battleship/internal/app"
	. "battleship/internal/config"
)

func main() {
	app.ReadUserMap(&Player)
	State = true
	for State == true {
		app.Play()
	}
}
