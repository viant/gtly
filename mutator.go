package gtly

import (
	"github.com/viant/xunsafe"
	"reflect"
	"time"
)

//Mutator represents object mutator
type Mutator struct {
	index int
	*xunsafe.Field
}

func (m *Mutator) init(index int, field *xunsafe.Field) {
	m.index = index
	m.Field = field
}

//SetValue sets value
func (m *Mutator) SetValue(object *Object, value interface{}) {
	switch m.Field.Type.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Struct, reflect.Func, reflect.Interface, reflect.Map:
		m.Field.SetValue(object.addr, value)
	case reflect.Float32:
		switch actual := value.(type) {
		case float32:
			m.Field.SetFloat32(object.addr, actual)
		case float64:
			m.Field.SetFloat32(object.addr, float32(actual))
		case int:
			m.Field.SetFloat32(object.addr, float32(actual))
		case int64:
			m.Field.SetFloat64(object.addr, float64(actual))
		default:
			m.Field.Set(object.addr, actual)
		}
	case reflect.Float64:
		switch actual := value.(type) {
		case float32:
			m.Field.SetFloat64(object.addr, float64(actual))
		case float64:
			m.Field.SetFloat64(object.addr, actual)
		case int:
			m.Field.SetFloat64(object.addr, float64(actual))
		case int64:
			m.Field.SetFloat64(object.addr, float64(actual))
		default:
			m.Field.Set(object.addr, actual)
		}
	case reflect.Int:
		switch actual := value.(type) {
		case float32:
			m.Field.SetInt(object.addr, int(actual))
		case float64:
			m.Field.SetInt(object.addr, int(actual))
		case int:
			m.Field.SetInt(object.addr, actual)
		case int64:
			m.Field.SetInt(object.addr, int(actual))
		default:
			m.Field.Set(object.addr, actual)
		}
	default:
		m.Field.Set(object.addr, value)
	}
	object.markFieldSet(m.index)
}

//Int sets int value
func (m *Mutator) Int(object *Object, value int) {
	m.Field.SetInt(object.addr, value)
	object.markFieldSet(m.index)
}

//Int64 sets int64 value
func (m *Mutator) Int64(object *Object, value int64) {
	m.Field.SetInt64(object.addr, value)
	object.markFieldSet(m.index)
}

//Float32 sets float32 value
func (m *Mutator) Float32(object *Object, value float32) {
	m.Field.SetFloat32(object.addr, value)
	object.markFieldSet(m.index)
}

//Float64 sets float64 value
func (m *Mutator) Float64(object *Object, value float64) {
	m.Field.SetFloat64(object.addr, value)
	object.markFieldSet(m.index)
}

//Bool sets bool value
func (m *Mutator) Bool(object *Object, value bool) {
	m.Field.SetBool(object.addr, value)
	object.markFieldSet(m.index)
}

//String sets string value
func (m *Mutator) String(object *Object, value string) {
	m.Field.SetString(object.addr, value)
	object.markFieldSet(m.index)
}

//StringPtr sets *string value
func (m *Mutator) StringPtr(object *Object, value *string) {
	m.Field.SetStringPtr(object.addr, value)
	object.markFieldSet(m.index)
}

//Time sets time value
func (m *Mutator) Time(object *Object, value time.Time) {
	m.Field.SetTime(object.addr, value)
	object.markFieldSet(m.index)
}

//TimePtr sets time value
func (m *Mutator) TimePtr(object *Object, value *time.Time) {
	m.Field.SetTimePtr(object.addr, value)
	object.markFieldSet(m.index)
}

//Bytes sets []byte value
func (m *Mutator) Bytes(object *Object, value []byte) {
	m.Field.SetBytes(object.addr, value)
	object.markFieldSet(m.index)
}
