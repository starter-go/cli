package cli

// DefaultContextFactory ...
type DefaultContextFactory struct {
}

func (inst *DefaultContextFactory) _Impl() ContextFactory {
	return inst
}

// NewContext ...
func (inst *DefaultContextFactory) NewContext(cfg *Configuration) (*Context, error) {

	ctx := &Context{}
	ctx.Config = cfg

	client, err := cfg.ClientFactory.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	server, err := cfg.ServerFactory.NewServer(ctx)
	if err != nil {
		return nil, err
	}

	filters, err := inst.prepareFilters(ctx)
	if err != nil {
		return nil, err
	}

	handlers, err := inst.prepareHandlers(ctx)
	if err != nil {
		return nil, err
	}

	chain, err := inst.makeFilterChain(ctx, filters)
	if err != nil {
		return nil, err
	}

	ctx.Client = client
	ctx.Server = server
	ctx.CLI = &implCLI{context: ctx}
	ctx.Chain = chain
	ctx.Filters = filters
	ctx.Handlers = handlers

	return ctx, nil
}

func (inst *DefaultContextFactory) makeFilterChain(ctx *Context, src []*FilterRegistration) (FilterChain, error) {
	builder := FilterChainBuilder{}
	for _, hr := range src {
		builder.AddFilterRegistration(hr)
	}
	return builder.Create(), nil
}

func (inst *DefaultContextFactory) prepareFilters(ctx *Context) ([]*FilterRegistration, error) {
	src := ctx.Config.Filters
	dst := make([]*FilterRegistration, 0)
	table := make(map[string]*FilterRegistration)
	for _, hr1 := range src {
		group := hr1.GetFilters()
		for _, hr2 := range group {
			name := hr2.Name
			old := table[name]
			table[name] = hr2
			if old != nil {
				panic("the command name is duplicated, name=" + name)
			}
			inst.doInit(ctx, hr2.OnInit)
			dst = append(dst, hr2)
		}
	}
	return dst, nil
}

func (inst *DefaultContextFactory) prepareHandlers(ctx *Context) ([]*HandlerRegistration, error) {
	src := ctx.Config.Handlers
	dst := make([]*HandlerRegistration, 0)
	table := make(map[string]*HandlerRegistration)
	for _, hr1 := range src {
		group := hr1.GetHandlers()
		for _, hr2 := range group {
			name := hr2.Name
			old := table[name]
			table[name] = hr2
			if old != nil {
				panic("the command name is duplicated, name=" + name)
			}
			inst.doInit(ctx, hr2.OnInit)
			dst = append(dst, hr2)
		}
	}
	return dst, nil
}

func (inst *DefaultContextFactory) doInit(c *Context, fn OnInitFunc) {
	if c == nil || fn == nil {
		return
	}
	err := fn(c)
	if err != nil {
		panic(err)
	}
}
