package cli

import (
	"context"
	"io"
	"os"
)

// Binding ...
type Binding struct {
	WD      string
	Env     map[string]string
	Error   io.Writer
	Output  io.Writer
	Input   io.Reader
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

	// make new Binding
	b = &Binding{}
	c2 := context.WithValue(c, key, b)
	b.Context = c2
	b.Error = os.Stderr
	b.Output = os.Stdout
	b.Input = os.Stdin
	b.WD = wd

	return b
}

////////////////////////////////////////////////////////////////////////////////
