package commands

import (
	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/filters"
	"bitwormhole.com/starter/cli/handlers"
)

// DefaultConfig ....
func DefaultConfig(cfg *cli.Configuration) *cli.Configuration {

	if cfg == nil {
		cfg = &cli.Configuration{}
	}

	if cfg.ClientFactory == nil {
		cfg.ClientFactory = &cli.DefaultClientFactory{}
	}

	if cfg.ServerFactory == nil {
		cfg.ServerFactory = &cli.DefaultServerFactory{}
	}

	if cfg.ContextFactory == nil {
		cfg.ContextFactory = &cli.DefaultContextFactory{}
	}

	builder := configBuilder{config: cfg}

	builder.RegisterFilter(&filters.BindingFilter{})
	builder.RegisterFilter(&filters.ClientServerInjectingFilter{})
	builder.RegisterFilter(&filters.CommandPrepareFilter{})
	builder.RegisterFilter(&filters.ErrorFilter{})
	builder.RegisterFilter(&filters.HandlingFilter{})
	builder.RegisterFilter(&filters.MultilineCommandFilter{})

	builder.RegisterHandler(&handlers.ChdirHandler{})
	builder.RegisterHandler(&handlers.HelpHandler{})
	builder.RegisterHandler(&handlers.LsHandler{})
	builder.RegisterHandler(&handlers.MkdirHandler{})
	builder.RegisterHandler(&handlers.PwdHandler{})
	builder.RegisterHandler(&handlers.SleepHandler{})

	return cfg
}

type configBuilder struct {
	config *cli.Configuration
}

func (inst *configBuilder) RegisterFilter(x cli.FilterRegistry) {
	list := inst.config.Filters
	list = append(list, x)
	inst.config.Filters = list
}

func (inst *configBuilder) RegisterHandler(x cli.HandlerRegistry) {
	list := inst.config.Handlers
	list = append(list, x)
	inst.config.Handlers = list
}
