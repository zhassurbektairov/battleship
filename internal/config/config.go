package config

type Data struct {
	InitialMap      [][]byte
	ShowMap         [][]byte
	ShipCoordinates map[[2]int]int
	ShipCount       int
}

var (
	Player, Pc Data
	State      bool
)
