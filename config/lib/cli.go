package lib

import (
	"context"

	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/cli/config"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// ComCLI ... 实现一个自动配置的 CLI
type ComCLI struct {
	markup.Component `id:"cli" class:"life"`

	Handlers []cli.HandlerRegistry `inject:".cli-handler-registry"`
	Filters  []cli.FilterRegistry  `inject:".cli-filter-registry"`

	impl cli.CLI
}

func (inst *ComCLI) _Impl() (cli.CLI, application.LifeRegistry) {
	return inst, inst
}

// GetLifeRegistration ...
func (inst *ComCLI) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnStart: inst.start,
	}
}

func (inst *ComCLI) loadImpl() (cli.CLI, error) {
	cfg := config.GetDefaultConfiguration()
	cfg.Filters = inst.Filters
	cfg.Handlers = inst.Handlers
	i := cli.New(cfg)
	return i, nil
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
