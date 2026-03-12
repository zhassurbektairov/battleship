package app

import (
	. "battleship/internal/config"
	msg "battleship/internal/msg"
	utils "battleship/internal/utils"
	"fmt"
	"math/rand"
)

func getCoord() (bool, int, int) {
	var input string
	fmt.Scanln(&input)
	x, y := -1, -1

	if len(input) == 2 {
		x = int(input[0] - 'A')
		y = int(input[1]-'0') - 1
		if x < 0 || x > 9 {
			fmt.Println(msg.ErrorInputXY)
			return false, x, y
		}
		if y < -1 || y > 8 {
			fmt.Println(msg.ErrorInputXY)
			return false, x, y
		}
		if y == -1 {
			y = 9
		}
	} else {
		fmt.Println(msg.ErrorInputXY)
		return false, x, y
	}

	return true, x, y
}

func getXY(user *Data) (int, int) {
	x := rand.Intn(10)
	y := rand.Intn(10)

	if user.ID == 2 {
		flag := false
		for flag == false {
			flag, x, y = getCoord()
		}
		for user.ShowMap[x][y] == '-' || user.ShowMap[x][y] == 'x' {
			fmt.Println(msg.ErrorUsedXY)

			flag = false
			for flag == false {
				flag, x, y = getCoord()
			}
		}
	} else {
		for user.ShowMap[x][y] == '-' || user.ShowMap[x][y] == 'x' {
			x = rand.Intn(10)
			y = rand.Intn(10)
		}
		fmt.Println(string(Letters[x]), y+1)
	}

	return x, y
}

func Destroy(user *Data) bool {
	x, y := getXY(user)

	switch user.InitialMap[x][y] {
	case 'V':
		user.ShowMap[x][y] = 'x'

		id := user.Ship.ID[x][y]
		user.Ship.Health[id]--

		if user.Ship.Health[id] == 0 {
			utils.RevealNeighbours(user, x, y)
			user.Ship.Count--
		}

		if user.Ship.Count == 0 {
			if user.ID == 1 {
				fmt.Println("user lost")
			} else {
				fmt.Println("pc lost")
			}
			State = false
		}
		return true
	case '.':
		user.ShowMap[x][y] = '-'
	}
	return false
}
