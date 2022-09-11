//go:generate go run github.com/pangdogs/galaxy/core/eventcode --decl_file=$GOFILE --not_import_core gen_emit --package=$GOPACKAGE
package core

// EventCompMgrAddComponents [EmitUnExport] 事件定义：实体的组件管理器加入一些组件
type EventCompMgrAddComponents interface {
	OnCompMgrAddComponents(entity Entity, components []Component)
}

// EventCompMgrRemoveComponent [EmitUnExport] 事件定义：实体的组件管理器删除组件
type EventCompMgrRemoveComponent interface {
	OnCompMgrRemoveComponent(entity Entity, component Component)
}
