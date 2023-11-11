package gen4example

import "github.com/starter-go/application"

//starter:configen(version="4")

// ComForCLIExample ...
func ComForCLIExample(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
