package gtly

import (
	"reflect"
)

//NilValue is used to discriminate between unset fileds, and set filed with nil value (for REST patch operation)
var NilValue = make([]*interface{}, 1)[0]

//Field represents dynamic filed
type Field struct {
	Name          string `json:",omitempty"`
	Index         int
	StructTag     reflect.StructTag
	OmitEmpty     *bool        `json:",omitempty"`
	DateFormat    string       `json:",omitempty"`
	DataLayout    string       `json:",omitempty"`
	DataType      string       `json:",omitempty"`
	InputName     string       `json:",omitempty"`
	ComponentType string       `json:",omitempty"`
	Type          reflect.Type `json:"-"`
	provider      *Provider
	outputName    string
	hidden        bool
	kind          reflect.Kind
}

//ShallOmitEmpty return true if shall omit empty
func (f *Field) ShallOmitEmpty() bool {
	if f.OmitEmpty == nil {
		return f.provider.Proto.OmitEmpty
	}
	return *f.OmitEmpty
}

//TimeLayout returns timelayout
func (f *Field) TimeLayout() string {
	if f.DataLayout == "" {
		return f.provider.Proto.timeLayout
	}
	return f.DataLayout
}

func (f *Field) init(index int, provider *Provider) {
	f.Index = index
	f.provider = provider
	if f.Type == nil && f.DataType != "" {
		f.Type = getBaseType(f.DataType)
	}
	if f.DataType == "" && f.Type != nil {
		f.DataType = typeNameForType(f.Type)
	}
}

//OutputName returns Field output Name
func (f *Field) OutputName() string {
	if f.outputName == "" {
		return f.Name
	}
	return f.outputName
}

//Get returns Field value
func (f *Field) Get(values []interface{}) interface{} {
	if f.Index < len(values) {
		return Value(values[f.Index])
	}
	return nil
}

//SetProvider set provider
func (f *Field) SetProvider(provider *Provider) {
	f.provider = provider
}

//NewField creates new fields
func NewField(name, dataType string, options ...Option) *Field {
	field := &Field{
		Name:     name,
		DataType: dataType,
	}
	for _, option := range options {
		option(field)
	}
	if field.Type == nil {
		field.Type = getBaseType(field.DataType)
	}
	if field.provider != nil {
		if field.Type == nil {
			switch field.DataType {
			case FieldTypeArray:
				field.Type = reflect.SliceOf(field.provider.Type())
			case FieldTypeObject:
				field.Type = field.provider.Type()
			}
		}
	}
	return field
}
