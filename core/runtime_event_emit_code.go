// Code generated by eventcode --decl_file=runtime_event.go --not_import_core gen_emit --package=core; DO NOT EDIT.

package core

func emitEventUpdate(event IEvent) {
	if event == nil {
		panic("nil event")
	}
	event.Emit(func(delegate IFaceCache) bool {
		Cache2IFace[eventUpdate](delegate).Update()
		return true
	})
}

func emitEventLateUpdate(event IEvent) {
	if event == nil {
		panic("nil event")
	}
	event.Emit(func(delegate IFaceCache) bool {
		Cache2IFace[eventLateUpdate](delegate).LateUpdate()
		return true
	})
}
