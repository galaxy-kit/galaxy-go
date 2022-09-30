package pt

import (
	"github.com/pangdogs/galaxy/ec"
	"reflect"
)

type _CompConstructType int32

const (
	_CompConstructType_Reflect _CompConstructType = iota
	_CompConstructType_Creator
)

// ComponentPt 组件原型
type ComponentPt struct {
	Interface     string // 组件接口名称
	Tag           string // 组件标签
	Description   string // 组件描述信息
	constructType _CompConstructType
	tfComp        reflect.Type
	creator       func() ec.Component
}

// New 创建组件
func (pt *ComponentPt) New() ec.Component {
	switch pt.constructType {
	case _CompConstructType_Reflect:
		vfComp := reflect.New(pt.tfComp)

		comp := vfComp.Interface().(ec.Component)
		ec.UnsafeComponent(comp).SetReflectValue(vfComp)

		return comp

	case _CompConstructType_Creator:
		return pt.creator()
		
	default:
		panic("not support construct type")
	}
}
