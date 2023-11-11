package config

import (
	"github.com/starter-go/cli"
	"github.com/starter-go/cli/filters"
	"github.com/starter-go/cli/handlers"
)

// GetDefaultConfiguration 取默认配置
func GetDefaultConfiguration() *cli.Configuration {
	loader := defaultConfigurationLoader{}
	return loader.load()
}

////////////////////////////////////////////////////////////////////////////////

type defaultConfigurationLoader struct {
}

func (inst *defaultConfigurationLoader) load() *cli.Configuration {

	cfg := &cli.Configuration{}

	cfg.ClientFactory = &cli.DefaultClientFactory{}
	cfg.ServerFactory = &cli.DefaultServerFactory{}
	cfg.ContextFactory = &cli.DefaultContextFactory{}

	inst.regFilters(cfg)
	inst.regHandlers(cfg)

	return cfg
}

func (inst *defaultConfigurationLoader) regFilters(cfg *cli.Configuration) {

	inst.registerFilter(cfg, &filters.BindingFilter{})
	inst.registerFilter(cfg, &filters.ClientServerInjectingFilter{})
	inst.registerFilter(cfg, &filters.CommandPrepareFilter{})
	inst.registerFilter(cfg, &filters.ErrorFilter{})
	inst.registerFilter(cfg, &filters.HandlingFilter{})
	inst.registerFilter(cfg, &filters.MultilineCommandFilter{})
}

func (inst *defaultConfigurationLoader) regHandlers(cfg *cli.Configuration) {

	inst.registerHandler(cfg, &handlers.ChdirHandler{})
	inst.registerHandler(cfg, &handlers.HelpHandler{})
	inst.registerHandler(cfg, &handlers.LsHandler{})
	inst.registerHandler(cfg, &handlers.MkdirHandler{})
	inst.registerHandler(cfg, &handlers.NowHandler{})
	inst.registerHandler(cfg, &handlers.PwdHandler{})
	inst.registerHandler(cfg, &handlers.SleepHandler{})
}

func (inst *defaultConfigurationLoader) registerFilter(cfg *cli.Configuration, fr cli.FilterRegistry) {
	list := cfg.Filters
	list = append(list, fr)
	cfg.Filters = list
}

func (inst *defaultConfigurationLoader) registerHandler(cfg *cli.Configuration, hr cli.HandlerRegistry) {
	list := cfg.Handlers
	list = append(list, hr)
	cfg.Handlers = list
}
