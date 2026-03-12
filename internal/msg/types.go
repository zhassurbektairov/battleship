package msg

const (
	AskInput = "Write input: "

	PlayerMove = "Player's move: "
	PcMove     = "Computer's move: "

	Result = "Results of bombing:"

	Icons = ". -> free cell, x -> bombed ship, - -> bombed cell"
)

const (
	ErrorShipSize  = "Invalid ship size (1 <= size <= 4)."
	ErrorShipNum   = "Wrong ships number."
	ErrorShipTouch = "Ships touching each other."

	ErrorUsedXY   = "This coordinate already opened."
	ErrorInputXY  = "Invalid coordinates (A->J and 1->0)."
	ErrorInputLen = "Invalid length of characters in line"
	ErrorChar     = "Invalid character in "
)
