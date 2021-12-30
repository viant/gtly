package gtly

import (
	"github.com/viant/toolbox"
	"reflect"
	"time"
)

const ( //Data type definition
	//FieldTypeInt int type
	FieldTypeInt = "int"
	//FieldTypeInt64 int type
	FieldTypeInt64 = "int64"
	//FieldTypeFloat32 float type
	FieldTypeFloat32 = "float32"
	//FieldTypeFloat64 float type
	FieldTypeFloat64 = "float64"
	//FieldTypeBool bool type
	FieldTypeBool = "bool"
	//FieldTypeString string type
	FieldTypeString = "string"
	//FieldTypeTime time type
	FieldTypeTime = "time"
	//FieldTypeBytes bytes type
	FieldTypeBytes = "bytes"
	//FieldTypeArray array type
	FieldTypeArray = "array"
	//FieldTypeObject object type
	FieldTypeObject = "object"
)

var (
	typeInt     = reflect.TypeOf(0)
	typeInt64   = reflect.TypeOf(int64(0))
	typeFloat   = reflect.TypeOf(float32(0))
	typeFloat64 = reflect.TypeOf(float64(0))
	typeBool    = reflect.TypeOf(true)
	typeString  = reflect.TypeOf("")
	typeBytes   = reflect.TypeOf([]byte(""))
	typeTime    = reflect.TypeOf(time.Time{})
	typeTimePtr = reflect.TypeOf(&time.Time{})
)

//typeNameForValue returns base type
func typeNameForType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Float32:
		return FieldTypeFloat32
	case reflect.Float64:
		return FieldTypeFloat64
	case reflect.Int, reflect.Uint64, reflect.Uint:
		return FieldTypeInt
	case reflect.Int64:
		return FieldTypeInt64
	case reflect.Bool:
		return FieldTypeBool
	case reflect.Struct:
		if t == typeTime {
			return FieldTypeTime
		}
	case reflect.Slice:
		switch t.Elem().Kind() {
		case reflect.Uint8:
			return FieldTypeBytes
		}
	case reflect.Ptr:
		switch t.Elem().Kind() {
		case reflect.Float32:
			return FieldTypeFloat32
		case reflect.Float64:
			return FieldTypeFloat64
		case reflect.Int, reflect.Uint64, reflect.Uint:
			return FieldTypeInt
		case reflect.Int64:
			return FieldTypeInt64
		case reflect.Bool:
			return FieldTypeBool
		case reflect.Struct:
			if t == typeTimePtr {
				return FieldTypeTime
			}
		}
	}
	return FieldTypeString
}

//typeNameForValue returns base type
func typeNameForValue(value interface{}) string {
	switch val := value.(type) {
	case float32, *float32:
		return FieldTypeFloat32
	case float64, *float64:
		return FieldTypeFloat64
	case int, int8, int16, int32, uint, uint8, uint16, uint32, uint64, *int, *int8, *int16, *int32, *uint, *uint8, *uint16, *uint32, *uint64:
		return FieldTypeInt
	case int64, *int64:
		return FieldTypeInt64
	case time.Time, *time.Time:
		return FieldTypeTime
	case bool, *bool:
		return FieldTypeBool
	case []byte:
		if _, err := toolbox.ToFloat(val); err == nil {
			return FieldTypeFloat32
		}
	}
	return FieldTypeString
}

//getBaseType returns base type for supplied name
func getBaseType(typeName string) reflect.Type {
	switch typeName {
	case FieldTypeInt:
		return typeInt
	case FieldTypeInt64:
		return typeInt64
	case FieldTypeFloat32:
		return typeFloat
	case FieldTypeFloat64:
		return typeFloat64
	case FieldTypeBool:
		return typeBool
	case FieldTypeString:
		return typeString
	case FieldTypeTime:
		return typeTime
	case FieldTypeBytes:
		return typeBytes
	}
	return nil
}
