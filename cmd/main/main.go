package main

import (
	"github.com/totoval/consoler/pkg"
	"github.com/totoval/consoler/pkg/facade"
	"github.com/totoval/consoler/pkg/structs"
	"github.com/totoval/framework/helpers/toto"
	logger_pkg "github.com/totoval/logger/pkg"
	logger_facade "github.com/totoval/logger/pkg/facade"
	logger_structs "github.com/totoval/logger/pkg/structs"
)

//@todo register into global vars
var log logger_facade.Logger
var console facade.Consoler

func main()  {
	// init logger
	// use driver then config
	l := &logger_pkg.Log{}
	if err := l.Use(logger_pkg.DriverLogrus).Config(map[string]interface{}{
		"level": logger_structs.LevelTrace,
	}); err != nil {
		panic(err)
	}
	// get the facade
	log = l.Component().(logger_facade.Logger)

	log.Info("test", toto.V{"haha": 123}, toto.V{"1": 2})


	// use driver then config
	c := &pkg.Console{}
	if err := c.Use(pkg.DriverDefault).Config(map[string]interface{}{
		"logger": log,
	}); err != nil {
		panic(err)
	}
	// get the facade
	console = c.Component().(facade.Consoler)

	console.Println(structs.ColorSuccess, "hello world")
}
