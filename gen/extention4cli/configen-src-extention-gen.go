package extention4cli
import (
    pce9319ab5 "github.com/starter-go/cli/implements/extention/filters"
    p1eae9dc6c "github.com/starter-go/cli/implements/extention/handlers"
     "github.com/starter-go/application"
)

// type pce9319ab5.BindingFilter in package:github.com/starter-go/cli/implements/extention/filters
//
// id:com-ce9319ab50fd5480-filters-BindingFilter
// class:class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry
// alias:
// scope:singleton
//
type pce9319ab50_filters_BindingFilter struct {
}

func (inst* pce9319ab50_filters_BindingFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ce9319ab50fd5480-filters-BindingFilter"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pce9319ab50_filters_BindingFilter) new() any {
    return &pce9319ab5.BindingFilter{}
}

func (inst* pce9319ab50_filters_BindingFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pce9319ab5.BindingFilter)
	nop(ie, com)

	


    return nil
}



// type pce9319ab5.ClientServerInjectingFilter in package:github.com/starter-go/cli/implements/extention/filters
//
// id:com-ce9319ab50fd5480-filters-ClientServerInjectingFilter
// class:class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry
// alias:
// scope:singleton
//
type pce9319ab50_filters_ClientServerInjectingFilter struct {
}

func (inst* pce9319ab50_filters_ClientServerInjectingFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ce9319ab50fd5480-filters-ClientServerInjectingFilter"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pce9319ab50_filters_ClientServerInjectingFilter) new() any {
    return &pce9319ab5.ClientServerInjectingFilter{}
}

func (inst* pce9319ab50_filters_ClientServerInjectingFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pce9319ab5.ClientServerInjectingFilter)
	nop(ie, com)

	


    return nil
}



// type pce9319ab5.CommandPrepareFilter in package:github.com/starter-go/cli/implements/extention/filters
//
// id:com-ce9319ab50fd5480-filters-CommandPrepareFilter
// class:class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry
// alias:
// scope:singleton
//
type pce9319ab50_filters_CommandPrepareFilter struct {
}

func (inst* pce9319ab50_filters_CommandPrepareFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ce9319ab50fd5480-filters-CommandPrepareFilter"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pce9319ab50_filters_CommandPrepareFilter) new() any {
    return &pce9319ab5.CommandPrepareFilter{}
}

func (inst* pce9319ab50_filters_CommandPrepareFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pce9319ab5.CommandPrepareFilter)
	nop(ie, com)

	


    return nil
}



// type pce9319ab5.ErrorFilter in package:github.com/starter-go/cli/implements/extention/filters
//
// id:com-ce9319ab50fd5480-filters-ErrorFilter
// class:class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry
// alias:
// scope:singleton
//
type pce9319ab50_filters_ErrorFilter struct {
}

func (inst* pce9319ab50_filters_ErrorFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ce9319ab50fd5480-filters-ErrorFilter"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pce9319ab50_filters_ErrorFilter) new() any {
    return &pce9319ab5.ErrorFilter{}
}

func (inst* pce9319ab50_filters_ErrorFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pce9319ab5.ErrorFilter)
	nop(ie, com)

	


    return nil
}



// type pce9319ab5.ExampleFilter in package:github.com/starter-go/cli/implements/extention/filters
//
// id:com-ce9319ab50fd5480-filters-ExampleFilter
// class:class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry
// alias:
// scope:singleton
//
type pce9319ab50_filters_ExampleFilter struct {
}

func (inst* pce9319ab50_filters_ExampleFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ce9319ab50fd5480-filters-ExampleFilter"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pce9319ab50_filters_ExampleFilter) new() any {
    return &pce9319ab5.ExampleFilter{}
}

func (inst* pce9319ab50_filters_ExampleFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pce9319ab5.ExampleFilter)
	nop(ie, com)

	


    return nil
}



// type pce9319ab5.HandlingFilter in package:github.com/starter-go/cli/implements/extention/filters
//
// id:com-ce9319ab50fd5480-filters-HandlingFilter
// class:class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry
// alias:
// scope:singleton
//
type pce9319ab50_filters_HandlingFilter struct {
}

func (inst* pce9319ab50_filters_HandlingFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ce9319ab50fd5480-filters-HandlingFilter"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pce9319ab50_filters_HandlingFilter) new() any {
    return &pce9319ab5.HandlingFilter{}
}

func (inst* pce9319ab50_filters_HandlingFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pce9319ab5.HandlingFilter)
	nop(ie, com)

	


    return nil
}



// type pce9319ab5.MultilineCommandFilter in package:github.com/starter-go/cli/implements/extention/filters
//
// id:com-ce9319ab50fd5480-filters-MultilineCommandFilter
// class:class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry
// alias:
// scope:singleton
//
type pce9319ab50_filters_MultilineCommandFilter struct {
}

func (inst* pce9319ab50_filters_MultilineCommandFilter) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ce9319ab50fd5480-filters-MultilineCommandFilter"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-FilterRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pce9319ab50_filters_MultilineCommandFilter) new() any {
    return &pce9319ab5.MultilineCommandFilter{}
}

func (inst* pce9319ab50_filters_MultilineCommandFilter) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pce9319ab5.MultilineCommandFilter)
	nop(ie, com)

	


    return nil
}



// type p1eae9dc6c.ChdirHandler in package:github.com/starter-go/cli/implements/extention/handlers
//
// id:com-1eae9dc6c0c30769-handlers-ChdirHandler
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type p1eae9dc6c0_handlers_ChdirHandler struct {
}

