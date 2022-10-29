package filters

import (
	"bitwormhole.com/starter/cli"
)

// ClientServerInjectingFilter ...
type ClientServerInjectingFilter struct {
	Context *cli.Context
}

func (inst *ClientServerInjectingFilter) _Impl() (cli.FilterRegistry, cli.Filter) {
	return inst, inst
}

func (inst *ClientServerInjectingFilter) init(c *cli.Context) error {
	inst.Context = c
	return nil
}

// GetFilters ...
func (inst *ClientServerInjectingFilter) GetFilters() []*cli.FilterRegistration {
	fr := &cli.FilterRegistration{
		Name:   "ClientServerInjectingFilter",
		Order:  OrderClientServer,
		Filter: inst,
		OnInit: inst.init,
	}
	return []*cli.FilterRegistration{fr}
}

// Pass ...
func (inst *ClientServerInjectingFilter) Pass(task *cli.Task, chain cli.FilterChain) error {

	ctx := inst.Context
	task.Client = ctx.Client
	task.Server = ctx.Server

	return chain.Pass(task)
}
