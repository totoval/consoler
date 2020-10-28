package facade

import "github.com/totoval/consoler/pkg/structs"

type Consoler interface {
	Println(color structs.Color, msg string)
	Sprintf(color structs.Color, format string, a ...interface{}) string
	Sprint(color structs.Color, a ...interface{}) string
}
