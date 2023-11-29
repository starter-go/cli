package extention4cli
import (
    paa85bf66e "github.com/starter-go/cli/implements/extention"
     "github.com/starter-go/application"
)

// type paa85bf66e.Demo in package:github.com/starter-go/cli/implements/extention
//
// id:com-aa85bf66edb60b91-extention-Demo
// class:
// alias:
// scope:singleton
//
type paa85bf66ed_extention_Demo struct {
}

func (inst* paa85bf66ed_extention_Demo) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-aa85bf66edb60b91-extention-Demo"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* paa85bf66ed_extention_Demo) new() any {
    return &paa85bf66e.Demo{}
}

func (inst* paa85bf66ed_extention_Demo) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*paa85bf66e.Demo)
	nop(ie, com)

	


    return nil
}


