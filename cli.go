package cli

import (
	"context"

	"bitwormhole.com/starter/vlog"
)

// CLI ...
type CLI interface {
	GetClient() Client
	GetServer() Server
}

// Init 初始化 cli 模块
func Init(c context.Context, config *Configuration) context.Context {

	if c == nil {
		c = context.Background()
	}

	if config == nil {
		config = makeDefaultConfig()
	}

	ba := bindingAccess{}
	c = ba.setup(c)
	b, err := ba.getBinding(c)
	if err != nil {
		panic(err)
	}

	ctx := b.context
	if ctx == nil {
		b.config = config
		ctx, err = config.ContextFactory.NewContext(config)
		if err != nil {
			panic(err)
		}
		b.context = ctx
	}

	return c
}

// GetClient ...
func GetClient(c context.Context) Client {
	ba := bindingAccess{}
	b, err := ba.getBinding(c)
	if err != nil {
		vlog.Error(err)
		panic("need for cli.Init()")
	}
	ctx := b.context
	return ctx.Client
}

func makeDefaultConfig() *Configuration {
	c := &Configuration{}
	c.ContextFactory = &DefaultContextFactory{}
	c.ClientFactory = &DefaultClientFactory{}
	c.ServerFactory = &DefaultServerFactory{}
	c.Handlers = nil
	c.Filters = nil
	return c
}
