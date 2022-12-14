package runtime

import (
	"github.com/golaxy-kit/golaxy/internal"
	"github.com/golaxy-kit/golaxy/localevent"
	"github.com/golaxy-kit/golaxy/service"
	"github.com/golaxy-kit/golaxy/util"
	"github.com/golaxy-kit/golaxy/util/container"
)

// NewContext 创建运行时上下文
func NewContext(serviceCtx service.Context, options ...ContextOption) Context {
	opts := ContextOptions{}
	WithContextOption.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return UnsafeNewContext(serviceCtx, opts)
}

func UnsafeNewContext(serviceCtx service.Context, options ContextOptions) Context {
	if !options.Inheritor.IsNil() {
		options.Inheritor.Iface.init(serviceCtx, &options)
		return options.Inheritor.Iface
	}

	ctx := &ContextBehavior{}
	ctx.init(serviceCtx, &options)

	return ctx.opts.Inheritor.Iface
}

// Context 运行时上下文接口
type Context interface {
	_InnerGC
	container.GCCollector
	internal.Context
	internal.RunningMark
	_SafeCall

	// GetServiceCtx 获取服务上下文
	GetServiceCtx() service.Context
	// GetFrame 获取帧
	GetFrame() Frame
	// GetEntityMgr 获取实体管理器
	GetEntityMgr() IEntityMgr
	// GetECTree 获取主EC树
	GetECTree() IECTree
	// GetFaceCache 获取Face缓存
	GetFaceCache() *container.Cache[util.FaceAny]
	// GetHookCache 获取Hook缓存
	GetHookCache() *container.Cache[localevent.Hook]

	init(serviceCtx service.Context, opts *ContextOptions)
	getOptions() *ContextOptions
	setFrame(frame Frame)
}

// ContextBehavior 运行时上下文行为，在需要扩展运行时上下文能力时，匿名嵌入至运行时上下文结构体中
type ContextBehavior struct {
	internal.ContextBehavior
	internal.RunningMarkBehavior
	opts       ContextOptions
	serviceCtx service.Context
	frame      Frame
	entityMgr  _EntityMgr
	ecTree     ECTree
	callee     internal.Callee
	gcList     []container.GC
	innerGC    _ContextInnerGC
}

// GetServiceCtx 获取服务上下文
func (ctx *ContextBehavior) GetServiceCtx() service.Context {
	return ctx.serviceCtx
}

// GetFrame 获取帧
func (ctx *ContextBehavior) GetFrame() Frame {
	return ctx.frame
}

// GetEntityMgr 获取实体管理器
func (ctx *ContextBehavior) GetEntityMgr() IEntityMgr {
	return &ctx.entityMgr
}

// GetECTree 获取主EC树
func (ctx *ContextBehavior) GetECTree() IECTree {
	return &ctx.ecTree
}

// GetFaceCache 获取Face缓存
func (ctx *ContextBehavior) GetFaceCache() *container.Cache[util.FaceAny] {
	return ctx.opts.FaceCache
}

// GetHookCache 获取Hook缓存
func (ctx *ContextBehavior) GetHookCache() *container.Cache[localevent.Hook] {
	return ctx.opts.HookCache
}

// CollectGC 收集GC
func (ctx *ContextBehavior) CollectGC(gc container.GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	ctx.gcList = append(ctx.gcList, gc)
}

func (ctx *ContextBehavior) init(serviceCtx service.Context, opts *ContextOptions) {
	if serviceCtx == nil {
		panic("nil serviceCtx")
	}

	if opts == nil {
		panic("nil opts")
	}

	ctx.opts = *opts

	if ctx.opts.Inheritor.IsNil() {
		ctx.opts.Inheritor = util.NewFace[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = serviceCtx
	}

	ctx.innerGC.Init(ctx)

	internal.UnsafeContext(&ctx.ContextBehavior).Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.serviceCtx = serviceCtx
	ctx.entityMgr.Init(ctx.getOptions().Inheritor.Iface)
	ctx.ecTree.init(ctx.opts.Inheritor.Iface, true)
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}

func (ctx *ContextBehavior) setFrame(frame Frame) {
	ctx.frame = frame
}

func (ctx *ContextBehavior) getInnerGC() container.GC {
	return &ctx.innerGC
}
