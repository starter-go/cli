package handlers

import (
	"errors"
	"strings"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
	"github.com/starter-go/cli"
)

// ChdirHandler ...
type ChdirHandler struct {

	//starter:component
	_as func(cli.HandlerRegistry) //starter:as(".")

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
	} else if strings.HasPrefix(dst, "file:/") {
		next = inst.getTargetWithFileURL(current, dst)
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

// getTargetWithFileURL 解析 "file:/" 形式的路径
func (inst *ChdirHandler) getTargetWithFileURL(current afs.Path, to string) afs.Path {
	const prefix = "file:/"
	to = to[len(prefix):]
	fs1 := current.GetFS()
	sep1 := fs1.Separator()
	sep := ""
	builder := strings.Builder{}
	elements := strings.Split(to, "/")
	for _, part := range elements {
		if part == "" {
			continue
		}
		builder.WriteString(sep)
		builder.WriteString(part)
		sep = sep1
	}
	path := builder.String()
	if sep1 == "/" {
		path = "/" + path
	}
	return fs1.NewPath(path)
}
