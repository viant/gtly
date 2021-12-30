package gtly

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

//Object represents dynamic object
type Object struct {
	proto *Proto
	setAt []bool
	addr  unsafe.Pointer
	value reflect.Value
}

//Set sets a value
func (o *Object) Set(val interface{}) error {
	switch actual := val.(type) {
	case map[string]interface{}:
		for k, v := range actual {
			m := o.Mutator(k)
			m(v)
		}
		return nil
	case []interface{}:
		for k, v := range actual {
			m := o.MutatorAt(k)
			m(v)
		}
		return nil
	}

	return fmt.Errorf("unsupported type: %T", val)
}

//Proto returns object proto
func (o *Object) Proto() *Proto {
	return o.proto
}

//Addr returns struct unsafe pointer
func (o *Object) Addr() unsafe.Pointer {
	return o.addr
}

//Interface returns a struct value
func (o *Object) Interface() interface{} {
	return o.proto.xType.Interface(o.addr)
}

//SetValue sets fieldValues
func (o *Object) SetValue(fieldName string, value interface{}) {
	field := o.proto.Field(fieldName)
	o.set(field, value)
}

func (o *Object) set(field *Field, value interface{}) {
	switch field.kind {
	case reflect.Ptr, reflect.Slice, reflect.Struct, reflect.Func, reflect.Interface, reflect.Map:
		field.xField.SetValue(o.addr, value)
	case reflect.Float32:
		switch actual := value.(type) {
		case float32:
			field.xField.SetFloat32(o.addr, actual)
		case float64:
			field.xField.SetFloat32(o.addr, float32(actual))
		case int:
			field.xField.SetFloat32(o.addr, float32(actual))
		case int64:
			field.xField.SetFloat64(o.addr, float64(actual))
		default:
			field.xField.Set(o.addr, actual)
		}
	case reflect.Float64:
		switch actual := value.(type) {
		case float32:
			field.xField.SetFloat64(o.addr, float64(actual))
		case float64:
			field.xField.SetFloat64(o.addr, actual)
		case int:
			field.xField.SetFloat64(o.addr, float64(actual))
		case int64:
			field.xField.SetFloat64(o.addr, float64(actual))
		default:
			field.xField.Set(o.addr, actual)
		}
	case reflect.Int:
		switch actual := value.(type) {
		case float32:
			field.xField.SetInt(o.addr, int(actual))
		case float64:
			field.xField.SetInt(o.addr, int(actual))
		case int:
			field.xField.SetInt(o.addr, actual)
		case int64:
			field.xField.SetInt(o.addr, int(actual))
		default:
			field.xField.Set(o.addr, actual)
		}
	default:
		field.xField.Set(o.addr, value)
	}
	o.markFieldSet(field.Index)

}

func (o *Object) mutator(field *Field) func(value interface{}) {
	return func(value interface{}) {
		o.set(field, value)
	}
}

//Mutator returns field mutator
func (o *Object) Mutator(fieldName string) func(value interface{}) {
	field := o.proto.Field(fieldName)
	return o.mutator(field)
}

//MutatorAt returns field mutator
func (o *Object) MutatorAt(index int) func(value interface{}) {
	field := o.proto.FieldAt(index)
	return o.mutator(field)
}

//Accessor returns a field mutator
func (o *Object) Accessor(fieldName string) func() interface{} {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func() interface{} {
		return xField.Value(o.addr)
	}
}

//IntMutator returns an int mutator
func (o *Object) IntMutator(fieldName string) func(value int) {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func(value int) {
		xField.SetInt(o.addr, value)
		o.markFieldSet(field.Index)
	}
}

//IntAccessor returns an int accessor
func (o *Object) IntAccessor(fieldName string) func() int {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func() int {
		return xField.Int(o.addr)
	}
}

//Float64Mutator returns a float mutator
func (o *Object) Float64Mutator(fieldName string) func(value float64) {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func(value float64) {
		xField.SetFloat64(o.addr, value)
		o.markFieldSet(field.Index)
	}
}

//Float64Accessor returns a float accesor
func (o *Object) Float64Accessor(fieldName string) func() float64 {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func() float64 {
		return xField.Float64(o.addr)
	}
}

