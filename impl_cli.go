package cli

type implCLI struct {
	context *Context
}

func (inst *implCLI) _Impl() CLI {
	return inst
}

func (inst *implCLI) GetClient() Client {
	return inst.context.Client
}

func (inst *implCLI) GetServer() Server {
	return inst.context.Server
}
