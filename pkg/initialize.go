package pkg

import (
	"errors"
	"github.com/totoval/consoler/internal"
	"github.com/totoval/consoler/internal/defaults"
	"github.com/totoval/logger/pkg/facade"
)

type Console struct {
	console internal.Consoler
}

func (c *Console) Component() interface{} {
	return c.console
}

func (c *Console) Use(driver string) Componentor {
	switch driver {
	case DriverDefault:
		c.console = &defaults.Console{}
	default:
		c.console = &defaults.Console{}
	}
	return c
}

func (c *Console) Config(configuration map[string]interface{}) error {
	logger, ok := configuration["logger"].(facade.Logger)
	if !ok {
		return errors.New("unknown configuration logger")
	}

	return c.console.New().SetLogger(logger)
}