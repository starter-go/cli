package filters

import (
	"bitwormhole.com/starter/cli"
)

// MultilineCommandFilter ...
type MultilineCommandFilter struct {
}

func (inst *MultilineCommandFilter) _Impl() (cli.FilterRegistry, cli.Filter) {
	return inst, inst
}

// GetFilters ...
func (inst *MultilineCommandFilter) GetFilters() []*cli.FilterRegistration {
	fr := &cli.FilterRegistration{
		Name:   "MultilineCommandFilter",
		Order:  OrderMultiline,
		Filter: inst,
	}
	return []*cli.FilterRegistration{fr}
}

// Pass ...
func (inst *MultilineCommandFilter) Pass(task *cli.Task, chain cli.FilterChain) error {
	return chain.Pass(task)
}
