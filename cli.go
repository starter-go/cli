package cli

import "context"

// CLI ... [inject:"#cli"]
type CLI interface {

	// 取客户端
	GetClient() Client

	// 取服务端
	GetServer() Server

	// 把这个 CLI 绑定到指定的 Context
	Bind(cc context.Context) context.Context
}

// // New 初始化 cli 模块
// func New(cfg *Configuration) CLI {
// 	cfg = prepareConfig(cfg)
// 	ctx, err := cfg.ContextFactory.NewContext(cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return ctx.CLI
// }

////////////////////////////////////////////////////////////////////////////////

// func prepareConfig(c *Configuration) *Configuration {

// 	if c == nil {
// 		c = &Configuration{}
// 	}

// 	if c.ContextFactory == nil {
// 		c.ContextFactory = &DefaultContextFactory{}
// 	}

// 	if c.ClientFactory == nil {
// 		c.ClientFactory = &DefaultClientFactory{}
// 	}

// 	if c.ServerFactory == nil {
// 		c.ServerFactory = &DefaultServerFactory{}
// 	}

// 	// c.Handlers = nil
// 	// c.Filters = nil

// 	return c
// }
