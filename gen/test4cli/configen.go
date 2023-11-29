package test4cli

import "github.com/starter-go/application"

//starter:configen(version="4")

// Components ...
func Components(cr application.ComponentRegistry) error {
	return registerComponents(cr)
	// return nil
}
