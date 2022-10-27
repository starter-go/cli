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

////////////////////////////////////////////////////////////////////////////////

// DefaultConsole ...
type DefaultConsole struct {
	i io.Reader
	o io.Writer
	e io.Writer

	bi *bufio.Reader
	bo *bufio.Writer
	be *bufio.Writer
}

func (inst *DefaultConsole) _Impl() Console {
	return inst
}

// Input ...
func (inst *DefaultConsole) Input() io.Reader {
	return inst.i
}

// Output ...
func (inst *DefaultConsole) Output() io.Writer {
	return inst.o
}

// Error ...
func (inst *DefaultConsole) Error() io.Writer {
	return inst.e
}

// In ...
func (inst *DefaultConsole) In() *bufio.Reader {
	r := inst.bi
	if r == nil {
		r = bufio.NewReader(inst.i)
		inst.bi = r
	}
	return r
}

// Out ...
func (inst *DefaultConsole) Out() *bufio.Writer {
	w := inst.bo
	if w == nil {
		w = bufio.NewWriter(inst.o)
		inst.bo = w
	}
	return w
}

// Err ...
func (inst *DefaultConsole) Err() *bufio.Writer {
	w := inst.be
	if w == nil {
		w = bufio.NewWriter(inst.e)
		inst.be = w
	}
	return w
}

////////////////////////////////////////////////////////////////////////////////

// ConsoleBuilder ...
type ConsoleBuilder struct {
	Input  io.Reader
	Output io.Writer
	Error  io.Writer
}

// Create ...
func (inst *ConsoleBuilder) Create() Console {
	return &DefaultConsole{
		i: inst.Input,
		o: inst.Output,
		e: inst.Error,
	}
}

////////////////////////////////////////////////////////////////////////////////
