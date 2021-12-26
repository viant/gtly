package gtly

import (
	"reflect"
	"time"
	"unsafe"
)

//Object represents dynamic object
type Object struct {
	_proto    *Proto
	_setAt    []bool
	_dataAddr unsafe.Pointer
	value     reflect.Value
}

//Proto returns object _proto
func (o *Object) Proto() *Proto {
	return o._proto
}

//StructValue returns a struct value
func (o *Object) StructValue() interface{} {
	return o.value.Elem().Interface()
}

//SetValue sets fieldValues
func (o *Object) SetValue(fieldName string, value interface{}) {
	field := o._proto.Field(fieldName)
	switch valueTyped := value.(type) {
	case float64:
		if field.kind == reflect.Int {
			field.xField.SetInt(o._dataAddr, int(valueTyped))
		} else {
			field.xField.SetFloat64(o._dataAddr, valueTyped)
		}
	case float32:
		field.xField.SetFloat32(o._dataAddr, valueTyped)
	case int:
		field.xField.SetInt(o._dataAddr, valueTyped)
	case int64:
		field.xField.SetInt64(o._dataAddr, valueTyped)
	case int32:
		field.xField.SetInt32(o._dataAddr, valueTyped)
	case int16:
		field.xField.SetInt16(o._dataAddr, valueTyped)
	case int8:
		field.xField.SetInt8(o._dataAddr, valueTyped)
	case uint:
		field.xField.SetUint(o._dataAddr, valueTyped)
	case uint64:
		field.xField.SetUint64(o._dataAddr, valueTyped)
	case uint32:
		field.xField.SetUint32(o._dataAddr, valueTyped)
	case uint16:
		field.xField.SetUint16(o._dataAddr, valueTyped)
	case uint8:
		field.xField.SetUint8(o._dataAddr, valueTyped)
	case string:
		field.xField.SetString(o._dataAddr, valueTyped)
	case time.Time:
		field.xField.SetTime(o._dataAddr, valueTyped)
	case bool:
		field.xField.SetBool(o._dataAddr, valueTyped)
	default:
		field.xField.SetValue(o._dataAddr, value)
	}
	o.markFieldSet(field.Index)
}

//Mutator returns field mutator
func (o *Object) Mutator(fieldName string) func(value interface{}) {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func(value interface{}) {
		XField.SetValue(o._dataAddr, value)
		o.markFieldSet(field.Index)
	}
}

//Accessor returns a field mutator
func (o *Object) Accessor(fieldName string) func() interface{} {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func() interface{} {
		return XField.Value(o._dataAddr)
	}
}

//IntMutator returns an int mutator
func (o *Object) IntMutator(fieldName string) func(value int) {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func(value int) {
		XField.SetInt(o._dataAddr, value)
		o.markFieldSet(field.Index)
	}
}

//IntAccessor returns an int accessor
func (o *Object) IntAccessor(fieldName string) func() int {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func() int {
		return XField.Int(o._dataAddr)
	}
}

//FloatMutator returns a float mutator
func (o *Object) FloatMutator(fieldName string) func(value float64) {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func(value float64) {
		XField.SetFloat64(o._dataAddr, value)
		o.markFieldSet(field.Index)
	}
}

//FloatAccessor returns a float accesor
func (o *Object) FloatAccessor(fieldName string) func() float64 {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func() float64 {
		return XField.Float64(o._dataAddr)
	}
}

//StringMutator returns a string mutator
func (o *Object) StringMutator(fieldName string) func(value string) {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func(value string) {
		XField.SetString(o._dataAddr, value)
		o.markFieldSet(field.Index)
	}
}

//StringAccessor returns string accessor
func (o *Object) StringAccessor(fieldName string) func() string {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func() string {
		return XField.String(o._dataAddr)
	}
}

//
func (o *Object) BoolMutator(fieldName string) func(value bool) {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func(value bool) {
		XField.SetBool(o._dataAddr, value)
		o.markFieldSet(field.Index)
	}
}

