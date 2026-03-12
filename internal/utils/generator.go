package utils

import (
	. "battleship/internal/config"
	"math/rand"
)

var ships = []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}

func canPlace(user *Data, x, y, size, d int) bool {
	for i := 0; i < size; i++ {
		nx := x
		ny := y

		if d == 0 {
			ny += i
		} else {
			nx += i
		}

		if !inBound(nx, ny) {
			return false
		}

		if user.InitialMap[nx][ny] != '.' {
			return false
		}

		for k := 0; k < 8; k++ {
			tx := nx + dir[k][0]
			ty := ny + dir[k][1]

			if !inBound(tx, ty) {
				continue
			}

			if user.InitialMap[tx][ty] == '#' {
				return false
			}
		}
	}

	return true
}

func placeShip(user *Data, x, y, size, d int) {
	for i := 0; i < size; i++ {
		nx := x
		ny := y

		if d == 0 {
			ny += i
		} else {
			nx += i
		}

		user.InitialMap[nx][ny] = '#'
	}
}

func GenerateMap(user *Data) {
	for {
		for i := range 10 {
			for j := range 10 {
				user.InitialMap[i][j] = '.'
			}
		}
		ok := true
		for _, size := range ships {
			placed := false
			for attempts := 0; attempts < 100; attempts++ {
				x := rand.Intn(10)
				y := rand.Intn(10)

				d := rand.Intn(2)

				if canPlace(user, x, y, size, d) {
					placeShip(user, x, y, size, d)
					placed = true
					break
				}
			}
			if !placed {
				ok = false
				break
			}
		}
		if ok {
			break
		}
	}
}

func GenerateUserMap(user *Data) {
	Initialize(user)
	GenerateMap(user)
	CountShip(user)
}
