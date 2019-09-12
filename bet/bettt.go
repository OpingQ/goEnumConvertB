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

func (myEnum BetZone) FullString() string {
	return [...]string{"Default","","Top","Down","","","","","My2","Jiod","","","","","","","","","","","Right","Left"}[myEnum]
}
