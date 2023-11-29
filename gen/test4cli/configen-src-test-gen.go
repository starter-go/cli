package test4cli
import (
    p1336d65ed "github.com/starter-go/cli"
    p4d9a94685 "github.com/starter-go/cli/implements/test"
     "github.com/starter-go/application"
)

// type p4d9a94685.TPoint1 in package:github.com/starter-go/cli/implements/test
//
// id:com-4d9a94685042a0d8-test-TPoint1
// class:
// alias:
// scope:singleton
//
type p4d9a946850_test_TPoint1 struct {
}

func (inst* p4d9a946850_test_TPoint1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-4d9a94685042a0d8-test-TPoint1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p4d9a946850_test_TPoint1) new() any {
    return &p4d9a94685.TPoint1{}
}

func (inst* p4d9a946850_test_TPoint1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p4d9a94685.TPoint1)
	nop(ie, com)

	
    com.CLI = inst.getCLI(ie)


    return nil
}


func (inst*p4d9a946850_test_TPoint1) getCLI(ie application.InjectionExt)p1336d65ed.CLI{
    return ie.GetComponent("#alias-1336d65edeed550b78a5d5b61e92d726-CLI").(p1336d65ed.CLI)
}


