package core

import (
	"context"

	"github.com/starter-go/cli"
)

type implCLI struct {
	context *cli.Context
}

func (inst *implCLI) _Impl() cli.CLI {
	return inst
}

func (inst *implCLI) GetClient() cli.Client {
	return inst.context.Client
}

func (inst *implCLI) GetServer() cli.Server {
	return inst.context.Server
}

func (inst *implCLI) Bind(cc context.Context) context.Context {
	cc = cli.Bind(cc)
	b := cli.GetBinding(cc)
	b.SetCLI(inst)
	return cc
}
