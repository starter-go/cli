// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfgexample

import (
	cli0xf7c71e "bitwormhole.com/starter/cli"
	example0x090997 "bitwormhole.com/starter/cli/config/example"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComTestPoint struct {
	instance *example0x090997.TestPoint
	 markup0x23084a.Component `class:"life"`
	CLI cli0xf7c71e.CLI `inject:"#cli"`
}

