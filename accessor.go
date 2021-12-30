package gtly

import (
	"github.com/viant/xunsafe"
	"time"
)

//Accessor represents object mutator
type Accessor struct {
	index int
	*xunsafe.Field
}

func (m *Accessor) init(index int, field *xunsafe.Field) {
	m.index = index
	m.Field = field
}

//Value returns value
func (m *Accessor) Value(object *Object) interface{} {
	return m.Field.Value(object.addr)
}

//Int returns int value
func (m *Accessor) Int(object *Object) int {
	return m.Field.Int(object.addr)
}

//Int64 returns int64 value
func (m *Accessor) Int64(object *Object) int64 {
	return m.Field.Int64(object.addr)

}

//Float32 returns float32 value
func (m *Accessor) Float32(object *Object) float32 {
	return m.Field.Float32(object.addr)

}

//Float64 returns float64 value
func (m *Accessor) Float64(object *Object) float64 {
	return m.Field.Float64(object.addr)

}

//Bool returns bool value
func (m *Accessor) Bool(object *Object) bool {
	return m.Field.Bool(object.addr)

}

//String returns string value
func (m *Accessor) String(object *Object) string {
	return m.Field.String(object.addr)
}

//StringPtr returns *string value
func (m *Accessor) StringPtr(object *Object) *string {
	return m.Field.StringPtr(object.addr)
}

//Time returns time value
func (m *Accessor) Time(object *Object) time.Time {
	return m.Field.Time(object.addr)
}

//TimePtr returns *time.Time value
func (m *Accessor) TimePtr(object *Object) *time.Time {
	return m.Field.TimePtr(object.addr)
}

//Bytes returns []byte value
func (m *Accessor) Bytes(object *Object) []byte {
	return m.Field.Bytes(object.addr)
}
