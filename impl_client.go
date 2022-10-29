package cli

import "context"

////////////////////////////////////////////////////////////////////////////////

type clientImpl struct {
	// chain FilterChain
	context *Context
}

func (inst *clientImpl) _Impl() Client {
	return inst
}

func (inst *clientImpl) RunCCA(ctx context.Context, cmd string, args []string) error {
	t := &Task{}
	t.Context = ctx
	t.Command = cmd
	t.Arguments = args
	return inst.Run(t)
}

func (inst *clientImpl) Run(t *Task) error {
	chain := inst.context.Chain
	return chain.Pass(t)
}

////////////////////////////////////////////////////////////////////////////////

// DefaultClientFactory ...
type DefaultClientFactory struct {
}

func (inst *DefaultClientFactory) _Impl() ClientFactory {
	return inst
}

// NewClient ...
func (inst *DefaultClientFactory) NewClient(c *Context) (Client, error) {
	client := &clientImpl{
		context: c,
	}
	return client, nil
}

////////////////////////////////////////////////////////////////////////////////
