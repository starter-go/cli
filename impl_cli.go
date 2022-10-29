package cli

import "context"

type implCLI struct {
	context *Context
}

func (inst *implCLI) _Impl() CLI {
	return inst
}

func (inst *implCLI) GetClient() Client {
	return inst.context.Client
}

func (inst *implCLI) GetServer() Server {
	return inst.context.Server
}

func (inst *implCLI) Bind(cc context.Context) context.Context {
	cc = Bind(cc)
	b := GetBinding(cc)
	b.SetCLI(inst)
	return cc
}
