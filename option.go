package gtly

import (
	"fmt"
	"github.com/viant/toolbox"
	"reflect"
)

//Option represents field option
type Option func(field *Field)

//DateLayoutOpt field with data layout option
func DateLayoutOpt(layout string) Option {
	return func(field *Field) {
		field.DataLayout = layout
	}
}

//ProviderOpt return a field provider option
func ProviderOpt(provider *Provider) Option {
	return func(field *Field) {
		field.provider = provider
	}
}

//OmitEmptyOpt returns a field omit empty option
func OmitEmptyOpt(omitEmpty bool) Option {
	return func(field *Field) {
		field.OmitEmpty = &omitEmpty
	}
}

//ComponentTypeOpt return a field component type option
func ComponentTypeOpt(componentType string) Option {
	return func(field *Field) {
		field.ComponentType = componentType
	}
}

//ValueOpt derives type from supplied value
func ValueOpt(value interface{}) (Option, error) {
	if value == nil {
		return nil, fmt.Errorf("value was empty")
	}

	switch val := value.(type) {
	case *Object:
		return func(field *Field) {
			field.DataType = FieldTypeObject
			field.provider = &Provider{Proto: val.Proto()}
		}, nil
	case *Array:
		return func(field *Field) {
			field.DataType = FieldTypeArray
			field.provider = &Provider{Proto: val.Proto()}
		}, nil
	case *Map:
		return func(field *Field) {
			field.DataType = FieldTypeArray
			field.provider = &Provider{Proto: val.Proto()}
		}, nil
	case *Multimap:
		return func(field *Field) {
			field.DataType = FieldTypeArray
			field.provider = &Provider{Proto: val.Proto()}
		}, nil

	default:

		typeName := typeNameForValue(value)
		if typeName != "" {
			rType := getBaseType(typeName)
			return func(field *Field) {
				if field.Type == nil {
					field.Type = rType
				}
				if field.DataType == "" {
					field.DataType = typeName
				}
			}, nil
		}

		if toolbox.IsMap(value) || toolbox.IsStruct(value) {
			fields, err := MapFields(value)
			if err != nil {
				return nil, err
			}
			provider, err := NewProvider("", fields...)
			if err != nil {
				return nil, err
			}
			return func(field *Field) {
				provider.Name = field.Name
				field.provider = provider
				field.Type = provider.dataType
			}, nil
		}

		if toolbox.IsSlice(value) {
			componentType := toolbox.DiscoverComponentType(value)
			typeName := typeNameForType(componentType)
			if typeName != "" {
				rType := getBaseType(typeName)
				return func(field *Field) {
					field.Type = reflect.SliceOf(rType)
					field.DataType = FieldTypeArray
				}, nil
			}
			fields, err := MapFields(reflect.New(componentType).Interface())
			if err != nil {
				return nil, err
			}
			provider, err := NewProvider("", fields...)
			if err != nil {
				return nil, err
			}
			return func(field *Field) {
				field.Type = provider.Type()
				field.DataType = FieldTypeArray
			}, nil
		}
	}
	return nil, fmt.Errorf("unsupported type %T", value)
}
