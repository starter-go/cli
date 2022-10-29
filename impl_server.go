package cli

import (
	"errors"
	"strings"

	"bitwormhole.com/starter/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type serverImpl struct {
	context  *Context
	handlers map[string]*HandlerRegistration // cache
}

func (inst *serverImpl) _Impl() Server {
	return inst
}

func (inst *serverImpl) makeNoHandlerError(name string) (*HandlerRegistration, error) {
	return nil, errors.New("no handler for command:" + name)
}

func (inst *serverImpl) getTab() map[string]*HandlerRegistration {
	t := inst.handlers
	if t == nil {
		t = make(map[string]*HandlerRegistration)
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

func (inst *serverImpl) FindHandler(name string) (*HandlerRegistration, error) {

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

	hr2 := &HandlerRegistration{}
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

func (inst *serverImpl) RegisterHandler(hr *HandlerRegistration) error {
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

func (inst *serverImpl) RegisterHandlers(hr HandlerRegistry) error {
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

func (inst *DefaultServerFactory) _Impl() ServerFactory {
	return inst
}

// NewServer ...
func (inst *DefaultServerFactory) NewServer(c *Context) (Server, error) {
	server := &serverImpl{
		context: c,
	}
	return server, nil
}

////////////////////////////////////////////////////////////////////////////////
