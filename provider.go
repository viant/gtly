package gtly

import (
	"encoding/json"
	"github.com/viant/xunsafe"
	"reflect"
)

//Provider provides shares proto data across all dynamic types
type Provider struct {
	*Proto
}

//NewObject creates an object
func (p *Provider) NewObject() *Object {
	pType := p.dataType
	instance := reflect.New(pType)
	obj := &Object{
		value: instance,
		proto: p.Proto,
		addr:  xunsafe.ValuePointer(&instance),
		setAt: make([]bool, len(p.Fields())),
	}
	return obj
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
	return result, result.Set(value)
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
func NewProvider(name string, fields ...*Field) (*Provider, error) {
	for _, field := range fields {
		field.init()
	}
	p := &Provider{Proto: newProto(name, fields...)}
	for _, field := range fields {
		field.provider = p
	}
	return p, nil
}

//UnMarshall unmarshal objects
func (p *Provider) UnMarshall(data []byte) (*Object, error) {
	var resultMap = make(map[string]interface{})
	err := json.Unmarshal(data, &resultMap)
	if err != nil {
		return nil, err
	}
	anObject := p.NewObject()
	return anObject, anObject.Set(resultMap)
}
