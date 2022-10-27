package handlers

import "bitwormhole.com/starter/cli"

// ExampleHandler ...
type ExampleHandler struct {
}

func (inst *ExampleHandler) _Impl() (cli.Handler, cli.HandlerRegistry, cli.Help) {
	return inst, inst, inst
}

// GetHandlers ...
func (inst *ExampleHandler) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "example",
		Handler: inst.Run,
		Help:    inst,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *ExampleHandler) GetHelp() *cli.HelpInfo {
	info := &cli.HelpInfo{
		Name:    "example",
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
