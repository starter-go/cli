package handlers

import "bitwormhole.com/starter/cli"

// PwdHandler ...
type PwdHandler struct {
}

func (inst *PwdHandler) _Impl() (cli.Handler, cli.HandlerRegistry, cli.Help) {
	return inst, inst, inst
}

// GetHandlers ...
func (inst *PwdHandler) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "pwd",
		Handler: inst.Run,
		Help:    inst,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *PwdHandler) GetHelp() *cli.HelpInfo {
	info := &cli.HelpInfo{
		Name:    "pwd",
		Title:   "Pwd",
		Usage:   "todo...",
		Content: "todo...",
	}
	return info
}

// Run ...
func (inst *PwdHandler) Run(task *cli.Task) error {
	wd := task.WD
	out := task.Console.Out()
	out.WriteString(wd + "\n")
	return nil
}
