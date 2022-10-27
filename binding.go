package cli

import (
	"context"
	"os"
)

// Binding ...
type Binding struct {
	WD      string
	Env     map[string]string
	Console Console
	Context context.Context
}

////////////////////////////////////////////////////////////////////////////////

type tagBindingContextKey string

// GetBinding ...
func GetBinding(c context.Context) *Binding {

	const key tagBindingContextKey = "bitwormhole.com/starter/cli/Binding#binding"

	o := c.Value(key)
	b, ok := o.(*Binding)
	if ok {
		return b
	}

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
	b = &Binding{}
	c2 := context.WithValue(c, key, b)
	b.Context = c2
	b.WD = wd
	b.Console = cb.Create()

	return b
}

////////////////////////////////////////////////////////////////////////////////
