package structs

type Color byte
const (
	ColorSuccess Color = iota
	ColorWarning
	ColorInfo
	ColorError
)
