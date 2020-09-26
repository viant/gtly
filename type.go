package gtly

import (
	"github.com/viant/toolbox"
	"time"
)

const ( //Data type definition
	//FieldTypeInt int type
	FieldTypeInt = "int"
	//FieldTypeFloat float type
	FieldTypeFloat = "float"
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

//getBaseType returns base type
func getBaseType(value interface{}) string {
	switch val := value.(type) {
	case float32, float64, *float32, *float64:
		return FieldTypeFloat
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64:
		return FieldTypeInt
	case time.Time, *time.Time:
		return FieldTypeTime
	case bool, *bool:
		return FieldTypeBool
	case []byte:
		if _, err := toolbox.ToFloat(val); err == nil {
			return FieldTypeFloat
		}
	}
	return FieldTypeString
}
