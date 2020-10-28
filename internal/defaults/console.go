package defaults

import (
	"errors"
	color2 "github.com/fatih/color"
	"github.com/totoval/consoler/internal"
	"github.com/totoval/consoler/pkg/structs"
	"github.com/totoval/logger/pkg/facade"
)

type Console struct {
	logger facade.Logger
}

func (c *Console) New() internal.Consoler {
	return c
}
func (c *Console) SetLogger(logger facade.Logger) error {
	c.logger = logger
	return nil
}

func (c *Console) Println(color structs.Color, msg string) {
	attribute, err := c.ConvertColor(color)
	if err != nil {
		_ = c.logger.Error(err)
	}
	if _, err := color2.New(attribute.(color2.Attribute)).Println(msg); err != nil {
		_ = c.logger.Error(err)
	}
}

func (c *Console) Sprintf(color structs.Color, format string, a ...interface{}) string {
	attribute, err := c.ConvertColor(color)
	if err != nil {
		_ = c.logger.Error(err)
	}
	return color2.New(attribute.(color2.Attribute)).Sprintf(format, a...)
}

func (c *Console) Sprint(color structs.Color, a ...interface{}) string {
	attribute, err := c.ConvertColor(color)
	if err != nil {
		_ = c.logger.Error(err)
	}
	return color2.New(attribute.(color2.Attribute)).Sprint(a...)
}

func (c *Console)ConvertColor(color structs.Color) (interface{}, error){
	switch color {
	case structs.ColorSuccess:
		return color2.FgGreen, nil
	case structs.ColorWarning:
		return color2.FgYellow, nil
	case structs.ColorInfo:
		return color2.FgBlue, nil
	case structs.ColorError:
		return color2.FgRed, nil
	default:
		return nil, errors.New("unknown console color")
	}
}

