package json

import (
	"github.com/francoispqt/gojay"
	"github.com/viant/gtly"
	"github.com/viant/toolbox"
	"reflect"
)

//Slice represents a primitive JSON slice
type Slice struct {
	_data interface{}
}

//IsNil returns true if empty
func (s Slice) IsNil() bool {
	if s._data == nil {
		return true
	}
	var sliceLen = reflect.ValueOf(s._data).Len()
	return sliceLen == 0
}

//MarshalJSONArray converts primitive collection into JSON array
func (s Slice) MarshalJSONArray(enc *gojay.Encoder) {
	toolbox.ProcessSlice(s._data, func(item interface{}) bool {
		if item == nil {
			enc.AddNull()
			return true
		}
		fields, err := gtly.MapFields(item)
		if err != nil {
			enc.AddInterface(item) // item has to be primitive
			return true
		}
		provider, err := gtly.NewProvider("foo", fields...)
		if object, err := provider.Object(item); err == nil {
			item = &Object{object}
		}
		enc.AddInterface(item)
		return true
	})
}

//NewSlice creates a slice
func NewSlice(data interface{}) *Slice {
	return &Slice{
		_data: data,
	}
}
