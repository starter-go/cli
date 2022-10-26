package cli

import "errors"

// Context ...
type Context struct {
	Client Client
	Server Server

	// Handlers []HandlerRegistry
	// Filters  []FilterRegistry
}

////////////////////////////////////////////////////////////////////////////////

// ContextBuilder ...
type ContextBuilder struct {
	handlers []HandlerRegistry
	filters  []FilterRegistry
	context  *Context
}

// Init ...
func (inst *ContextBuilder) Init(ctx *Context) {
	inst.context = ctx
}

// Reset ...
func (inst *ContextBuilder) Reset() {
	inst.context = nil
	inst.handlers = nil
	inst.filters = nil
}

// Create ...
func (inst *ContextBuilder) Create() *Context {

	chain, err := inst.makeFilterChain()
	if err != nil {
		panic(err)
	}

	hTable, err := inst.makeHandlerTable()
	if err != nil {
		panic(err)
	}

	client := &clientImpl{}
	server := &serverImpl{}

	client.chain = chain
	server.handlers = hTable

	ctx := inst.context
	if ctx == nil {
		ctx = &Context{}
	}
	ctx.Client = client
	ctx.Server = server
	return ctx
}

func (inst *ContextBuilder) makeFilterChain() (FilterChain, error) {
	fcb := FilterChainBuilder{}
	src := inst.filters
	for _, fr := range src {
		fcb.AddFilterRegistry(fr)
	}
	chain := fcb.Create()
	return chain, nil
}

func (inst *ContextBuilder) makeHandlerTable() (map[string]*HandlerRegistration, error) {
	table := make(map[string]*HandlerRegistration)
	src := inst.handlers
	for _, hr1 := range src {
		group := hr1.GetHandlers()
		for _, hr2 := range group {
			name := hr2.Name
			old := table[name]
			err := inst.checkHandler(name, old, hr2)
			if err != nil {
				return nil, err
			}
			table[name] = hr2
		}
	}
	return table, nil
}

func (inst *ContextBuilder) checkHandler(name string, old, new *HandlerRegistration) error {
	if old == nil {
		return nil
	}
	return errors.New("the command name is duplicate, name=" + name)
}

// RegisterHandler ...
func (inst *ContextBuilder) RegisterHandler(hr HandlerRegistry) {
	if hr != nil {
		inst.handlers = append(inst.handlers, hr)
	}
}

// RegisterFilter ...
func (inst *ContextBuilder) RegisterFilter(fr FilterRegistry) {
	if fr != nil {
		inst.filters = append(inst.filters, fr)
	}
}

////////////////////////////////////////////////////////////////////////////////
