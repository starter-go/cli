package handlers

import (
	"errors"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
	"github.com/starter-go/cli"
)

// MkdirHandler ...
type MkdirHandler struct {
}

func (inst *MkdirHandler) _Impl() (cli.Handler, cli.HandlerRegistry, cli.Help) {
	return inst, inst, inst
}

// GetHandlers ...
func (inst *MkdirHandler) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "mkdir",
		Handler: inst.Run,
		Help:    inst,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *MkdirHandler) GetHelp() *cli.HelpInfo {
	info := &cli.HelpInfo{
		Name:    "mkdir",
		Title:   "make directory",
		Usage:   "todo...",
		Content: "todo...",
	}
	return info
}

// Run ...
func (inst *MkdirHandler) Run(task *cli.Task) error {

	path := task.Arguments[0]
	wd := files.FS().NewPath(task.WD)
	child := wd.GetChild(path)
	parent := child.GetParent()

	if !parent.IsDirectory() {
		return errors.New("the path is not a directory, path=" + parent.GetPath())
	}

	if child.Exists() {
		return errors.New("the file/dir is exists, path=" + child.GetPath())
	}

	opt := &afs.Options{
		Create: true,
	}
	return child.Mkdir(opt)
}
