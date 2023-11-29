package extention4cli

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p1eae9dc6c0_handlers_ChdirHandler{})
    inst.register(&p1eae9dc6c0_handlers_ExampleHandler{})
    inst.register(&p1eae9dc6c0_handlers_HelpHandler{})
    inst.register(&p1eae9dc6c0_handlers_LsHandler{})
    inst.register(&p1eae9dc6c0_handlers_MkdirHandler{})
    inst.register(&p1eae9dc6c0_handlers_NowHandler{})
    inst.register(&p1eae9dc6c0_handlers_PwdHandler{})
    inst.register(&p1eae9dc6c0_handlers_SleepHandler{})
    inst.register(&paa85bf66ed_extention_Demo{})
    inst.register(&pce9319ab50_filters_BindingFilter{})
    inst.register(&pce9319ab50_filters_ClientServerInjectingFilter{})
    inst.register(&pce9319ab50_filters_CommandPrepareFilter{})
    inst.register(&pce9319ab50_filters_ErrorFilter{})
    inst.register(&pce9319ab50_filters_ExampleFilter{})
    inst.register(&pce9319ab50_filters_HandlingFilter{})
    inst.register(&pce9319ab50_filters_MultilineCommandFilter{})


    return nil
}
