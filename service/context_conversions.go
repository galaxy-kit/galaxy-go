package service

import "github.com/galaxy-kit/galaxy-go/util"

// GetInheritor 获取服务上下文的继承者
func GetInheritor[T any](ctx Context) T {
	return util.Cache2Iface[T](ctx.getOptions().Inheritor.Cache)
}
