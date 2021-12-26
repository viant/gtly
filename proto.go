package gtly

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/viant/toolbox/format"
	"github.com/viant/xunsafe"
	"reflect"
	"strings"
	"sync"
	"time"
)

var defaultEmptyValues = map[interface{}]bool{
	"":  true,
	nil: true,
}

//Proto represents generic type prototype
type Proto struct {
	Name             string
	lock             *sync.RWMutex
	fieldNames       map[string]*Field
	fields           []*Field
	nilTypes         []int
	OmitEmpty        bool
	emptyValues      map[interface{}]bool
	timeLayout       string
	caseFormat       format.Case
	outputCaseFormat format.Case
	inputCaseFormat  format.Case
	dataType         reflect.Type
	kind             reflect.Kind
}

func (p *Proto) Type() reflect.Type {
	p.buildTypeIfNeeded()
	return p.dataType
}

func (p *Proto) buildTypeIfNeeded() {
	if p.dataType == nil {
		p.buildType()
	}
}

func (p *Proto) buildType() {
	fields := p.fields
	structFields := make([]reflect.StructField, len(fields))
	for i, field := range fields {
		structFields[i] = reflect.StructField{
			Name: field.Name,
			Type: field.Type,
		}
		fields[i].kind = field.Type.Kind()
	}
	p.dataType = reflect.StructOf(structFields)
	for i, field := range fields {
		fields[i].xField = xunsafe.FieldByName(p.dataType, field.Name)
	}

}

//SimpleName returns simple name
func (p *Proto) SimpleName() string {
	if index := strings.LastIndex(p.Name, "."); index != -1 {
		return p.Name[index+1:]
	}
	return p.Name
}

//SetOmitEmpty sets omit empty flag
func (p *Proto) SetOmitEmpty(omitEmpty bool) {
	p.OmitEmpty = true
	if omitEmpty {
		if len(p.emptyValues) == 0 {
			p.emptyValues = defaultEmptyValues
		}
	}
}

//SetEmptyValues sets empty values, use only if empty values are non in default map: nil, empty string
func (p *Proto) SetEmptyValues(values ...interface{}) {
	p.emptyValues = make(map[interface{}]bool)
	for i := range values {
		p.emptyValues[values[i]] = true
	}
}

//OutputCaseFormat set output case format
func (p *Proto) OutputCaseFormat(source, output format.Case) error {
	for i, field := range p.fields {
		p.fields[i].outputName = source.Format(field.Name, output)
	}
	return nil
}

//InputCaseFormat set output case format
func (p *Proto) InputCaseFormat(source, input format.Case) error {
	p.inputCaseFormat = source
	for i, field := range p.fields {
		p.fields[i].outputName = source.Format(field.Name, input)
	}
	return nil
}

//Hide set hidden flag for the field
func (p *Proto) Hide(name string) {
	field := p.Field(name)
	if field == nil {
		return
	}
	field.hidden = true
}

//Show remove hidden flag for supplied field
func (p *Proto) Show(name string) {
	field := p.Field(name)
	if field == nil {
		return
	}
	field.hidden = false
}

//Size returns _proto size
func (p *Proto) Size() int {
	p.lock.RLock()
	result := len(p.fields)
	p.lock.RUnlock()
	return result
}

func (p *Proto) asMap(values []interface{}) map[string]interface{} {
	var result = make(map[string]interface{})
	for _, field := range p.fields {
		if field.hidden {
			continue
		}
		var value interface{}
		if field.Index < len(values) {
			value = values[field.Index]
		}
		value = Value(value)
		fieldName := field.Name
		if field.outputName != "" {
			fieldName = field.outputName
		}
		result[fieldName] = value
	}
	return result
}

func reallocateIfNeeded(size int, data []interface{}) []interface{} {
	if size >= len(data) {
		for i := len(data); i < size; i++ {
			data = append(data, nil)
		}
	}
	return data
}

//Fields returns fields list
func (p *Proto) Fields() []*Field {
	return p.fields
}

//Field returns field for specified Name
func (p *Proto) Field(name string) *Field {
	field := p.fieldNames[name]
	return field
}

func (p *Proto) FieldAt(index int) *Field {
	return p.fields[index]
}

//Object creates an object
func (p *Proto) Object(values []interface{}) (*Object, error) {
	if len(p.fields) < len(values) {
		return nil, errors.Errorf("invalid value count: %v, field count: %v", len(values), len(p.fields))
	}

	object := &Object{_proto: p}
	if len(p.nilTypes) > 0 {
		for _, index := range p.nilTypes {
			if values[index] != nil {
				p.fields[index].InitType(values[index])
			}
		}
		p.updateNilTypes()
	}
	return object, nil
}

//FieldWithValue returns existing filed , or create a new field
func (p *Proto) FieldWithValue(fieldName string, value interface{}) *Field {
	p.lock.RLock()
	field, ok := p.fieldNames[fieldName]
	p.lock.RUnlock()
	if ok {
		return field
	}

	field = &Field{Name: fieldName, Index: len(p.fieldNames)}
	if p.inputCaseFormat != p.caseFormat {
		field.InputName = field.Name
		field.Name = p.inputCaseFormat.Format(fieldName, p.caseFormat)
	}
	field.InitType(value)
	if value == nil {
		p.addNilType(field.Index)
	}

	return p.AddField(field)
}

//AddField add fields
func (p *Proto) AddField(field *Field) *Field {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.dataType != nil {
		panic(fmt.Sprintf("proto has been already in use"))
	}
	if p.caseFormat != p.outputCaseFormat {
		field.outputName = p.caseFormat.Format(field.Name, p.outputCaseFormat)
	}
	p.fieldNames[field.Name] = field
	if field.InputName != "" {
		p.fieldNames[field.InputName] = field
	}
	if field.outputName != "" {
		p.fieldNames[field.outputName] = field
	}
	p.fields = append(p.fields, field)
	field.Index = len(p.fields) - 1
	return field
}

func (p *Proto) updateNilTypes() {
	p.nilTypes = make([]int, 0)
	for i := range p.fields {
		if p.fields[i].DataType == "" {
			p.nilTypes = append(p.nilTypes, p.fields[i].Index)
		}
	}
}

func (p *Proto) addNilType(index int) {
	if len(p.nilTypes) == 0 {
		p.nilTypes = make([]int, 0)
	}
	p.nilTypes = append(p.nilTypes, index)
}

//newProto create a data type prototype
func newProto(name string, fields ...*Field) *Proto {
	result := &Proto{
		Name:       name,
		lock:       &sync.RWMutex{},
		fieldNames: make(map[string]*Field),
		fields:     make([]*Field, 0),
	}
	result.timeLayout = time.RFC3339
	for i := range fields {
		result.AddField(fields[i])
		fields[i].Index = i
	}
	return result
}
