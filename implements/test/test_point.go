package test

import (
	"github.com/starter-go/cli"
	"github.com/starter-go/units"
)

// TestPoint ...
type TestPoint struct {

	//starter:component

	CLI cli.CLI //starter:inject("#")
}

func (inst *TestPoint) _Impl() units.Units {
	return inst
}

// Units ...
func (inst *TestPoint) Units(list []*units.Registration) []*units.Registration {

	r1 := &units.Registration{
		Name:     "test-point-1",
		Enabled:  true,
		Priority: 0,
		Test:     inst.run,
	}

	list = append(list, r1)
	return list
}

func (inst *TestPoint) run() error {
	c := inst.CLI
	// client := c.GetClient()
	ctx := c.Bind(nil)
	return c.GetClient().RunCCA(ctx, "pwd", nil)
}
