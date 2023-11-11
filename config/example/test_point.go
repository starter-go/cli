package example

import (
	"github.com/starter-go/application"
	"github.com/starter-go/cli"
)

// TestPoint ...
type TestPoint struct {

	//starter:component
	_as func(application.Lifecycle) //starter:as(".")

	CLI cli.CLI //starter:inject("#")
}

func (inst *TestPoint) _Impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *TestPoint) Life() *application.Life {
	return &application.Life{
		OnStart: inst.start,
	}
}

func (inst *TestPoint) start() error {
	cli := inst.CLI
	client := cli.GetClient()
	ctx := cli.Bind(nil)
	return client.RunCCA(ctx, "pwd", nil)
}
