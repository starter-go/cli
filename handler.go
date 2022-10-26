package cli

// HandlerFunc 定义命令处理函数的签名
type HandlerFunc func(task *Task) error

// Handler 定义命令处理函数，但是接口
type Handler interface {
	Run(task *Task) error
}
