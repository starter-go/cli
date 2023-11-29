package handlers

import (
	"strconv"
	"time"

	"github.com/starter-go/cli"
)

// SleepHandler ...
type SleepHandler struct {

	//starter:component
	_as func(cli.HandlerRegistry) //starter:as(".")

}

func (inst *SleepHandler) _Impl() (cli.Handler, cli.HandlerRegistry, cli.Help) {
	return inst, inst, inst
}

// GetHandlers ...
func (inst *SleepHandler) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "sleep",
		Handler: inst.Run,
		Help:    inst,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *SleepHandler) GetHelp() *cli.HelpInfo {
	info := &cli.HelpInfo{
		Name:    "sleep",
		Title:   "Sleep",
		Usage:   "todo...",
		Content: "todo...",
	}
	return info
}

// Run ...
func (inst *SleepHandler) Run(task *cli.Task) error {
	a1 := task.Arguments[0]
	ms, err := strconv.ParseInt(a1, 10, 0)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond * time.Duration(ms))
	return nil
}
