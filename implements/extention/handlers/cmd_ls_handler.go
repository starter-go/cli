package handlers

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
	"github.com/starter-go/cli"
)

// LsHandler ...
type LsHandler struct {

	//starter:component
	_as func(cli.HandlerRegistry) //starter:as(".")

}

func (inst *LsHandler) _Impl() (cli.Handler, cli.HandlerRegistry, cli.Help) {
	return inst, inst, inst
}

// GetHandlers ...
func (inst *LsHandler) GetHandlers() []*cli.HandlerRegistration {
	hr := &cli.HandlerRegistration{
		Name:    "ls",
		Handler: inst.Run,
		Help:    inst,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *LsHandler) GetHelp() *cli.HelpInfo {
	info := &cli.HelpInfo{
		Name:    "ls",
		Title:   "Ls",
		Usage:   "todo...",
		Content: "todo...",
	}
	return info
}

// Run ...
func (inst *LsHandler) Run(task *cli.Task) error {

	wd := task.WD
	out := task.Console.Out()
	keys := make([]string, 0)
	table := make(map[string]afs.Path)
	builder := strings.Builder{}

	list, err := inst.getFileList(wd)
	if err != nil {
		return err
	}

	for _, child := range list {
		name := child.GetName()
		table[name] = child
		keys = append(keys, name)
	}

	sort.Strings(keys)
	out.WriteString("|updated_at|type|size|name|\n")
	out.WriteString("|----------|----|----|----|\n")

	for i, name := range keys {
		item := table[name]
		builder.Reset()
		inst.formatItem(item, i, &builder)
		out.WriteString(builder.String())
	}

	return nil
}

func (inst *LsHandler) formatItem(item afs.Path, index int, b *strings.Builder) {

	name := item.GetName()
	info := item.GetInfo()
	size := info.Length()
	date := info.UpdatedAt()
	dir := info.IsDirectory()

	strSize := strconv.FormatInt(size, 10)
	strDate := inst.formatDate(date)
	strDir := ""
	if dir {
		strDir = "[DIR]"
	}

	// 01234567890123456789012345
	// 2022-10-26T13:24:42+08:00

	inst.formatColumn(strDate, 26, 'l', b)
	inst.formatColumn(strDir, 6, 'l', b)
	inst.formatColumn(strSize, 10, 'r', b)
	b.WriteString("  " + name)
	b.WriteRune('\n')
}

func (inst *LsHandler) formatDate(date time.Time) string {
	str := date.Format(time.RFC3339)
	i := strings.IndexRune(str, '+')
	if i > 0 {
		str = str[0:i]
	}
	return strings.Replace(str, "T", " ", 1)
}

func (inst *LsHandler) formatColumn(str string, width int, align rune, out *strings.Builder) {

	const space = ' '
	size := len(str)
	paddingSize := width - size

	if align == 'l' {
		out.WriteString(str)
	}

	if paddingSize > 0 {
		for i := 0; i < paddingSize; i++ {
			out.WriteRune(space)
		}
	} else {
		out.WriteRune(space)
	}

	if align == 'r' {
		out.WriteString(str)
	}
}

func (inst *LsHandler) getFileList(wd string) ([]afs.Path, error) {

	f := files.FS()
	dir := f.NewPath(wd)

	for dir != nil {
		if dir.IsDirectory() {
			break
		} else {
			dir = dir.GetParent()
		}
	}

	if dir == nil {
		return nil, errors.New("The current working directory is inaccessible, path=" + wd)
	}

	list := dir.ListChildren()
	return list, nil
}