func (inst* p1eae9dc6c0_handlers_ChdirHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eae9dc6c0c30769-handlers-ChdirHandler"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eae9dc6c0_handlers_ChdirHandler) new() any {
    return &p1eae9dc6c.ChdirHandler{}
}

func (inst* p1eae9dc6c0_handlers_ChdirHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eae9dc6c.ChdirHandler)
	nop(ie, com)

	


    return nil
}



// type p1eae9dc6c.ExampleHandler in package:github.com/starter-go/cli/implements/extention/handlers
//
// id:com-1eae9dc6c0c30769-handlers-ExampleHandler
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type p1eae9dc6c0_handlers_ExampleHandler struct {
}

func (inst* p1eae9dc6c0_handlers_ExampleHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eae9dc6c0c30769-handlers-ExampleHandler"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eae9dc6c0_handlers_ExampleHandler) new() any {
    return &p1eae9dc6c.ExampleHandler{}
}

func (inst* p1eae9dc6c0_handlers_ExampleHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eae9dc6c.ExampleHandler)
	nop(ie, com)

	


    return nil
}



// type p1eae9dc6c.HelpHandler in package:github.com/starter-go/cli/implements/extention/handlers
//
// id:com-1eae9dc6c0c30769-handlers-HelpHandler
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type p1eae9dc6c0_handlers_HelpHandler struct {
}

func (inst* p1eae9dc6c0_handlers_HelpHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eae9dc6c0c30769-handlers-HelpHandler"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eae9dc6c0_handlers_HelpHandler) new() any {
    return &p1eae9dc6c.HelpHandler{}
}

func (inst* p1eae9dc6c0_handlers_HelpHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eae9dc6c.HelpHandler)
	nop(ie, com)

	


    return nil
}



// type p1eae9dc6c.LsHandler in package:github.com/starter-go/cli/implements/extention/handlers
//
// id:com-1eae9dc6c0c30769-handlers-LsHandler
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type p1eae9dc6c0_handlers_LsHandler struct {
}

func (inst* p1eae9dc6c0_handlers_LsHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eae9dc6c0c30769-handlers-LsHandler"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eae9dc6c0_handlers_LsHandler) new() any {
    return &p1eae9dc6c.LsHandler{}
}

func (inst* p1eae9dc6c0_handlers_LsHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eae9dc6c.LsHandler)
	nop(ie, com)

	


    return nil
}



// type p1eae9dc6c.MkdirHandler in package:github.com/starter-go/cli/implements/extention/handlers
//
// id:com-1eae9dc6c0c30769-handlers-MkdirHandler
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type p1eae9dc6c0_handlers_MkdirHandler struct {
}

func (inst* p1eae9dc6c0_handlers_MkdirHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eae9dc6c0c30769-handlers-MkdirHandler"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eae9dc6c0_handlers_MkdirHandler) new() any {
    return &p1eae9dc6c.MkdirHandler{}
}

func (inst* p1eae9dc6c0_handlers_MkdirHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eae9dc6c.MkdirHandler)
	nop(ie, com)

	


    return nil
}



// type p1eae9dc6c.NowHandler in package:github.com/starter-go/cli/implements/extention/handlers
//
// id:com-1eae9dc6c0c30769-handlers-NowHandler
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type p1eae9dc6c0_handlers_NowHandler struct {
}

func (inst* p1eae9dc6c0_handlers_NowHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eae9dc6c0c30769-handlers-NowHandler"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eae9dc6c0_handlers_NowHandler) new() any {
    return &p1eae9dc6c.NowHandler{}
}

func (inst* p1eae9dc6c0_handlers_NowHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eae9dc6c.NowHandler)
	nop(ie, com)

	


    return nil
}



// type p1eae9dc6c.PwdHandler in package:github.com/starter-go/cli/implements/extention/handlers
//
// id:com-1eae9dc6c0c30769-handlers-PwdHandler
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type p1eae9dc6c0_handlers_PwdHandler struct {
}

func (inst* p1eae9dc6c0_handlers_PwdHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eae9dc6c0c30769-handlers-PwdHandler"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eae9dc6c0_handlers_PwdHandler) new() any {
    return &p1eae9dc6c.PwdHandler{}
}

func (inst* p1eae9dc6c0_handlers_PwdHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eae9dc6c.PwdHandler)
	nop(ie, com)

	


    return nil
}



// type p1eae9dc6c.RunHandler in package:github.com/starter-go/cli/implements/extention/handlers
//
// id:com-1eae9dc6c0c30769-handlers-RunHandler
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type p1eae9dc6c0_handlers_RunHandler struct {
}

func (inst* p1eae9dc6c0_handlers_RunHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eae9dc6c0c30769-handlers-RunHandler"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eae9dc6c0_handlers_RunHandler) new() any {
    return &p1eae9dc6c.RunHandler{}
}

func (inst* p1eae9dc6c0_handlers_RunHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eae9dc6c.RunHandler)
	nop(ie, com)

	


    return nil
}



// type p1eae9dc6c.SleepHandler in package:github.com/starter-go/cli/implements/extention/handlers
//
// id:com-1eae9dc6c0c30769-handlers-SleepHandler
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type p1eae9dc6c0_handlers_SleepHandler struct {
}

func (inst* p1eae9dc6c0_handlers_SleepHandler) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1eae9dc6c0c30769-handlers-SleepHandler"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1eae9dc6c0_handlers_SleepHandler) new() any {
    return &p1eae9dc6c.SleepHandler{}
}

func (inst* p1eae9dc6c0_handlers_SleepHandler) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1eae9dc6c.SleepHandler)
	nop(ie, com)

	


    return nil
}


