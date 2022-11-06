package gunit

import (
	"strings"
	"testing"
)

func TestMultiline(t *testing.T) {

	const nl = "\n"
	cmd := strings.Builder{}

	cmd.WriteString("pwd" + nl)
	cmd.WriteString("sleep 200" + nl)
	cmd.WriteString("help" + nl)
	cmd.WriteString("now" + nl)
	cmd.WriteString("ls" + nl)

	testCommand(t, cmd.String())
}
