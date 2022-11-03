package lib

import (
	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/filters"
	"github.com/bitwormhole/starter/markup"
)

// Filters ...
type Filters struct {
	markup.Component `class:"cli-filter-registry"`
}

func (inst *Filters) _Impl() cli.FilterRegistry {
	return inst
}

// GetFilters ...
func (inst *Filters) GetFilters() []*cli.FilterRegistration {
	all := []*cli.FilterRegistration{}

	all = inst.fill(all, &filters.BindingFilter{})
	all = inst.fill(all, &filters.ClientServerInjectingFilter{})
	all = inst.fill(all, &filters.CommandPrepareFilter{})
	all = inst.fill(all, &filters.ErrorFilter{})
	all = inst.fill(all, &filters.HandlingFilter{})
	all = inst.fill(all, &filters.MultilineCommandFilter{})

	return all
}

func (inst *Filters) fill(dst []*cli.FilterRegistration, src cli.FilterRegistry) []*cli.FilterRegistration {
	some := src.GetFilters()
	return append(dst, some...)
}
