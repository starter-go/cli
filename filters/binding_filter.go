package filters

import (
	"context"
	"os"

	"bitwormhole.com/starter/cli"
	"bitwormhole.com/starter/vlog"
)

// BindingFilter ...
type BindingFilter struct {
}

func (inst *BindingFilter) _Impl() (cli.FilterRegistry, cli.Filter) {
	return inst, inst
}

// GetFilters ...
func (inst *BindingFilter) GetFilters() []*cli.FilterRegistration {
	fr := &cli.FilterRegistration{
		Name:   "BindingFilter",
		Order:  OrderBinding,
		Filter: inst,
	}
	return []*cli.FilterRegistration{fr}
}

// Pass ...
func (inst *BindingFilter) Pass(task *cli.Task, chain cli.FilterChain) error {

	// get or create binding
	ctx := task.Context
	if ctx == nil {
		ctx = context.Background()
	}
	b := cli.GetBinding(ctx)
	ctx = b.Context
	task.Context = ctx

	// pre-process
	err := inst.preprocess(task, b)
	if err != nil {
		return err
	}
	// post-process
	defer func() {
		err = inst.postProcess(task, b)
		if err != nil {
			vlog.Warn(err)
		}
	}()

	return chain.Pass(task)
}

func (inst *BindingFilter) prepareConsole(task *cli.Task, b *cli.Binding) (cli.Console, error) {

	con := task.Console
	if con == nil {
		con = b.Console
	}
	if con != nil {
		return con, nil
	}

	cb := &cli.ConsoleBuilder{}
	cb.Error = os.Stderr
	cb.Output = os.Stdout
	cb.Input = os.Stdin

	con = cb.Create()
	return con, nil
}

func (inst *BindingFilter) prepareEnv(task *cli.Task, b *cli.Binding) (map[string]string, error) {
	src := task.Env
	dst := b.Env
	if dst == nil {
		dst = make(map[string]string)
	}
	for k, v := range src {
		dst[k] = v
	}
	return dst, nil
}

func (inst *BindingFilter) prepareWD(task *cli.Task, b *cli.Binding) (string, error) {
	wd := task.WD
	if wd == "" {
		wd = b.WD
	}
	if wd != "" {
		return wd, nil
	}
	return os.Getwd()
}

func (inst *BindingFilter) preprocess(task *cli.Task, b *cli.Binding) error {

	// wd
	wd, err := inst.prepareWD(task, b)
	if err != nil {
		return err
	}
	task.WD = wd
	b.WD = wd

	// console
	console, err := inst.prepareConsole(task, b)
	if err != nil {
		return err
	}
	task.Console = console
	b.Console = console

	// env
	env, err := inst.prepareEnv(task, b)
	if err != nil {
		return err
	}
	task.Env = env
	b.Env = env

	return nil
}

func (inst *BindingFilter) postProcess(task *cli.Task, b *cli.Binding) error {

	console := task.Console
	if console != nil {
		console.Err().Flush()
		console.Out().Flush()
	}

	b.Env = task.Env
	b.WD = task.WD
	b.Console = console
	return nil
}
