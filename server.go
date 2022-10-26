package cli

// Server 表示 CLI 服务器
type Server interface {
	FindHandler(name string) (*HandlerRegistration, error)
}
