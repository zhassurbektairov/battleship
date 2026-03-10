package config

type Data struct {
	InitialMap [][]byte
	ShowMap    [][]byte
	ShipID     [10][10]int
	ShipHealth map[int]int
	ShipCount  int
}

var (
	Player, Pc Data
	State      bool
)