func (o *Object) BoolAccessor(fieldName string) func() bool {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func() bool {
		return XField.Bool(o._dataAddr)
	}
}

func (o *Object) TimeMutator(fieldName string) func(time time.Time) {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func(value time.Time) {
		XField.SetTime(o._dataAddr, value)
		o.markFieldSet(field.Index)
	}
}

func (o *Object) TimeAccessor(fieldName string) func() time.Time {
	field := o._proto.Field(fieldName)
	XField := field.xField
	return func() time.Time {
		return XField.Time(o._dataAddr)
	}
}

func (o *Object) Value(fieldName string) interface{} {
	field := o._proto.Field(fieldName)
	return field.xField.Value(o._dataAddr)
}

//Init initialise entire object
func (o *Object) Init(values map[string]interface{}) {
	for k, v := range values {
		o.SetValue(k, v)
	}
}

//SetAt returns true if value was set at given index
func (o *Object) SetAt(index int) bool {
	return index < len(o._setAt) && o._setAt[index]
}

func (o *Object) markFieldSet(index int) {
	o._setAt[index] = true
}

//IsNil returns true if object is nil
func (o *Object) IsNil() bool {
	for _, v := range o._setAt {
		if v == true {
			return false
		}
	}
	return true
}

//ValueAt get value for supplied filed Index
func (o *Object) ValueAt(index int) (interface{}, bool) {
	if index < 0 || index >= len(o._proto.fields) || !o._setAt[index] {
		return nil, false
	}

	field := o._proto.FieldAt(index)
	return field.xField.Value(o._dataAddr), true
}

//IntAt returns int value for specified index
func (o *Object) IntAt(index int) (int, bool) {
	if index < 0 || index >= len(o._proto.fields) || !o._setAt[index] {
		return 0, false
	}

	field := o._proto.FieldAt(index)
	return field.xField.Int(o._dataAddr), true
}

//StringAt returns int value for specified index
func (o *Object) StringAt(index int) (string, bool) {
	if index < 0 || index >= len(o._proto.fields) || !o._setAt[index] {
		return "", false
	}

	field := o._proto.FieldAt(index)
	return field.xField.String(o._dataAddr), true
}

//BoolAt returns bool value for specified index
func (o *Object) BoolAt(index int) (bool, bool) {
	if index < 0 || index >= len(o._proto.fields) || !o._setAt[index] {
		return false, false
	}

	field := o._proto.FieldAt(index)
	return field.xField.Bool(o._dataAddr), true
}

//FloatAt returns float value for specified index
func (o *Object) FloatAt(index int) (float64, bool) {
	if index < 0 || index >= len(o._proto.fields) || !o._setAt[index] {
		return 0.0, false
	}

	field := o._proto.FieldAt(index)
	return field.xField.Float64(o._dataAddr), true
}

//TimeAt returns time value for specified index
func (o *Object) TimeAt(index int) (time.Time, bool) {
	if index < 0 || index >= len(o._proto.fields) || !o._setAt[index] {
		return time.Time{}, false
	}

	field := o._proto.FieldAt(index)
	return field.xField.Time(o._dataAddr), true
}

//AsMap return map
func (o *Object) AsMap() map[string]interface{} {
	result := map[string]interface{}{}
	objAddr := o._dataAddr
	for _, field := range o._proto.Fields() {
		if !o.SetAt(field.Index) || field.hidden {
			continue
		}
		outputName := o.FieldOutputName(field)
		if outputName == "" {
			continue
		}
		value := field.xField.Value(objAddr)
		result[outputName] = value
	}
	return result
}

func (o *Object) FieldOutputName(field *Field) string {
	outputName := field.Name
	if field.outputName != "" {
		outputName = field.outputName
	}
	return outputName
}

func (o *Object) Field(name string) *Field {
	return o._proto.Field(name)
}
