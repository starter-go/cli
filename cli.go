package cli

import "context"

// CLI ...
type CLI interface {
	GetClient() Client
	GetServer() Server
}

// New 初始化 cli 模块
func New(cfg *Configuration) CLI {
	cfg = prepareConfig(cfg)
	ctx, err := cfg.ContextFactory.NewContext(cfg)
	if err != nil {
		panic(err)
	}
	return ctx.CLI
}

// Bind ...
func Bind(cc context.Context) context.Context {
	ba := bindingAccess{}
	b, err := ba.getBinding(cc)
	if err == nil && b != nil {
		return cc
	}
	cc = ba.setup(cc)
	b, err = ba.getBinding(cc)
	if err != nil {
		panic(err)
	}
	return cc
}

////////////////////////////////////////////////////////////////////////////////

func prepareConfig(c *Configuration) *Configuration {

	if c == nil {
		c = &Configuration{}
	}

	if c.ContextFactory == nil {
		c.ContextFactory = &DefaultContextFactory{}
	}

	if c.ClientFactory == nil {
		c.ClientFactory = &DefaultClientFactory{}
	}

	if c.ServerFactory == nil {
		c.ServerFactory = &DefaultServerFactory{}
	}

	// c.Handlers = nil
	// c.Filters = nil

	return c
}
