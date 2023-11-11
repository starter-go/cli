package gen4common

import "github.com/starter-go/application"

//starter:configen(version="4")

// ComForCLICommon ...
func ComForCLICommon(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
