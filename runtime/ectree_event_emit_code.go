// Code generated by eventcode --decl_file=ectree_event.go gen_emit --package=runtime; DO NOT EDIT.

package runtime

import (
	localevent "github.com/galaxy-kit/galaxy-go/localevent"
	"github.com/galaxy-kit/galaxy-go/ec"
	"github.com/galaxy-kit/galaxy-go/util"
)

func emitEventECTreeAddChild(event localevent.IEvent, ecTree IECTree, parent, child ec.Entity) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[EventECTreeAddChild](delegate).OnAddChild(ecTree, parent, child)
		return true
	})
}

func emitEventECTreeRemoveChild(event localevent.IEvent, ecTree IECTree, parent, child ec.Entity) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[EventECTreeRemoveChild](delegate).OnRemoveChild(ecTree, parent, child)
		return true
	})
}
