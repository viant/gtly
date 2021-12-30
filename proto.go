package gtly

import (
	"github.com/viant/toolbox/format"
	"github.com/viant/xunsafe"
	"reflect"
	"strings"
	"time"
)

var defaultEmptyValues = map[interface{}]bool{
	"":  true,
	nil: true,
}

//Proto represents generic type prototype
type Proto struct {
	Name       string
	fieldNames map[string]int
	fields     []Field
	accessors  []Accessor
	mutators   []Mutator

	OmitEmpty        bool
	emptyValues      map[interface{}]bool
	timeLayout       string
	caseFormat       format.Case
	outputCaseFormat format.Case
	inputCaseFormat  format.Case
	dataType         reflect.Type
	xType            *xunsafe.Type
	kind             reflect.Kind
}

//Type returns proto data type
func (p *Proto) Type() reflect.Type {
	return p.dataType
}

func (p *Proto) buildType() reflect.Type {
	structFields := make([]reflect.StructField, len(p.fields))
	for i, field := range p.fields {
		if field.Type == nil {
			field.Type = typeString
		}
		structFields[i] = reflect.StructField{
			Name: field.Name,
			Type: field.Type,
		}
		p.fields[i].kind = structFields[i].Type.Kind()
	}
	dataType := reflect.StructOf(structFields)
	for i := range p.fields {
		p.initField(&p.fields[i], dataType)
	}
	return dataType
}

func (p *Proto) initField(field *Field, dataType reflect.Type) {
	xField := xunsafe.FieldByName(dataType, field.Name)
	if p.caseFormat != p.outputCaseFormat {
		field.outputName = p.caseFormat.Format(field.Name, p.outputCaseFormat)
	}

	p.accessors[field.Index].init(field.Index, xField)
	p.mutators[field.Index].init(field.Index, xField)
	p.indexByNames(field)
}

func (p *Proto) indexByNames(field *Field) {
	p.fieldNames[field.Name] = field.Index
	if field.InputName != "" && field.InputName != field.Name {
		p.fieldNames[field.InputName] = field.Index
	}
	if field.outputName != "" && field.outputName != field.Name {
		p.fieldNames[field.outputName] = field.Index
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

//Hide set hidden flag for the Field
func (p *Proto) Hide(name string) {
	field := p.Field(name)
	if field == nil {
		return
	}
	field.hidden = true
}

//Show remove hidden flag for supplied Field
func (p *Proto) Show(name string) {
	field := p.Field(name)
	if field == nil {
		return
	}
	field.hidden = false
}

//Size returns proto size
func (p *Proto) Size() int {
	return len(p.fields)
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
func (p *Proto) Fields() []Field {
	return p.fields
}

//Field returns Field for specified Name
func (p *Proto) Field(name string) *Field {
	index := p.fieldNames[name]
	return &p.fields[index]
}

//FieldAt returns Field at position
func (p *Proto) FieldAt(index int) *Field {
	return &p.fields[index]
}

//Mutator returns field mutator
func (p *Proto) Mutator(fieldName string) *Mutator {
	field := p.Field(fieldName)
	return &p.mutators[field.Index]
}

//MutatorAt returns field mutator
func (p *Proto) MutatorAt(index int) *Mutator {
	return &p.mutators[index]
}

//Accessor returns a field accessor
func (p *Proto) Accessor(fieldName string) *Accessor {
	field := p.Field(fieldName)
	return &p.accessors[field.Index]
}

//AccessorAt returns a field accessor
func (p *Proto) AccessorAt(index int) *Accessor {
	return &p.accessors[index]
}

//newProto create a data type prototype
func newProto(name string, fields []*Field) *Proto {
	result := &Proto{
		Name:       name,
		fieldNames: make(map[string]int),
		fields:     make([]Field, len(fields)),
		accessors:  make([]Accessor, len(fields)),
		mutators:   make([]Mutator, len(fields)),
	}
	result.timeLayout = time.RFC3339
	for i := range fields {
		result.fields[i] = *fields[i]
	}
	result.dataType = result.buildType()
	result.xType = xunsafe.NewType(result.dataType)

	return result
}
