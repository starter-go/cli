package cli

import "context"

type clientImpl struct {
	chain FilterChain
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
	return inst.chain.Pass(t)
}
