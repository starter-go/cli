package cli

import "context"

// Task 表示要执行的任务
type Task struct {
	Context   context.Context
	Server    Server
	Client    Client
	Command   string
	Arguments []string
	Console   Console
	WD        string
	Env       map[string]string
}