//StringMutator returns a string mutator
func (o *Object) StringMutator(fieldName string) func(value string) {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func(value string) {
		xField.SetString(o.addr, value)
		o.markFieldSet(field.Index)
	}
}

//StringAccessor returns string accessor
func (o *Object) StringAccessor(fieldName string) func() string {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func() string {
		return xField.String(o.addr)
	}
}

//BoolMutator returns bool mutator
func (o *Object) BoolMutator(fieldName string) func(value bool) {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func(value bool) {
		xField.SetBool(o.addr, value)
		o.markFieldSet(field.Index)
	}
}

//BoolAccessor returns bool accessor
func (o *Object) BoolAccessor(fieldName string) func() bool {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func() bool {
		return xField.Bool(o.addr)
	}
}

//TimeMutator return time mutator
func (o *Object) TimeMutator(fieldName string) func(time time.Time) {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func(value time.Time) {
		xField.SetTime(o.addr, value)
		o.markFieldSet(field.Index)
	}
}

//TimeAccessor returns time accessor
func (o *Object) TimeAccessor(fieldName string) func() time.Time {
	field := o.proto.Field(fieldName)
	xField := field.xField
	return func() time.Time {
		return xField.Time(o.addr)
	}
}

//Value returns field value
func (o *Object) Value(fieldName string) interface{} {
	field := o.proto.Field(fieldName)
	return field.xField.Value(o.addr)
}

//SetAt returns true if value was set at given index
func (o *Object) SetAt(index int) bool {
	if index >= len(o.setAt) {
		return false
	}
	return o.setAt[index]
}

func (o *Object) markFieldSet(index int) {
	o.setAt[index] = true
}

//IsNil returns true if object is nil
func (o *Object) IsNil() bool {
	for _, v := range o.setAt {
		if v == true {
			return false
		}
	}
	return true
}

//ValueAt get value for supplied filed Index
func (o *Object) ValueAt(index int) (interface{}, bool) {
	if index >= len(o.setAt) || !o.setAt[index] {
		return nil, false
	}
	field := o.proto.FieldAt(index)
	return field.xField.Value(o.addr), true
}

//IntAt returns int value for specified index
func (o *Object) IntAt(index int) (int, bool) {
	if index >= len(o.setAt) || !o.setAt[index] {
		return 0, false
	}
	field := o.proto.FieldAt(index)
	return field.xField.Int(o.addr), true
}

//StringAt returns int value for specified index
func (o *Object) StringAt(index int) (string, bool) {
	if index >= len(o.setAt) || !o.setAt[index] {
		return "", false
	}
	field := o.proto.FieldAt(index)
	return field.xField.String(o.addr), true
}

//BoolAt returns bool value for specified index
func (o *Object) BoolAt(index int) (bool, bool) {
	if index >= len(o.setAt) || !o.setAt[index] {
		return false, false
	}
	field := o.proto.FieldAt(index)
	return field.xField.Bool(o.addr), true
}

//FloatAt returns float value for specified index
func (o *Object) FloatAt(index int) (float64, bool) {
	if index >= len(o.setAt) || !o.setAt[index] {
		return 0, false
	}
	field := o.proto.FieldAt(index)
	return field.xField.Float64(o.addr), true
}

//TimeAt returns time value for specified index
func (o *Object) TimeAt(index int) (time.Time, bool) {
	if index >= len(o.setAt) || !o.setAt[index] {
		return time.Time{}, false
	}
	field := o.proto.FieldAt(index)
	return field.xField.Time(o.addr), true
}

//AsMap return map
func (o *Object) AsMap() map[string]interface{} {
	result := map[string]interface{}{}
	objAddr := o.addr
	for _, field := range o.proto.Fields() {
		if !o.SetAt(field.Index) || field.hidden {
			continue
		}
		outputName := o.FieldOutputName(field)
		if outputName == "" {
			continue
		}
		value := field.xField.Interface(objAddr)
		result[outputName] = value
	}
	return result
}

//FieldOutputName returns field output name
func (o *Object) FieldOutputName(field *Field) string {
	outputName := field.Name
	if field.outputName != "" {
		outputName = field.outputName
	}
	return outputName
}

//Field returns field by name
func (o *Object) Field(name string) *Field {
	return o.proto.Field(name)
}
