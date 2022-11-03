// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfglib

import (
	cli0xf7c71e "bitwormhole.com/starter/cli"
	lib0x973f62 "bitwormhole.com/starter/cli/config/lib"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: cli
	cominfobuilder.Next()
	cominfobuilder.ID("cli").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComComCLI{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	// component: com1-lib0x973f62.Filters
	cominfobuilder.Next()
	cominfobuilder.ID("com1-lib0x973f62.Filters").Class("cli-filter-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComFilters{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComComCLI : the factory of component: cli
type comFactory4pComComCLI struct {

    mPrototype * lib0x973f62.ComCLI

	
	mHandlersSelector config.InjectionSelector
	mFiltersSelector config.InjectionSelector

}

func (inst * comFactory4pComComCLI) init() application.ComponentFactory {

	
	inst.mHandlersSelector = config.NewInjectionSelector(".cli-handler-registry",nil)
	inst.mFiltersSelector = config.NewInjectionSelector(".cli-filter-registry",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComComCLI) newObject() * lib0x973f62.ComCLI {
	return & lib0x973f62.ComCLI {}
}

func (inst * comFactory4pComComCLI) castObject(instance application.ComponentInstance) * lib0x973f62.ComCLI {
	return instance.Get().(*lib0x973f62.ComCLI)
}

func (inst * comFactory4pComComCLI) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComComCLI) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComComCLI) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComComCLI) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComComCLI) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComComCLI) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Handlers = inst.getterForFieldHandlersSelector(context)
	obj.Filters = inst.getterForFieldFiltersSelector(context)
	return context.LastError()
}

//getterForFieldHandlersSelector
func (inst * comFactory4pComComCLI) getterForFieldHandlersSelector (context application.InstanceContext) []cli0xf7c71e.HandlerRegistry {
	list1 := inst.mHandlersSelector.GetList(context)
	list2 := make([]cli0xf7c71e.HandlerRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(cli0xf7c71e.HandlerRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}

//getterForFieldFiltersSelector
func (inst * comFactory4pComComCLI) getterForFieldFiltersSelector (context application.InstanceContext) []cli0xf7c71e.FilterRegistry {
	list1 := inst.mFiltersSelector.GetList(context)
	list2 := make([]cli0xf7c71e.FilterRegistry, 0, len(list1))
	for _, item1 := range list1 {
		item2, ok := item1.(cli0xf7c71e.FilterRegistry)
		if ok {
			list2 = append(list2, item2)
		}
	}
	return list2
}



////////////////////////////////////////////////////////////////////////////////

// comFactory4pComFilters : the factory of component: com1-lib0x973f62.Filters
type comFactory4pComFilters struct {

    mPrototype * lib0x973f62.Filters

	

}

func (inst * comFactory4pComFilters) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComFilters) newObject() * lib0x973f62.Filters {
	return & lib0x973f62.Filters {}
}

func (inst * comFactory4pComFilters) castObject(instance application.ComponentInstance) * lib0x973f62.Filters {
	return instance.Get().(*lib0x973f62.Filters)
}

func (inst * comFactory4pComFilters) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComFilters) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComFilters) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComFilters) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFilters) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComFilters) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




