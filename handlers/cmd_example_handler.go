package handlers

import "bitwormhole.com/starter/cli"

// ExampleHandler ...
type ExampleHandler struct {
}

func (inst *ExampleHandler) _Impl() (cli.Handler, cli.HandlerRegistry, cli.Help) {
	return inst, inst, inst
}

func (inst *ExampleHandler) name() string {
	return "example"
}

// GetHandlers ...
func (inst *ExampleHandler) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Handler: inst.Run,
		Help:    inst,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *ExampleHandler) GetHelp() *cli.HelpInfo {
	name := inst.name()
	info := &cli.HelpInfo{
		Name:    name,
		Title:   "Example",
		Usage:   "todo...",
		Content: "todo...",
	}
	return info
}

// Run ...
func (inst *ExampleHandler) Run(task *cli.Task) error {
	return nil
}
