// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfgexample

import (
	cli0xf7c71e "bitwormhole.com/starter/cli"
	example0x090997 "bitwormhole.com/starter/cli/config/example"
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

	// component: com0-example0x090997.TestPoint
	cominfobuilder.Next()
	cominfobuilder.ID("com0-example0x090997.TestPoint").Class("life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComTestPoint{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComTestPoint : the factory of component: com0-example0x090997.TestPoint
type comFactory4pComTestPoint struct {

    mPrototype * example0x090997.TestPoint

	
	mCLISelector config.InjectionSelector

}

func (inst * comFactory4pComTestPoint) init() application.ComponentFactory {

	
	inst.mCLISelector = config.NewInjectionSelector("#cli",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComTestPoint) newObject() * example0x090997.TestPoint {
	return & example0x090997.TestPoint {}
}

func (inst * comFactory4pComTestPoint) castObject(instance application.ComponentInstance) * example0x090997.TestPoint {
	return instance.Get().(*example0x090997.TestPoint)
}

func (inst * comFactory4pComTestPoint) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComTestPoint) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComTestPoint) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComTestPoint) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestPoint) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComTestPoint) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.CLI = inst.getterForFieldCLISelector(context)
	return context.LastError()
}

//getterForFieldCLISelector
func (inst * comFactory4pComTestPoint) getterForFieldCLISelector (context application.InstanceContext) cli0xf7c71e.CLI {

	o1 := inst.mCLISelector.GetOne(context)
	o2, ok := o1.(cli0xf7c71e.CLI)
	if !ok {
		eb := &util.ErrorBuilder{}
		eb.Message("bad cast")
		eb.Set("com", "com0-example0x090997.TestPoint")
		eb.Set("field", "CLI")
		eb.Set("type1", "?")
		eb.Set("type2", "cli0xf7c71e.CLI")
		context.HandleError(eb.Create())
		return nil
	}
	return o2
}




