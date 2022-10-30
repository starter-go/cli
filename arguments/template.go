package arguments

////////////////////////////////////////////////////////////////////////////////

// Value ... 表示一个简单的参数项
type Value interface {
	Exists() bool
	String() string
	Boolean() bool
	Int() int
}

// Option 表示一个带名称的选项
type Option interface {
	Name() string
	Value(index int) Value
	Exists() bool
}

// Arguments 表示一个命令模板，用来分析命令参数
type Arguments interface {

	// 原始的未解析的参数
	Raw() []string

	// 命名选项
	Options() []Option

	// 未命名项
	Items() []Value

	GetItem(index int) Value

	GetOption(name string) Option
}

// Template 表示一个命令模板，用来分析命令参数
type Template interface {
	Parse(args []string) Arguments
}

// TemplateBuilder 用来创建模板
type TemplateBuilder interface {
	AcceptOption(name string, countValue int) TemplateBuilder
	Create() Template
}

////////////////////////////////////////////////////////////////////////////////

// NewTemplateBuilder 新建一个 TemplateBuilder
func NewTemplateBuilder() TemplateBuilder {
	return &templateBuilderImpl{}
}
