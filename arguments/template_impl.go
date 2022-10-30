package arguments

import (
	"errors"
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////

type argsImpl struct {
	raw     []string
	options map[string]Option
	items   []Value
}

func (inst *argsImpl) _Impl() Arguments {
	return inst
}

func (inst *argsImpl) Raw() []string {
	return inst.raw
}

// 命名选项
func (inst *argsImpl) Options() []Option {
	src := inst.options
	dst := make([]Option, 0)
	for _, item := range src {
		dst = append(dst, item)
	}
	return dst
}

// 未命名项
func (inst *argsImpl) Items() []Value {
	return inst.items
}

func (inst *argsImpl) GetItem(index int) Value {
	list := inst.items
	if list == nil {
		return &valueImpl{}
	}
	size := len(list)
	if 0 <= index && index < size {
		item := list[index]
		if item != nil {
			return item
		}
	}
	return &valueImpl{}
}

func (inst *argsImpl) GetOption(name string) Option {
	table := inst.options
	if table == nil {
		return &optionImpl{}
	}
	opt := table[name]
	if opt == nil {
		opt = &optionImpl{}
	}
	return opt
}

////////////////////////////////////////////////////////////////////////////////

type argsBuilder struct {
	acceptOptions map[string]*optionDescriptor
	args          argsImpl
	currentOption *optionImpl
}

func (inst *argsBuilder) init(t *templateImpl) {
	inst.acceptOptions = t.acceptOptions
	inst.args.options = make(map[string]Option)
}

func (inst *argsBuilder) append(item string) {

	// append to raw
	inst.args.raw = append(inst.args.raw, item)
	nextOptDesc := inst.acceptOptions[item]
	currentOpt := inst.currentOption

	if nextOptDesc != nil {
		// new option
		inst.addNewOption(nextOptDesc)
		return
	} else if currentOpt != nil {
		// append value to option
		err := currentOpt.appendValue(item)
		if err == nil {
			return
		}
		inst.currentOption = nil
	}
	// add free value
	inst.addFreeItem(item)
}

func (inst *argsBuilder) addNewOption(desc *optionDescriptor) {
	opt := &optionImpl{
		name:            desc.name,
		countValueLimit: desc.countValueLimit,
		values:          nil,
		exists:          true,
	}
	inst.currentOption = opt
	inst.args.options[opt.name] = opt
}

func (inst *argsBuilder) addFreeItem(item string) {
	v := &valueImpl{
		exists: true,
		value:  item,
	}
	inst.args.items = append(inst.args.items, v)
}

func (inst *argsBuilder) create() Arguments {
	args := &argsImpl{}
	*args = inst.args
	return args
}

////////////////////////////////////////////////////////////////////////////////

type templateImpl struct {
	acceptOptions map[string]*optionDescriptor
}

func (inst *templateImpl) _Impl() Template {
	return inst
}

func (inst *templateImpl) Parse(args []string) Arguments {
	ab := argsBuilder{}
	ab.init(inst)
	for _, a := range args {
		ab.append(a)
	}
	return ab.create()
}

////////////////////////////////////////////////////////////////////////////////

type optionDescriptor struct {
	name            string // 选项的名称
	countValueLimit int    // 接受值的最大数量
}

////////////////////////////////////////////////////////////////////////////////

type templateBuilderImpl struct {
	acceptOptions map[string]*optionDescriptor
}

func (inst *templateBuilderImpl) _Impl() TemplateBuilder {
	return inst
}

func (inst *templateBuilderImpl) getOptionTable() map[string]*optionDescriptor {
	t := inst.acceptOptions
	if t == nil {
		t = make(map[string]*optionDescriptor)
		inst.acceptOptions = t
	}
	return t
}

func (inst *templateBuilderImpl) AcceptOption(name string, countValue int) TemplateBuilder {
	opt := &optionDescriptor{
		name:            name,
		countValueLimit: countValue,
	}
	table := inst.getOptionTable()
	table[name] = opt
	return inst
}

func (inst *templateBuilderImpl) Create() Template {
	t := &templateImpl{}
	t.acceptOptions = inst.getOptionTable()
	return t
}

////////////////////////////////////////////////////////////////////////////////

type optionImpl struct {
	name            string
	values          []Value
	exists          bool
	countValueLimit int
}

func (inst *optionImpl) _Impl() Option {
	return inst
}

func (inst *optionImpl) Name() string {
	return inst.name
}

func (inst *optionImpl) Exists() bool {
	return inst.exists
}

func (inst *optionImpl) Value(index int) Value {
	list := inst.values
	if list != nil {
		size := len(list)
		if 0 <= index && index < size {
			return list[index]
		}
	}
	return emptyValue()
}

func (inst *optionImpl) appendValue(value string) error {
	size := len(inst.values)
	if size >= inst.countValueLimit {
		return errors.New("the count of value is over limit")
	}
	v := &valueImpl{
		exists: true,
		value:  value,
	}
	inst.values = append(inst.values, v)
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type valueImpl struct {
	exists bool
	value  string
}

func emptyValue() Value {
	return &valueImpl{}
}

func (inst *valueImpl) _Impl() Value {
	return inst
}

func (inst *valueImpl) Exists() bool {
	return inst.exists
}

func (inst *valueImpl) String() string {
	return inst.value
}

func (inst *valueImpl) Boolean() bool {
	str := inst.value
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return b
}

func (inst *valueImpl) Int() int {
	str := inst.value
	n, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0
	}
	return int(n)
}

////////////////////////////////////////////////////////////////////////////////
