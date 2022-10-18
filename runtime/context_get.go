package runtime

import (
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/service"
	"github.com/pangdogs/galaxy/util"
)

// EntityContext 从实体上获取运行时上下文
func EntityContext(entity ec.Entity) Context {
	if entity == nil {
		panic("nil entity")
	}

	ctx := ec.UnsafeEntity(entity).GetContext()
	if ctx == util.NilIfaceCache {
		panic("nil context")
	}

	return util.Cache2Iface[Context](ctx)
}

// ComponentContext 从组件上获取运行时上下文
func ComponentContext(comp ec.Component) Context {
	if comp == nil {
		panic("nil comp")
	}

	return EntityContext(comp.GetEntity())
}

func entityServiceContext(entity ec.Entity) service.Context {
	return EntityContext(entity).GetServiceCtx()
}

func componentServiceContext(comp ec.Component) service.Context {
	return ComponentContext(comp).GetServiceCtx()
}
