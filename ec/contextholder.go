package ec

import (
	"github.com/pangdogs/galaxy/util"
)

// ContextHolder 上下文持有者，用于从实体或组件上获取上下文
type ContextHolder interface {
	getContext() util.IfaceCache
}
