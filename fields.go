package gtly

import (
	"fmt"
	"github.com/viant/toolbox"
	"reflect"
)

//Fields represents a fields
type Fields []*Field

//MapFields creates fields for supplied values
func MapFields(source interface{}) (Fields, error) {
	if aMap, ok := source.(map[string]interface{}); ok {
		var fields = make(Fields, len(aMap))
		i := 0
		for k, v := range aMap {
			opt, err := ValueOpt(v)
			if err != nil {
				return nil, err
			}
			fields[i] = NewField(k, "", opt)
			i++
		}
		return fields, nil
	}

	if toolbox.IsStruct(source) {
		var fields = make(Fields, 0)
		return fields, toolbox.ProcessStruct(source, func(fieldType reflect.StructField, field reflect.Value) error {
			fields = append(fields, &Field{
				Name: fieldType.Name,
				Type: fieldType.Type,
			})
			return nil
		})
	}
	return nil, fmt.Errorf("unsupported type: %v", source)
}
