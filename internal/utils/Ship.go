package utils

import (
	. "battleship/internal/config"
)

var dir = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
}

func Run(user *Data, x, y int) {
	user.InitialMap[x][y] = 'V'

	user.ShipID[x][y] = user.ShipCount
	user.ShipHealth[user.ShipCount]++

	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]

		if nx < 0 || ny < 0 || nx >= 10 || ny >= 10 {
			continue
		}

		if user.InitialMap[nx][ny] == '#' {
			Run(user, nx, ny)
		}
	}
}

func RevealNeighbours(user *Data, x, y int) {
	if user.InitialMap[x][y] == 'V' {
		user.ShowMap[x][y] = 'x'
		user.InitialMap[x][y] = 'x'
	}
	for i := 0; i < 8; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]

		if nx < 0 || ny < 0 || nx >= 10 || ny >= 10 {
			continue
		}

		if user.InitialMap[nx][ny] == 'V' {
			RevealNeighbours(user, nx, ny)
		} else if user.InitialMap[nx][ny] == '.' {
			user.ShowMap[nx][ny] = '-'
		}
	}
}
