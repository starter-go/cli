package filters

// 定义 Filter 的顺序
const (
	OrderExample = iota
	OrderError
	OrderBinding
	OrderClientServer
	OrderMultiline
	OrderCommandPrepare
	OrderHandling
)
