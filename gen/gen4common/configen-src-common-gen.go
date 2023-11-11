package gen4common
import (
    pd0bd963d6 "github.com/starter-go/cli/config/common"
     "github.com/starter-go/application"
)

// type pd0bd963d6.Handlers in package:github.com/starter-go/cli/config/common
//
// id:com-d0bd963d6723799c-common-Handlers
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type pd0bd963d67_common_Handlers struct {
}

func (inst* pd0bd963d67_common_Handlers) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-d0bd963d6723799c-common-Handlers"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pd0bd963d67_common_Handlers) new() any {
    return &pd0bd963d6.Handlers{}
}

func (inst* pd0bd963d67_common_Handlers) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pd0bd963d6.Handlers)
	nop(ie, com)

	


    return nil
}


