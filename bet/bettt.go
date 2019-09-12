package bet

type BetZone int

const (
	Default BetZone = iota
	Top BetZone = 1 + iota
	Down
	My2 BetZone = 5 + iota
	Jiod
	Right BetZone = 15 + iota
	Left
	
)

func (myEnum BetZone) String() string {
	return string(myEnum)
}

func (myEnum BetZone) Int() int {
	return int(myEnum)
}
