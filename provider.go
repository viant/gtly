package gtly

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/viant/toolbox"
	"github.com/viant/xunsafe"
	"reflect"
)

//Provider provides shares proto data across all dynamic types
type Provider struct {
	*Proto
}

//NewObject creates an object
func (p *Provider) NewObject() *Object {
	pType := p.Type()
	instance := reflect.New(pType)
	return &Object{
		value: instance,
		proto: p.Proto,
		addr:  xunsafe.ValuePointer(&instance),
		setAt: make([]bool, len(p.Fields())),
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

func (p *Provider) UnMarshall(data []byte) (*Object, error) {
	resultMap := new(map[string]interface{})
	err := json.Unmarshal(data, resultMap)
	if err != nil {
		return nil, err
	}
	anObject := p.NewObject()
	anObject.Init(*resultMap)
	return anObject, err
}
