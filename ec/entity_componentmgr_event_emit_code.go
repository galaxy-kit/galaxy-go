// Code generated by eventcode --decl_file=entity_componentmgr_event.go gen_emit --package=ec; DO NOT EDIT.

package ec

import (
	localevent "github.com/golaxy-kit/golaxy/localevent"
	"github.com/golaxy-kit/golaxy/util"
)

func emitEventCompMgrAddComponents(event localevent.IEvent, entity Entity, components []Component) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[EventCompMgrAddComponents](delegate).OnCompMgrAddComponents(entity, components)
		return true
	})
}

func emitEventCompMgrRemoveComponent(event localevent.IEvent, entity Entity, component Component) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[EventCompMgrRemoveComponent](delegate).OnCompMgrRemoveComponent(entity, component)
		return true
	})
}

func emitEventCompMgrFirstAccessComponent(event localevent.IEvent, entity Entity, component Component) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[EventCompMgrFirstAccessComponent](delegate).OnCompMgrFirstAccessComponent(entity, component)
		return true
	})
}
