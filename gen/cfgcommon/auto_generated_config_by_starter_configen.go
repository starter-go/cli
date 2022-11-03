// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfgcommon

import (
	common0x169b46 "bitwormhole.com/starter/cli/config/common"
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

	// component: com0-common0x169b46.Handlers
	cominfobuilder.Next()
	cominfobuilder.ID("com0-common0x169b46.Handlers").Class("cli-handler-registry").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComHandlers{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComHandlers : the factory of component: com0-common0x169b46.Handlers
type comFactory4pComHandlers struct {

    mPrototype * common0x169b46.Handlers

	

}

func (inst * comFactory4pComHandlers) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComHandlers) newObject() * common0x169b46.Handlers {
	return & common0x169b46.Handlers {}
}

func (inst * comFactory4pComHandlers) castObject(instance application.ComponentInstance) * common0x169b46.Handlers {
	return instance.Get().(*common0x169b46.Handlers)
}

func (inst * comFactory4pComHandlers) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComHandlers) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComHandlers) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComHandlers) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHandlers) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComHandlers) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}




