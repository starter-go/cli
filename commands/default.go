package commands

import (
	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/filters"
	"bitwormhole.com/starter/cli/handlers"
)

// DefaultContextBuilder ....
func DefaultContextBuilder() *cli.ContextBuilder {

	ctx := &cli.Context{}
	builder := &cli.ContextBuilder{}
	builder.Init(ctx)

	builder.RegisterFilter(&filters.BindingFilter{})
	builder.RegisterFilter(&filters.ClientServerInjectingFilter{Context: ctx})
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

	return builder
}
