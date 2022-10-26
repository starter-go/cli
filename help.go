package cli

// HelpInfo ... 提供一条命令的帮助信息
type HelpInfo struct {
	Name    string // 命令名称
	Title   string // 命令标题
	Usage   string // 使用方法
	Content string // 具体内容
}

// Help 接口提供帮助
type Help interface {
	GetHelp() *HelpInfo
}
