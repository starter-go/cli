package core

import (
	"errors"
	"strings"

	"github.com/starter-go/cli"
	"github.com/starter-go/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type serverImpl struct {
	context  *cli.Context
	handlers map[string]*cli.HandlerRegistration // cache
}

func (inst *serverImpl) _Impl() cli.Server {
	return inst
}

func (inst *serverImpl) makeNoHandlerError(name string) (*cli.HandlerRegistration, error) {
	return nil, errors.New("no handler for command:" + name)
}

func (inst *serverImpl) getTab() map[string]*cli.HandlerRegistration {
	t := inst.handlers
	if t == nil {
		t = make(map[string]*cli.HandlerRegistration)
	} else {
		return t
	}
	src := inst.context.Handlers
	for _, hr := range src {
		name := hr.Name
		old := t[name]
		if old == nil {
			t[name] = hr
		} else {
			vlog.Warn("the command name is duplicated, name=" + name)
		}
	}
	inst.handlers = t
	return t
}

func (inst *serverImpl) FindHandler(name string) (*cli.HandlerRegistration, error) {

	name = strings.TrimSpace(name)

	table := inst.getTab()
	if table == nil {
		return inst.makeNoHandlerError(name)
	}

	hr := table[name]
	if hr == nil {
		return inst.makeNoHandlerError(name)
	}

	if hr.Handler == nil {
		return inst.makeNoHandlerError(name)
	}

	hr2 := &cli.HandlerRegistration{}
	*hr2 = *hr
	return hr2, nil
}

func (inst *serverImpl) ListNames() []string {
	src := inst.handlers
	dst := make([]string, 0)
	for name := range src {
		dst = append(dst, name)
	}
	return dst
}

func (inst *serverImpl) RegisterHandler(hr *cli.HandlerRegistration) error {
	if hr == nil {
		return nil
	}
	name := hr.Name
	dst := inst.getTab()
	old := dst[name]
	if old != nil {
		return errors.New("the command name is duplicated, name=" + name)
	}
	dst[name] = hr
	return nil
}

func (inst *serverImpl) RegisterHandlers(hr cli.HandlerRegistry) error {
	src := hr.GetHandlers()
	for _, hr2 := range src {
		err := inst.RegisterHandler(hr2)
		if err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// DefaultServerFactory ...
type DefaultServerFactory struct {
}

func (inst *DefaultServerFactory) _Impl() cli.ServerFactory {
	return inst
}

// NewServer ...
func (inst *DefaultServerFactory) NewServer(c *cli.Context) (cli.Server, error) {
	server := &serverImpl{
		context: c,
	}
	return server, nil
}

////////////////////////////////////////////////////////////////////////////////
