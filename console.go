package cli

import (
	"bufio"
	"io"
)

// Console ... 表示命令的输入输出对象
type Console interface {
	Input() io.Reader
	Output() io.Writer
	Error() io.Writer

	In() *bufio.Reader
	Out() *bufio.Writer
	Err() *bufio.Writer
}
