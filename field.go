package gtly

import (
	"github.com/viant/toolbox"
	"github.com/viant/xunsafe"
	"reflect"
	"time"
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
	xField        *xunsafe.Field
	outputName    string
	hidden        bool
	kind          reflect.Kind
}

//IsEmpty returns true if field value is empty
func (f *Field) IsEmpty(proto *Proto, value interface{}) bool {
	if value == nil || value == NilValue {
		return true
	}
	if !f.ShallOmitEmpty(proto) {
		return false
	}
	if nillable, ok := value.(Nilable); ok {
		return nillable.IsNil()
	}
	switch f.DataType {
	case FieldTypeBool:
		if !toolbox.AsBoolean(value) {
			return true
		}
	case FieldTypeInt:
		if toolbox.AsInt(value) == 0 {
			return true
		}
	case FieldTypeFloat:
		if toolbox.AsFloat(value) == 0 {
			return true
		}
	case FieldTypeString:
		if toolbox.AsString(value) == "" {
			return true
		}
	}
	if toolbox.IsSlice(value) {
		return reflect.ValueOf(value).Len() == 0
	}
	return false
}

//ShallOmitEmpty return true if shall omit empty
func (f *Field) ShallOmitEmpty(proto *Proto) bool {
	if f.OmitEmpty == nil {
		return proto.OmitEmpty
	}
	return *f.OmitEmpty
}

//TimeLayout returns timelayout
func (f *Field) TimeLayout(proto *Proto) string {
	if f.DataLayout == "" {
		return proto.timeLayout
	}
	return f.DataLayout
}

//InitType initialise filed type
func (f *Field) InitType(value interface{}) {
	if value == nil {
		return
	}
	switch val := value.(type) {
	case *Object:
		f.DataType = FieldTypeObject
		if val == nil {
			f.provider = NewProvider(f.Name)
			return
		}
		f.provider = &Provider{Proto: val.Proto()}
		return
	case *Array:
		f.DataType = FieldTypeArray
		if val == nil {
			f.provider = NewProvider(f.Name)
			return
		}
		f.provider = &Provider{Proto: val.Proto()}
		return
	case *Map:
		f.DataType = FieldTypeArray
		if val == nil {
			f.provider = NewProvider(f.Name)
			return
		}
		f.provider = &Provider{Proto: val.Proto()}
		return
	case *Multimap:
		f.DataType = FieldTypeArray
		if val == nil {
			f.provider = NewProvider(f.Name)
			return
		}
		f.provider = &Provider{Proto: val.Proto()}
		return
	case time.Time, *time.Time, **time.Time, string, []byte:
		f.DataType = getBaseTypeName(value)
		return

	default:
		f.DataType = getBaseTypeName(value)
	}

	if toolbox.IsMap(value) || toolbox.IsStruct(value) {
		f.provider = NewProvider(f.Name)
		f.DataType = FieldTypeObject
		return
	}
	if toolbox.IsSlice(value) {
		f.provider = NewProvider(f.Name)
		f.DataType = FieldTypeArray
		componentType := toolbox.DiscoverComponentType(value)
		componentValue := reflect.New(componentType).Interface()
		f.ComponentType = getBaseTypeName(componentValue)
		return
	}

}

//OutputName returns field output Name
func (f *Field) OutputName() string {
	if f.outputName == "" {
		return f.Name
	}
	return f.outputName
}

//Get returns field value
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
	return field
}
