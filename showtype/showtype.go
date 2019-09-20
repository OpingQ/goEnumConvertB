package showtype

import "strconv"

// ShowType Enum
type ShowType int

// ShowType Enum�s��
const (
	Default ShowType = iota
	Normal
	Hot
	New
)

// string fmt
func (myEnum ShowType) String() string {
	return strconv.Itoa(int(myEnum))
}

// Int fmt
func (myEnum ShowType) Int() int {
	return int(myEnum)
}

func (myEnum ShowType) FullString() string {
	return [...]string{"Default", "Normal", "Hot", "New"}[myEnum]
}
