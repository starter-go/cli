package core

import (
	"github.com/starter-go/cli"
	"github.com/starter-go/cli/handlers"
)

// Handlers ...
type Handlers struct {

	//starter:component
	_as func(cli.HandlerRegistry) //starter:as(".")

}

func (inst *Handlers) _Impl() cli.HandlerRegistry {
	return inst
}

// GetHandlers ...
func (inst *Handlers) GetHandlers() []*cli.HandlerRegistration {
	all := []*cli.HandlerRegistration{}

	all = inst.fill(all, &handlers.ChdirHandler{})
	all = inst.fill(all, &handlers.HelpHandler{})
	all = inst.fill(all, &handlers.LsHandler{})
	all = inst.fill(all, &handlers.MkdirHandler{})
	all = inst.fill(all, &handlers.NowHandler{})
	all = inst.fill(all, &handlers.PwdHandler{})
	all = inst.fill(all, &handlers.SleepHandler{})

	return all
}

func (inst *Handlers) fill(dst []*cli.HandlerRegistration, src cli.HandlerRegistry) []*cli.HandlerRegistration {
	some := src.GetHandlers()
	return append(dst, some...)
}
