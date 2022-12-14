package ec

import (
	"github.com/golaxy-kit/golaxy/util"
)

// ContextResolver 上下文供应者，用于从实体或组件上获取上下文
type ContextResolver interface {
	getContext() util.IfaceCache
}
