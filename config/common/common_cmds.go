package common

import (
	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/handlers"
	"github.com/bitwormhole/starter/markup"
)

// Handlers ...
type Handlers struct {
	markup.Component `class:"cli-handler-registry"`
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
