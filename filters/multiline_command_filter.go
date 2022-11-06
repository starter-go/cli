package filters

import (
	"strings"

	"bitwormhole.com/starter/cli"
)

// MultilineCommandFilter ...
type MultilineCommandFilter struct {
}

func (inst *MultilineCommandFilter) _Impl() (cli.FilterRegistry, cli.Filter) {
	return inst, inst
}

// GetFilters ...
func (inst *MultilineCommandFilter) GetFilters() []*cli.FilterRegistration {
	fr := &cli.FilterRegistration{
		Name:   "MultilineCommandFilter",
		Order:  OrderMultiline,
		Filter: inst,
	}
	return []*cli.FilterRegistration{fr}
}

// Pass ...
func (inst *MultilineCommandFilter) Pass(task1 *cli.Task, chain cli.FilterChain) error {
	task2 := task1
	script := inst.getScript(task2)
	rows := inst.scriptToRows(script)
	defer func() {
		inst.postProcess(task1, task2)
	}()
	for _, row := range rows {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}
		task2 = inst.prepareTask(task2, row)
		err := chain.Pass(task2)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *MultilineCommandFilter) getScript(task *cli.Task) string {
	builder := strings.Builder{}
	builder.WriteString(task.Command)
	for _, a := range task.Arguments {
		if builder.Len() > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(a)
	}
	return builder.String()
}

func (inst *MultilineCommandFilter) scriptToRows(script string) []string {
	const (
		nl  = "\n"
		nlx = "\r"
	)
	script = strings.ReplaceAll(script, nlx, nl)
	return strings.Split(script, nl)
}

func (inst *MultilineCommandFilter) prepareTask(src *cli.Task, cmd string) *cli.Task {
	dst := &cli.Task{
		Client:    src.Client,
		Console:   src.Console,
		Context:   src.Context,
		Env:       src.Env,
		Server:    src.Server,
		WD:        src.WD,
		Command:   cmd,
		Arguments: nil,
	}
	return dst
}

func (inst *MultilineCommandFilter) postProcess(task1, task2 *cli.Task) {
	task1.Context = task2.Context
	task1.Console = task2.Console
	task1.Env = task2.Env
	task1.WD = task2.WD
}
