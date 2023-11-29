package test

import (
	"github.com/starter-go/cli"
	"github.com/starter-go/units"
)

// TPoint1 ... 测试点：No.1
type TPoint1 struct {

	//starter:component

	CLI cli.CLI //starter:inject("#")
}

func (inst *TPoint1) _Impl() units.Units {
	return inst
}

// Units ...
func (inst *TPoint1) Units(list []*units.Registration) []*units.Registration {

	r1 := &units.Registration{
		Name:     "test-point-1",
		Enabled:  true,
		Priority: 0,
		Test:     inst.run,
	}

	list = append(list, r1)
	return list
}

func (inst *TPoint1) run() error {
	c := inst.CLI
	// client := c.GetClient()
	ctx := c.Bind(nil)
	return c.GetClient().RunCCA(ctx, "pwd", nil)
}
