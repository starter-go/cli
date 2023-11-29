package handlers

import (
	"strconv"
	"strings"
	"time"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/cli"
)

// NowHandler ...
type NowHandler struct {
}

func (inst *NowHandler) _Impl() (cli.Handler, cli.HandlerRegistry, cli.Help) {
	return inst, inst, inst
}

func (inst *NowHandler) name() string {
	return "now"
}

// GetHandlers ...
func (inst *NowHandler) GetHandlers() []*cli.HandlerRegistration {
	name := inst.name()
	hr := &cli.HandlerRegistration{
		Name:    name,
		Handler: inst.Run,
		Help:    inst,
	}
	return []*cli.HandlerRegistration{hr}
}

// GetHelp ...
func (inst *NowHandler) GetHelp() *cli.HelpInfo {
	name := inst.name()
	info := &cli.HelpInfo{
		Name:    name,
		Title:   "Now",
		Usage:   "todo...",
		Content: "todo...",
	}
	return info
}

// Run ...
func (inst *NowHandler) Run(task *cli.Task) error {
	now := lang.Now()
	t := now.Int() // util.NewTime(now).Int64()

	builder := strings.Builder{}
	builder.WriteString(now.Time().Format(time.RFC3339))
	builder.WriteString(" (t=")
	builder.WriteString(strconv.FormatInt(t, 10))
	builder.WriteString(")\n")
	out := task.Console.Out()
	out.WriteString(builder.String())
	return nil
}
