package cli

import (
	"errors"
	"strings"
)

type serverImpl struct {
	handlers map[string]*HandlerRegistration
}

func (inst *serverImpl) _Impl() Server {
	return inst
}

func (inst *serverImpl) makeNoHandlerError(name string) (*HandlerRegistration, error) {
	return nil, errors.New("no handler for command:" + name)
}

func (inst *serverImpl) FindHandler(name string) (*HandlerRegistration, error) {

	name = strings.TrimSpace(name)

	table := inst.handlers
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
