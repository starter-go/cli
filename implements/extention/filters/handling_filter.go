package filters

import (
	"github.com/starter-go/cli"
)

// HandlingFilter ...
type HandlingFilter struct {
}

func (inst *HandlingFilter) _Impl() (cli.FilterRegistry, cli.Filter) {
	return inst, inst
}

// GetFilters ...
func (inst *HandlingFilter) GetFilters() []*cli.FilterRegistration {
	fr := &cli.FilterRegistration{
		Name:   "HandlingFilter",
		Order:  OrderHandling,
		Filter: inst,
	}
	return []*cli.FilterRegistration{fr}
}

// Pass ...
func (inst *HandlingFilter) Pass(task *cli.Task, chain cli.FilterChain) error {

	name := task.Command
	h, err := task.Server.FindHandler(name)
	if err != nil {
		return err
	}

	fn := h.Handler
	err = fn(task)
	if err != nil {
		return err
	}

	return chain.Pass(task)
}
