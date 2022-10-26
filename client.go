package cli

import "context"

// Client 表示 CLI 客户端
type Client interface {
	Handler

	// run a task with Context & Command & Arguments
	RunCCA(ctx context.Context, cmd string, args []string) error
}
