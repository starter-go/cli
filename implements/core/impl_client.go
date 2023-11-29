package core

import (
	"context"

	"github.com/starter-go/cli"
)

////////////////////////////////////////////////////////////////////////////////

type clientImpl struct {
	// chain FilterChain
	context *cli.Context
}

func (inst *clientImpl) _Impl() cli.Client {
	return inst
}

func (inst *clientImpl) RunCCA(ctx context.Context, cmd string, args []string) error {
	t := &cli.Task{}
	t.Context = ctx
	t.Command = cmd
	t.Arguments = args
	return inst.Run(t)
}

func (inst *clientImpl) Run(t *cli.Task) error {
	chain := inst.context.Chain

	t.Client = inst.context.Client
	t.Server = inst.context.Server

	return chain.Pass(t)
}

////////////////////////////////////////////////////////////////////////////////

// DefaultClientFactory ...
type DefaultClientFactory struct {
}

func (inst *DefaultClientFactory) _Impl() cli.ClientFactory {
	return inst
}

// NewClient ...
func (inst *DefaultClientFactory) NewClient(c *cli.Context) (cli.Client, error) {
	client := &clientImpl{
		context: c,
	}
	return client, nil
}

////////////////////////////////////////////////////////////////////////////////
