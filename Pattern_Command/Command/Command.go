package Command

import "L2/Pattern_Command/Engine"

type StartEngine struct {
	en *Engine.Engine
}

func (cm StartEngine) Execute() {
	cm.en.On()
}

func NewStartEngine(en *Engine.Engine) *StartEngine {
	st := new(StartEngine)
	st.en = en
	return st
}

type StopEngine struct {
	en *Engine.Engine
}

func NewStopEngine(en *Engine.Engine) *StopEngine {
	st := new(StopEngine)
	st.en = en
	return st
}
func (cm StopEngine) Execute() {
	cm.en.Off()
}
