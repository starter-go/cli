package filters

import (
	"errors"
	"fmt"

	"bitwormhole.com/starter/cli"
)

type errorFilterHolder struct {
	err error
}

////////////////////////////////////////////////////////////////////////////////

// ErrorFilter ...
type ErrorFilter struct {
}

func (inst *ErrorFilter) _Impl() (cli.FilterRegistry, cli.Filter) {
	return inst, inst
}

// GetFilters ...
func (inst *ErrorFilter) GetFilters() []*cli.FilterRegistration {
	fr := &cli.FilterRegistration{
		Name:   "ErrorFilter",
		Order:  OrderError,
		Filter: inst,
	}
	return []*cli.FilterRegistration{fr}
}

// Pass ...
func (inst *ErrorFilter) Pass(task *cli.Task, chain cli.FilterChain) error {
	h := errorFilterHolder{}
	inst.do(task, chain, &h)
	return h.err
}

func (inst *ErrorFilter) do(task *cli.Task, chain cli.FilterChain, h *errorFilterHolder) {
	defer func() {
		x := recover()
		err2 := inst.x2err(x)
		if err2 != nil {
			h.err = err2
		}
	}()
	err1 := chain.Pass(task)
	h.err = err1
}

func (inst *ErrorFilter) x2err(x any) error {
	if x == nil {
		return nil
	}
	err, ok := x.(error)
	if ok {
		return err
	}
	msg := fmt.Sprint(x)
	return errors.New(msg)
}
