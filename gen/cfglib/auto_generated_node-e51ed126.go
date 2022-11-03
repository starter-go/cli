// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cfglib

import (
	cli0xf7c71e "bitwormhole.com/starter/cli"
	lib0x973f62 "bitwormhole.com/starter/cli/config/lib"
	markup0x23084a "github.com/bitwormhole/starter/markup"
)

type pComComCLI struct {
	instance *lib0x973f62.ComCLI
	 markup0x23084a.Component `id:"cli" class:"life"`
	Handlers []cli0xf7c71e.HandlerRegistry `inject:".cli-handler-registry"`
	Filters []cli0xf7c71e.FilterRegistry `inject:".cli-filter-registry"`
}


type pComFilters struct {
	instance *lib0x973f62.Filters
	 markup0x23084a.Component `class:"cli-filter-registry"`
}

