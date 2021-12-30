package json

import (
	"encoding/base64"
	"github.com/francoispqt/gojay"
	"github.com/viant/gtly"
	"github.com/viant/toolbox"
)

//Object JSON wrapper
type Object struct {
	*gtly.Object
}

//MarshalJSONObject converts an object into JSON object
func (o Object) MarshalJSONObject(enc *gojay.Encoder) {
	fields := o.Proto().Fields()
	for _, field := range fields {
		value, ok := o.ValueAt(field.Index)
		omitEmpty := field.ShallOmitEmpty()
		if omitEmpty && !ok {
			continue
		}

		if value == nil {
			enc.AddNullKey(field.OutputName())
			continue
		}
		if err := o.encodeJSONValue(field, value, enc); err != nil {
			enc.AddInterface(err)
		}
	}
}

func (o *Object) encodeJSONValue(field *gtly.Field, value interface{}, enc *gojay.Encoder) error {
	filedName := field.OutputName()
	switch field.DataType {
	case gtly.FieldTypeInt:
		enc.IntKey(filedName, toolbox.AsInt(value))
		return nil
	case gtly.FieldTypeFloat32:
		enc.FloatKey(filedName, toolbox.AsFloat(value))
		return nil
	case gtly.FieldTypeBool:

		enc.BoolKey(filedName, toolbox.AsBoolean(value))
		return nil
	case gtly.FieldTypeBytes:
		bs, ok := value.([]byte)
		if ok {
			value = base64.StdEncoding.EncodeToString(bs)
		}
		return nil
	case gtly.FieldTypeArray:

		var marshaler gojay.MarshalerJSONArray
		collection, ok := value.(gtly.Collection)
		if ok {
			switch val := collection.(type) {
			case *gtly.Array:
				if val == nil {
					return nil
				}
			case *gtly.Map:
				if val == nil {
					return nil
				}
			case *gtly.Multimap:
				if val == nil {
					return nil
				}
			}
			marshaler = &Collection{collection}
		} else {
			marshaler = NewSlice(value)
		}
		enc.ArrayKeyOmitEmpty(filedName, marshaler)

		return nil
	case gtly.FieldTypeObject:
		object, ok := value.(*gtly.Object)
		if !ok {
			provider, err := gtly.NewProvider("")
			if err != nil {
				return err
			}
			provider.SetOmitEmpty(o.Proto().OmitEmpty)
			object, err = provider.Object(value)
			if err != nil {
				return err
			}
		}
		marshaler := &Object{object}
		enc.ObjectKeyOmitEmpty(filedName, marshaler)
		return nil
	case gtly.FieldTypeTime:
		timeLayout := field.TimeLayout()
		if timeLayout != "" {
			if timeValue, err := toolbox.ToTime(value, timeLayout); err == nil {
				value = timeValue.Format(timeLayout)
			}
		}
	}
	if field.ShallOmitEmpty() {
		enc.StringKeyOmitEmpty(filedName, toolbox.AsString(value))
		return nil
	}
	enc.StringKey(filedName, toolbox.AsString(value))
	return nil
}
