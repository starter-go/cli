package gunit

import (
	"testing"

	"bitwormhole.com/starter/cli/commands"
)

func Test(t *testing.T) {
	cb := commands.DefaultContextBuilder()
	ctx := cb.Create()
	err := ctx.Client.RunCCA(nil, "help", nil)
	if err != nil {
		t.Error(err)
	}
}
