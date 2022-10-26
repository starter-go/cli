package cli

import "sort"

////////////////////////////////////////////////////////////////////////////////

// HandlerRegistration ... 定义一条命令的注册信息
type HandlerRegistration struct {
	Name    string
	Handler HandlerFunc
	Help    Help
}

// HandlerRegistry ... 该接口表示一个命令注册对象
type HandlerRegistry interface {
	GetHandlers() []*HandlerRegistration
}

////////////////////////////////////////////////////////////////////////////////

// FilterRegistration ... 定义一条命令的注册信息
type FilterRegistration struct {
	Name   string
	Filter Filter
	Order  int
}

// FilterRegistry ... 该接口表示一个命令注册对象
type FilterRegistry interface {
	GetFilters() []*FilterRegistration
}

////////////////////////////////////////////////////////////////////////////////

// FilterRegistrationSorter ...
type FilterRegistrationSorter struct {
	items   []*FilterRegistration
	reverse bool
}

// Sort ...
func (inst *FilterRegistrationSorter) Sort(list []*FilterRegistration, reverse bool) {
	child := &FilterRegistrationSorter{
		items:   list,
		reverse: reverse,
	}
	sort.Sort(child)
}

func (inst *FilterRegistrationSorter) Len() int {
	return len(inst.items)
}

func (inst *FilterRegistrationSorter) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	b := inst.reverse
	if o1.Order < o2.Order {
		b = !b
	}
	return b
}

func (inst *FilterRegistrationSorter) Swap(i1, i2 int) {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	inst.items[i1] = o2
	inst.items[i2] = o1
}

////////////////////////////////////////////////////////////////////////////////
