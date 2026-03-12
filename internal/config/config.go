package config

type ShipData struct {
	Count  int
	Health map[int]int
	ID     [10][10]int
}

type Data struct {
	InitialMap [][]byte
	ShowMap    [][]byte
	Ship       ShipData
	ID         int
	Moves      int
}

var (
	Player, Pc Data
	State      bool
)

const Letters = "ABCDEFGHIJ"
