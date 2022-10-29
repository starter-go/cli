package cli

// Configuration ...
type Configuration struct {
	Handlers       []HandlerRegistry
	Filters        []FilterRegistry
	ContextFactory ContextFactory
	ClientFactory  ClientFactory
	ServerFactory  ServerFactory
}
