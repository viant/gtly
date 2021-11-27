package gtly

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/viant/toolbox"
	"reflect"
	"unsafe"
)

//Provider provides shares _proto data across all dynamic types
type Provider struct {
	*Proto
}

//NewObject creates an object
func (p *Provider) NewObject() *Object {
	instance := reflect.New(p.Type())

	return &Object{
		value:     instance,
		_proto:    p.Proto,
		_dataAddr: unsafe.Pointer(instance.Elem().UnsafeAddr()),
		_setAt:    make([]bool, len(p.Fields())),
	}
}

//NewArray creates a slice
func (p *Provider) NewArray(items ...*Object) *Array {
	return &Array{
		_provider: p,
		_data:     items,
	}
}

//Object creates an object from struct or map
func (p *Provider) Object(value interface{}) (*Object, error) {
	result := p.NewObject()
	if toolbox.IsStruct(value) {
		return result, toolbox.ProcessStruct(value, func(fieldType reflect.StructField, field reflect.Value) error {
			result.SetValue(fieldType.Name, field.Interface())
			return nil
		})
	}
	if toolbox.IsMap(value) {
		toolbox.ProcessMap(value, func(key, value interface{}) bool {
			result.SetValue(toolbox.AsString(key), value)
			return true
		})
		return result, nil
	}
	return nil, errors.Errorf("unsupported object source: %T", value)
}

//NewMap creates a map of string and object
func (p *Provider) NewMap(keyProvider KeyProvider) *Map {
	return &Map{
		_map:        map[interface{}]*Object{},
		_provider:   p,
		keyProvider: keyProvider,
	}
}

//NewMultimap creates a multimap of string and slice
func (p *Provider) NewMultimap(keyProvider KeyProvider) *Multimap {
	return &Multimap{
		_map:        map[interface{}][]*Object{},
		_provider:   p,
		keyProvider: keyProvider,
	}
}

//NewProvider creates provider
func NewProvider(name string, fields ...*Field) *Provider {
	return &Provider{Proto: newProto(name, fields...)}
}

func (p *Provider) UnMarshall(data []byte) *Object {
	resultMap := new(map[string]interface{})
	json.Unmarshal(data, resultMap)
	anObject := p.NewObject()
	anObject.Init(*resultMap)
	return anObject
}
