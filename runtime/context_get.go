package runtime

import (
	"github.com/golaxy-kit/golaxy/ec"
	"github.com/golaxy-kit/golaxy/service"
	"github.com/golaxy-kit/golaxy/util"
)

// Get 获取运行时上下文
func Get(ctxHolder ec.ContextResolver) Context {
	if ctxHolder == nil {
		panic("nil ctxHolder")
	}

	ctx := ec.UnsafeContextResolver(ctxHolder).GetContext()
	if ctx == util.NilIfaceCache {
		panic("nil context")
	}

	return util.Cache2Iface[Context](ctx)
}

// TryGet 尝试获取运行时上下文
func TryGet(ctxHolder ec.ContextResolver) (Context, bool) {
	if ctxHolder == nil {
		return nil, false
	}

	ctx := ec.UnsafeContextResolver(ctxHolder).GetContext()
	if ctx == util.NilIfaceCache {
		return nil, false
	}

	return util.Cache2Iface[Context](ctx), true
}

func getServiceContext(ctxHolder ec.ContextResolver) service.Context {
	return Get(ctxHolder).GetServiceCtx()
}

func tryGetServiceContext(ctxHolder ec.ContextResolver) (service.Context, bool) {
	runtimeCtx, ok := TryGet(ctxHolder)
	if !ok {
		return nil, false
	}
	return runtimeCtx.GetServiceCtx(), true
}
