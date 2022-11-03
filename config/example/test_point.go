package example

import (
	"bitwormhole.com/starter/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
)

// TestPoint ...
type TestPoint struct {
	markup.Component `class:"life"`

	CLI cli.CLI `inject:"#cli"`
}

func (inst *TestPoint) _Impl() application.LifeRegistry {
	return inst
}

// GetLifeRegistration ...
func (inst *TestPoint) GetLifeRegistration() *application.LifeRegistration {
	return &application.LifeRegistration{
		OnStart: inst.start,
	}
}

func (inst *TestPoint) start() error {
	cli := inst.CLI
	client := cli.GetClient()
	ctx := cli.Bind(nil)
	return client.RunCCA(ctx, "pwd", nil)
}
