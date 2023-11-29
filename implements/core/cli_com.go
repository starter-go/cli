package core

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/cli"
)

// ComCLI ... 实现一个自动配置的 CLI
type ComCLI struct {
	//// markup.Component `id:"cli" class:"life"`

	//starter:component
	_as func(cli.CLI, application.Lifecycle) //starter:as("#",".")

	Handlers []cli.HandlerRegistry //starter:inject(".")
	Filters  []cli.FilterRegistry  //starter:inject(".")

	impl cli.CLI
}

func (inst *ComCLI) _Impl() (cli.CLI, application.Lifecycle) {
	return inst, inst
}

// Life ...
func (inst *ComCLI) Life() *application.Life {
	return &application.Life{
		OnStart: inst.start,
	}
}

func (inst *ComCLI) getConfiguration() *cli.Configuration {

	cfg := &cli.Configuration{}

	cfg.Filters = inst.Filters
	cfg.Handlers = inst.Handlers
	cfg.ClientFactory = &DefaultClientFactory{}
	cfg.ServerFactory = &DefaultServerFactory{}
	cfg.ContextFactory = &DefaultContextFactory{}

	return cfg
}

func (inst *ComCLI) loadImpl() (cli.CLI, error) {
	cfg := inst.getConfiguration()
	factory := cfg.ContextFactory
	ctx, err := factory.NewContext(cfg)
	if err != nil {
		return nil, err
	}
	return ctx.CLI, nil
}

func (inst *ComCLI) getImpl() (cli.CLI, error) {
	i := inst.impl
	if i != nil {
		return i, nil
	}
	i, err := inst.loadImpl()
	if err != nil {
		return nil, err
	}
	inst.impl = i
	return i, nil
}

func (inst *ComCLI) start() error {
	_, err := inst.getImpl()
	return err
}

func (inst *ComCLI) getRequiredImpl() cli.CLI {
	i, err := inst.getImpl()
	if err != nil {
		panic(err)
	}
	return i
}

// GetClient 取客户端
func (inst *ComCLI) GetClient() cli.Client {
	return inst.getRequiredImpl().GetClient()
}

// GetServer 取服务端
func (inst *ComCLI) GetServer() cli.Server {
	return inst.getRequiredImpl().GetServer()
}

// Bind 把这个 CLI 绑定到指定的 Context
func (inst *ComCLI) Bind(cc context.Context) context.Context {
	return inst.getRequiredImpl().Bind(cc)
}
