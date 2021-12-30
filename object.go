package gtly

import (
	"fmt"
	"reflect"
	"unsafe"
)

//Object represents dynamic object
type Object struct {
	proto *Proto
	setAt []bool
	addr  unsafe.Pointer
	value reflect.Value
}

//Set sets a value from a map of a slice (slice index has to match field index)
func (o *Object) Set(val interface{}) error {
	switch actual := val.(type) {
	case map[string]interface{}:
		for k, v := range actual {
			o.proto.Mutator(k).SetValue(o, v)
		}
		return nil
	case []interface{}:
		for k, v := range actual {
			o.proto.MutatorAt(k).SetValue(o, v)
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
	o.proto.mutators[field.Index].SetValue(o, value)
}

//SetValueAt sets field's value
func (o *Object) SetValueAt(fieldIndex int, value interface{}) {
	o.proto.mutators[fieldIndex].SetValue(o, value)
}

//Value get value for supplied filed name
func (o *Object) Value(fieldName string) interface{} {
	field := o.proto.Field(fieldName)
	if field.Index >= len(o.setAt) || !o.setAt[field.Index] {
		return nil
	}
	return o.proto.accessors[field.Index].Value(o)
}

//ValueAt get value for supplied filed Index
func (o *Object) ValueAt(fieldIndex int) (interface{}, bool) {
	if fieldIndex >= len(o.setAt) || !o.setAt[fieldIndex] {
		return nil, false
	}
	return o.proto.accessors[fieldIndex].Value(o), true
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

//AsMap return map
func (o *Object) AsMap() map[string]interface{} {
	result := map[string]interface{}{}
	fields := o.proto.Fields()
	for i := range o.proto.accessors {
		field := &fields[i]
		if !o.SetAt(field.Index) || field.hidden {
			continue
		}
		outputName := o.FieldOutputName(field)
		if outputName == "" {
			continue
		}
		value := o.proto.accessors[i].Value(o)
		result[outputName] = value
	}
	return result
}

//FieldOutputName returns Field output name
func (o *Object) FieldOutputName(field *Field) string {
	outputName := field.Name
	if field.outputName != "" {
		outputName = field.outputName
	}
	return outputName
}

//Field returns Field by name
func (o *Object) Field(name string) *Field {
	return o.proto.Field(name)
}
