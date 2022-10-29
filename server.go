package cli

// Server 表示 CLI 服务器
type Server interface {
	FindHandler(name string) (*HandlerRegistration, error)

	RegisterHandler(hr *HandlerRegistration) error

	RegisterHandlers(hr HandlerRegistry) error

	ListNames() []string
}

// ServerFactory ...
type ServerFactory interface {
	NewServer(c *Context) (Server, error)
}
