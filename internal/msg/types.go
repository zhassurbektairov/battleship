package msg

const (
	AskInput = "Write input: "

	PlayerMove = "Player's move: "
	PcMove     = "Computer's move: "

	Result = "Results of bombing:"

	TopInfo = " Your map    Computer's map\n 1234567890  1234567890"

	Icons = "[ . ] live cell, [ - ] bombed cell\n[ # ] live ship, [ x ] bombed ship"
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
