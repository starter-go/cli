package handlers

import (
	"errors"
	"strings"

	"bitwormhole.com/starter/afs/files"
	"bitwormhole.com/starter/cli"
)

// ChdirHandler ...
type ChdirHandler struct {
}

func (inst *ChdirHandler) _Impl() (cli.Handler, cli.HandlerRegistry, cli.Help) {
	return inst, inst, inst
}

// GetHandlers ...
func (inst *ChdirHandler) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "cd",
		Handler: inst.Run,
		Help:    inst,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *ChdirHandler) GetHelp() *cli.HelpInfo {
	info := &cli.HelpInfo{
		Name:    "cd",
		Title:   "change directory",
		Usage:   "todo...",
		Content: "todo...",
	}
	return info
}

// Run ...
func (inst *ChdirHandler) Run(task *cli.Task) error {

	wd := task.WD
	to := task.Arguments[0]
	current := files.FS().NewPath(wd)
	next := current
	dst := strings.TrimSpace(to)
	if strings.HasPrefix(dst, "/") {
		next = current.GetFS().NewPath(to)
	} else {
		next = current.GetChild(to)
	}

	path := next.GetPath()

	if !next.IsDirectory() {
		return errors.New("the path is not a directory, path=" + path)
	}

	task.WD = path
	return nil
}
