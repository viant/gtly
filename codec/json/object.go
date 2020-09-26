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
		value := o.ValueAt(field.Index)
		omitEmpty := field.ShallOmitEmpty(o.Proto())
		if omitEmpty {
			empty := field.IsEmpty(o.Proto(), value)
			if empty {
				continue
			}
		}

		if value == nil {
			enc.AddNullKey(field.OutputName())
			continue
		}
		if field.DataType == "" {
			field.InitType(value)
		}
		o.encodeJSONValue(field, value, enc)
	}
}

func (o *Object) encodeJSONValue(field *gtly.Field, value interface{}, enc *gojay.Encoder) {
	filedName := field.OutputName()
	switch field.DataType {
	case gtly.FieldTypeInt:
		enc.IntKey(filedName, toolbox.AsInt(value))
		return
	case gtly.FieldTypeFloat:
		enc.FloatKey(filedName, toolbox.AsFloat(value))
		return
	case gtly.FieldTypeBool:

		enc.BoolKey(filedName, toolbox.AsBoolean(value))
		return
	case gtly.FieldTypeBytes:
		bs, ok := value.([]byte)
		if ok {
			value = base64.StdEncoding.EncodeToString(bs)
		}
		return
	case gtly.FieldTypeArray:

		var marshaler gojay.MarshalerJSONArray
		collection, ok := value.(gtly.Collection)
		if ok {
			switch val := collection.(type) {
			case *gtly.Array:
				if val == nil {
					return
				}
			case *gtly.Map:
				if val == nil {
					return
				}
			case *gtly.Multimap:
				if val == nil {
					return
				}
			}
			marshaler = &Collection{collection}
		} else {
			marshaler = NewSlice(value)
		}
		enc.ArrayKeyOmitEmpty(filedName, marshaler)

		return
	case gtly.FieldTypeObject:
		object, ok := value.(*gtly.Object)
		if !ok {
			provider := gtly.NewProvider("")
			provider.SetOmitEmpty(o.Proto().OmitEmpty)
			var err error
			object, err = provider.Object(value)
			if err != nil {
				return
			}
		}
		marshaler := &Object{object}
		enc.ObjectKeyOmitEmpty(filedName, marshaler)
		return
	case gtly.FieldTypeTime:
		timeLayout := field.TimeLayout(o.Proto())
		if timeLayout != "" {
			if timeValue, err := toolbox.ToTime(value, timeLayout); err == nil {
				value = timeValue.Format(timeLayout)
			}
		}
	}
	if field.ShallOmitEmpty(o.Proto()) {
		enc.StringKeyOmitEmpty(filedName, toolbox.AsString(value))
		return
	}
	enc.StringKey(filedName, toolbox.AsString(value))
}
