// Code generated by eventcode --decl_file=runtime_event.go gen_emit --package=galaxy --default_export=0; DO NOT EDIT.

package galaxy

import (
	localevent "github.com/golaxy-kit/golaxy/localevent"
	"github.com/golaxy-kit/golaxy/util"
)

func emitEventUpdate(event localevent.IEvent) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[eventUpdate](delegate).Update()
		return true
	})
}

func emitEventLateUpdate(event localevent.IEvent) {
	if event == nil {
		panic("nil event")
	}
	localevent.UnsafeEvent(event).Emit(func(delegate util.IfaceCache) bool {
		util.Cache2Iface[eventLateUpdate](delegate).LateUpdate()
		return true
	})
}
