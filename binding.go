package cli

import (
	"context"
	"errors"
	"os"
)

// Binding ...
type Binding interface {
	GetWD() string
	GetEnv() map[string]string
	GetConsole() Console

	SetWD(wd string)
	SetEnv(kv map[string]string)
	SetConsole(c Console)
}

////////////////////////////////////////////////////////////////////////////////

type tagBindingKey string

const theBindingKey tagBindingKey = "bitwormhole.com/starter/cli/Binding#binding"

// binding ...
type binding struct {
	wd      string
	env     map[string]string
	console Console
	config  *Configuration
	context *Context
	facade  Binding
}

////////////////////////////////////////////////////////////////////////////////

type bindingAccess struct{}

// setupBinding ...
func (inst *bindingAccess) setup(c context.Context) context.Context {

	if c == nil {
		c = context.Background()
	}

	const key = theBindingKey
	o := c.Value(key)
	b, ok := o.(*binding)
	if ok && b != nil {
		return c
	}

	// init new binding
	b = &binding{}
	inst.initBinding(b)
	c = context.WithValue(c, key, b)
	return c
}

// getBinding ...
func (bindingAccess) getBinding(c context.Context) (*binding, error) {
	const key = theBindingKey
	o := c.Value(key)
	b, ok := o.(*binding)
	if ok && b != nil {
		return b, nil
	}
	return nil, errors.New("need for cli.bindingAccess.setup()")
}

func (bindingAccess) initBinding(b *binding) {

	// wd
	wd, err := os.Getwd()
	if err != nil {
		wd = "/home/x"
	}

	// new console
	cb := &ConsoleBuilder{}
	cb.Error = os.Stderr
	cb.Output = os.Stdout
	cb.Input = os.Stdin

	// make new Binding
	b.wd = wd
	b.console = cb.Create()
	b.facade = &bindingFacade{inner: b}
}

////////////////////////////////////////////////////////////////////////////////

type bindingFacade struct {
	inner *binding
}

func (inst *bindingFacade) _Impl() Binding {
	return inst
}

func (inst *bindingFacade) GetWD() string {
	return inst.inner.wd
}

func (inst *bindingFacade) GetEnv() map[string]string {
	return inst.inner.env
}

func (inst *bindingFacade) GetConsole() Console {
	return inst.inner.console
}

func (inst *bindingFacade) SetWD(wd string) {
	inst.inner.wd = wd
}

func (inst *bindingFacade) SetEnv(kv map[string]string) {
	inst.inner.env = kv
}

func (inst *bindingFacade) SetConsole(c Console) {
	inst.inner.console = c
}

////////////////////////////////////////////////////////////////////////////////

// GetBinding ... get the facade of binding
func GetBinding(cc context.Context) Binding {
	ba := bindingAccess{}
	b, err := ba.getBinding(cc)
	if err != nil {
		panic(err)
	}
	return b.facade
}
