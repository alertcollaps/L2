package main

import (
	"L2/Pattern_Command/Command"
	"L2/Pattern_Command/Engine"
	"fmt"
)

func main() {
	en := new(Engine.Engine)
	start := Command.NewStartEngine(en)
	stop := Command.NewStopEngine(en)
	dr := Engine.NewDriver(start)
	fmt.Println("State of engine -", en.Check())
	dr.Execute()
	fmt.Println("State of engine now -", en.Check())

	dr = Engine.NewDriver(stop)
	dr.Execute()

	fmt.Println("State of engine after -", en.Check())
}
