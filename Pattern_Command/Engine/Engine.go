package Engine

type Execute interface {
	Execute()
}

type Driver struct {
	command Execute
}

//Constructor
func NewDriver(command Execute) *Driver {
	drv := new(Driver)
	drv.command = command
	return drv
}

func (drv Driver) Execute() {
	drv.command.Execute()
}

type Engine struct {
	state bool
}

func (en *Engine) On() {
	en.state = true
}
func (en *Engine) Off() {
	en.state = false
}
func (en Engine) Check() bool {
	return en.state
}
