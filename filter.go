package cli

// Filter ... 用来过滤请求的任务
type Filter interface {
	Pass(task *Task, chain FilterChain) error
}

// FilterChain ... 表示一连串 Filter 构成的链条
type FilterChain interface {
	Pass(task *Task) error
}

////////////////////////////////////////////////////////////////////////////////

// FilterChainBuilder 用来创建 FilterChain
type FilterChainBuilder struct {
	filters []*FilterRegistration
}

// AddFilter ...
func (inst *FilterChainBuilder) AddFilter(f Filter, order int) {
	if f == nil {
		return
	}
	fr := &FilterRegistration{Filter: f, Order: order}
	inst.filters = append(inst.filters, fr)
}

// AddFilterRegistry ...
func (inst *FilterChainBuilder) AddFilterRegistry(fr FilterRegistry) {
	if fr == nil {
		return
	}
	src := fr.GetFilters()
	for _, item := range src {
		if item == nil {
			continue
		}
		inst.filters = append(inst.filters, item)
	}
}

// AddFilterRegistration ...
func (inst *FilterChainBuilder) AddFilterRegistration(fr *FilterRegistration) {
	if fr == nil {
		return
	}
	inst.filters = append(inst.filters, fr)
}

// Create ...
func (inst *FilterChainBuilder) Create() FilterChain {
	src := inst.filters
	(&FilterRegistrationSorter{}).Sort(src, true)
	var p FilterChain = &filterChainEnding{}
	for _, fr := range src {
		f := fr.Filter
		node := &filterChainNode{filter: f, next: p}
		p = node
	}
	return p
}

////////////////////////////////////////////////////////////////////////////////

type filterChainNode struct {
	filter Filter
	next   FilterChain
}

func (inst *filterChainNode) Pass(t *Task) error {
	return inst.filter.Pass(t, inst.next)
}

////////////////////////////////////////////////////////////////////////////////

type filterChainEnding struct {
}

func (inst *filterChainEnding) Pass(t *Task) error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
