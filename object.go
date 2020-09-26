package gtly

import (
	"github.com/viant/toolbox"
	"time"
)

//Object represents dynamic object
type Object struct {
	_proto *Proto
	_data  []interface{}
}

//Proto returns object _proto
func (o *Object) Proto() *Proto {
	return o._proto
}

//Init initialise entire object
func (o *Object) Init(values map[string]interface{}) {
	o._data = o._proto.asValues(values)
}

//AsMap return map
func (o *Object) AsMap() map[string]interface{} {
	return o._proto.asMap(o._data)
}

//SetInt sets int values
func (o *Object) SetInt(name string, value int) {
	field := o._proto.FieldWithValue(name, value)
	field.SetValue(value, &o._data)
}

//SetFloat sets float values
func (o *Object) SetFloat(name string, value float64) {
	field := o._proto.FieldWithValue(name, value)
	field.SetValue(value, &o._data)
}

//SetString sets string value
func (o *Object) SetString(name string, value string) {
	field := o._proto.FieldWithValue(name, value)
	field.SetValue(value, &o._data)
}

//SetBool sets string value
func (o *Object) SetBool(name string, value bool) {
	field := o._proto.FieldWithValue(name, value)
	field.SetValue(value, &o._data)
}

//SetTime sets string value
func (o *Object) SetTime(name string, value time.Time) {
	field := o._proto.FieldWithValue(name, value)
	field.SetValue(value, &o._data)
}

//SetValue sets values
func (o *Object) SetValue(name string, value interface{}) {
	field := o._proto.FieldWithValue(name, value)
	field.Set(value, &o._data)
}

//Value get value for supplied Name
func (o *Object) Value(name string) interface{} {
	field := o._proto.Field(name)
	if field == nil {
		return nil
	}
	return field.Get(o._data)
}

//Int returns int for supplied field name
func (o *Object) Int(name string) int {
	field := o._proto.Field(name)
	if field == nil {
		return 0
	}
	value := field.Get(o._data)
	if value == nil {
		return 0
	}
	return toolbox.AsInt(value)
}

//Float returns float for supplied field name
func (o *Object) Float(name string) float64 {
	field := o._proto.Field(name)
	if field == nil {
		return 0
	}
	value := field.Get(o._data)
	if value == nil {
		return 0
	}
	return toolbox.AsFloat(value)
}

//Bool return bool for supplied field name
func (o *Object) Bool(name string) bool {
	field := o._proto.Field(name)
	if field == nil {
		return false
	}
	value := field.Get(o._data)
	if value == nil {
		return false
	}
	return toolbox.AsBoolean(value)
}

//String return string for supplied field name
func (o *Object) String(name string) string {
	field := o._proto.Field(name)
	if field == nil {
		return ""
	}
	value := field.Get(o._data)
	if value == nil {
		return ""
	}
	return toolbox.AsString(value)
}

//ValueAt get value for supplied filed Index
func (o *Object) ValueAt(index int) interface{} {
	if index >= len(o._data) {
		return nil
	}
	return Value(o._data[index])
}

//IntAt returns int value for specified index
func (o *Object) IntAt(index int) int {
	if index >= len(o._data) {
		return 0
	}
	val, ok := o._data[index].(int)
	if !ok {
		return 0
	}
	return val
}

//FloatAt returns float value for specified index
func (o *Object) FloatAt(index int) float64 {
	if index >= len(o._data) {
		return 0
	}
	val, ok := o._data[index].(float64)
	if !ok {
		return 0
	}
	return val
}

//StringAt returns int value for specified index
func (o *Object) StringAt(index int) string {
	if index >= len(o._data) {
		return ""
	}
	val, ok := o._data[index].(string)
	if !ok {
		return ""
	}
	return val
}

//BoolAt returns bool value for specified index
func (o *Object) BoolAt(index int) bool {
	if index >= len(o._data) {
		return false
	}
	val, ok := o._data[index].(bool)
	if !ok {
		return false
	}
	return val
}

//HasAt returns true if has value
func (o *Object) HasAt(index int) bool {
	if index >= len(o._data) {
		return false
	}
	return o._data[index] != nil
}

//FloatValue return float for supplied Name
func (o *Object) FloatValue(name string) (*float64, error) {
	val := o.Value(name)
	if val == nil {
		return nil, nil
	}
	casted, err := toolbox.ToFloat(val)
	return &casted, err
}

//IntValue returns int value
func (o *Object) IntValue(name string) (*int, error) {
	val := o.Value(name)
	if val == nil {
		return nil, nil
	}
	casted, err := toolbox.ToInt(val)
	return &casted, err
}

//StringValue returns int value
func (o *Object) StringValue(name string) *string {
	val := o.Value(name)
	if val == nil {
		return nil
	}
	casted := toolbox.AsString(val)
	return &casted
}

//IsNil returns true if object is nil
func (o *Object) IsNil() bool {
	for i := range o._data {
		if !o._proto.fields[i].IsEmpty(o._proto, o._data[i]) {
			return false
		}
	}
	return true
}
