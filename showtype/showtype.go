package showtype

// ShowType Enum
type ShowType int

// ShowType Enum½s¸¹
const (
	Default ShowType = iota
	Normal
	Hot
	New
	
)

// string fmt
func (myEnum ShowType) String() string {
	return string(myEnum)
}

// Int fmt
func (myEnum ShowType) Int() int {
	return int(myEnum)
}


func (myEnum ShowType) FullString() string {
	return [...]string{"Default","Normal","Hot","New"}[myEnum]
}
