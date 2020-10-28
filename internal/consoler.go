package internal

import (
	"github.com/totoval/consoler/pkg/structs"
	"github.com/totoval/logger/pkg/facade"
)

type Consoler interface {
	New() Consoler
	SetLogger(logger facade.Logger) error
	ConvertColor(color structs.Color) (interface{}, error)
	Println(color structs.Color, msg string)
	Sprintf(color structs.Color, format string, a ...interface{}) string
	Sprint(color structs.Color, a ...interface{}) string
}