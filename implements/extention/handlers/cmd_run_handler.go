package handlers

import (
	"fmt"
	"os/exec"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
	"github.com/starter-go/cli"
)

// RunHandler ...
type RunHandler struct {

	//starter:component
	_as func(cli.HandlerRegistry) //starter:as(".")

}

func (inst *RunHandler) _Impl() (cli.Handler, cli.HandlerRegistry, cli.Help) {
	return inst, inst, inst
}

func (inst *RunHandler) name() string {
	return "run"
}

// GetHandlers ...
func (inst *RunHandler) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Handler: inst.Run,
		Help:    inst,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *RunHandler) GetHelp() *cli.HelpInfo {
	name := inst.name()
	info := &cli.HelpInfo{
		Name:    name,
		Title:   "Run",
		Usage:   "run command args ...",
		Content: "Run a executable with filepath",
	}
	return info
}

// Run ...
func (inst *RunHandler) Run(task *cli.Task) error {

	ctx := task.Context
	args := task.Arguments
	name := args[0]
	ar2 := args[1:]

	exefile, err := inst.resolveExePath(name)
	if err == nil && exefile != nil {
		// name = exefile.GetName()
		name = exefile.GetPath()
	} else {
		exefile = nil
	}

	cmd := exec.CommandContext(ctx, name, ar2...)
	cmd.Dir = task.WD
	if exefile != nil {
		cmd.Path = exefile.GetPath()
	}

	cmd.Stdout = task.Console.Output()
	cmd.Stderr = task.Console.Error()
	cmd.Stdin = task.Console.Input()

	return cmd.Run()
}

func (inst *RunHandler) resolveExePath(path string) (afs.Path, error) {
	file := files.FS().NewPath(path)
	if file.IsFile() {
		return file, nil
	}
	return nil, fmt.Errorf("path [%s] is not a file", path)
}
