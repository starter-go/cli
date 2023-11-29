package handlers

import (
	"errors"
	"sort"
	"strings"

	"github.com/starter-go/cli"
)

// HelpHandler ...
type HelpHandler struct {
}

func (inst *HelpHandler) _Impl() (cli.Handler, cli.HandlerRegistry, cli.Help) {
	return inst, inst, inst
}

// GetHandlers ...
func (inst *HelpHandler) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "help",
		Handler: inst.Run,
		Help:    inst,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *HelpHandler) GetHelp() *cli.HelpInfo {
	info := &cli.HelpInfo{
		Name:    "help",
		Title:   "Print Help Info",
		Usage:   "help [command]",
		Content: "show the help info about a command, or list command names.",
	}
	return info
}

// Run ...
func (inst *HelpHandler) Run(task *cli.Task) error {
	args := task.Arguments
	if len(args) > 0 {
		name := args[0]
		return inst.printDetail(task, name)
	}
	return inst.printCommandList(task)
}

func (inst *HelpHandler) findHelpInfo(task *cli.Task, cmd string) (*cli.HelpInfo, error) {

	h, err := task.Server.FindHandler(cmd)
	if err != nil {
		return nil, err
	}

	helper := h.Help
	if helper == nil {
		return nil, errors.New("no help info for command:" + cmd)
	}

	info := helper.GetHelp()
	if info == nil {
		return nil, errors.New("no help info for command:" + cmd)
	}

	return info, nil
}

func (inst *HelpHandler) printCommandList(task *cli.Task) error {

	const nl = "\n"
	console := task.Console
	out := console.Out()
	stderr := console.Err()
	list := task.Server.ListNames()
	sort.Strings(list)

	line := strings.Builder{}
	out.WriteString("Commands:" + nl)

	for _, cmd := range list {
		info, err := inst.findHelpInfo(task, cmd)
		if err != nil {
			stderr.WriteString(err.Error() + nl)
			continue
		}
		line.Reset()
		inst.appendPaddingTo(&line, 4)
		line.WriteString(cmd)
		inst.appendPaddingTo(&line, 16)
		line.WriteString(info.Title)
		line.WriteString(nl)
		out.WriteString(line.String())
	}

	return nil
}

func (inst *HelpHandler) appendPaddingTo(b *strings.Builder, n int) {
	size := b.Len()
	for ; size < n; size++ {
		b.WriteRune(' ')
	}
}

func (inst *HelpHandler) printDetail(task *cli.Task, cmd string) error {

	const nl = "\n"
	console := task.Console
	out := console.Out()
	stderr := console.Err()

	info, err := inst.findHelpInfo(task, cmd)
	if err != nil {
		stderr.WriteString(err.Error() + nl)
		return nil
	}

	out.WriteString("command: " + info.Name + nl)
	out.WriteString(info.Title + nl)
	out.WriteString("usage: " + info.Usage + nl)
	out.WriteString(info.Content + nl)

	return nil
}
