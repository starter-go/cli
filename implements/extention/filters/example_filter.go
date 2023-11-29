package filters

import (
	"github.com/starter-go/cli"
)

// ExampleFilter ...
type ExampleFilter struct {

	//starter:component
	_as func(cli.FilterRegistry) //starter:as(".")

}

func (inst *ExampleFilter) _Impl() (cli.FilterRegistry, cli.Filter) {
	return inst, inst
}

// GetFilters ...
func (inst *ExampleFilter) GetFilters() []*cli.FilterRegistration {
	fr := &cli.FilterRegistration{
		Name:   "ExampleFilter",
		Order:  OrderExample,
		Filter: inst,
	}
	return []*cli.FilterRegistration{fr}
}

// Pass ...
func (inst *ExampleFilter) Pass(task *cli.Task, chain cli.FilterChain) error {
	return chain.Pass(task)
}
