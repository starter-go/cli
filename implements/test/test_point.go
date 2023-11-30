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
		Name:     "test-point-1-1",
		Enabled:  true,
		Priority: 0,
		Test:     inst.run1,
	}

	r2 := &units.Registration{
		Name:     "test-point-1-2",
		Enabled:  true,
		Priority: 0,
		Test:     inst.run2,
	}

	r3 := &units.Registration{
		Name:     "test-point-1-3",
		Enabled:  true,
		Priority: 100,
		Test:     inst.run3,
	}

	r4 := &units.Registration{
		Name:     "test-point-1-4",
		Enabled:  true,
		Priority: 0,
		Test:     inst.run4,
	}

	list = append(list, r1)
	list = append(list, r2)
	list = append(list, r3)
	list = append(list, r4)

	return list
}

func (inst *TPoint1) runCommand(cmd string) error {
	c := inst.CLI
	ctx := c.Bind(nil)
	return c.GetClient().RunCCA(ctx, cmd, nil)
}

func (inst *TPoint1) run1() error {
	return inst.runCommand("pwd")
}

func (inst *TPoint1) run2() error {
	return inst.runCommand("now")
}

func (inst *TPoint1) run3() error {
	return inst.runCommand("help")
}

func (inst *TPoint1) run4() error {
	// return inst.runCommand("run git status  ")
	return inst.runCommand("run 'C:\\Program Files\\Mozilla Firefox\\firefox.exe'  https://www.cctv.com/ ")
}
