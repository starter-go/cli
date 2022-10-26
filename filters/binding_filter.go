package filters

import (
	"context"

	"bitwormhole.com/starter/cli"
)

// BindingFilter ...
type BindingFilter struct {
}

func (inst *BindingFilter) _Impl() (cli.FilterRegistry, cli.Filter) {
	return inst, inst
}

// GetFilters ...
func (inst *BindingFilter) GetFilters() []*cli.FilterRegistration {
	fr := &cli.FilterRegistration{
		Name:   "BindingFilter",
		Order:  OrderBinding,
		Filter: inst,
	}
	return []*cli.FilterRegistration{fr}
}

// Pass ...
func (inst *BindingFilter) Pass(task *cli.Task, chain cli.FilterChain) error {

	ctx := task.Context
	if ctx == nil {
		ctx = context.Background()
	}
	b := cli.GetBinding(ctx)
	ctx = b.Context
	task.Context = ctx

	return chain.Pass(task)
}
