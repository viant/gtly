package gtly_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/viant/gtly"
	"github.com/viant/toolbox/format"
	"testing"
	"time"
)

type setMethod string

const (
	noSet              setMethod = "no set"
	byInterfaceMutator setMethod = "by interface mutator"
	byTypeMutator      setMethod = "by type mutator"
	byInitialize       setMethod = "by initialize"
	byValue            setMethod = "by interface"
)

type ObjTestCase struct {
	description      string
	fields           map[string]ObjectTestField
	isNil            bool
	outputCaseFormat format.Case
	asMap            map[string]interface{}
}

type ObjectTestField struct {
	value       interface{}
	index       int
	hidden      bool
	set         setMethod
	expectedSet bool
}

func TestObject(t *testing.T) {
	testCases := []ObjTestCase{
		{
			outputCaseFormat: format.CaseLowerCamel,
			description:      "set by value",
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         byValue,
					expectedSet: true,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         byValue,
					expectedSet: true,
					hidden:      true,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         byValue,
					expectedSet: true,
				},
				"Prop4": {
					value:       true,
					index:       3,
					set:         byValue,
					expectedSet: true,
				},
				"Prop5": {
					value:       getIsoDate("2021-11-01"),
					index:       4,
					set:         byValue,
					expectedSet: true,
				},
				"Prop6": {
					value:       "",
					index:       5,
					set:         noSet,
					expectedSet: false,
				},
			},
			isNil: false,
			asMap: map[string]interface{}{
				"prop1": "abc",
				"prop3": 4,
				"prop4": true,
				"prop5": getIsoDate("2021-11-01"),
			},
		},
		{
			description:      "initialize object",
			outputCaseFormat: format.CaseLowerCamel,
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         byInitialize,
					expectedSet: true,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         byInitialize,
					expectedSet: true,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         byInitialize,
					expectedSet: true,
				},
				"Prop4": {
					value:       "",
					index:       3,
					set:         byInitialize,
					expectedSet: true,
				},
				"Prop5": {
					value:       getIsoDate("2021-11-01"),
					index:       4,
					set:         byInitialize,
					expectedSet: true,
				},
			},
			isNil: false,
			asMap: map[string]interface{}{
				"prop1": "abc",
				"prop2": 3.5,
				"prop3": 4,
				"prop4": "",
				"prop5": getIsoDate("2021-11-01"),
			},
		},
		{
			description:      "set by interface mutator",
			outputCaseFormat: format.CaseLowerCamel,
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         byInterfaceMutator,
					expectedSet: true,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         byInterfaceMutator,
					expectedSet: true,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         byInterfaceMutator,
					expectedSet: true,
				},
				"Prop4": {
					value:       "",
					index:       3,
					set:         byInterfaceMutator,
					expectedSet: true,
				},
				"Prop5": {
					value:       getIsoDate("2021-11-01"),
					index:       4,
					set:         byInterfaceMutator,
					expectedSet: true,
				},
			},
			isNil: false,
			asMap: map[string]interface{}{
				"prop1": "abc",
				"prop2": 3.5,
				"prop3": 4,
				"prop4": "",
				"prop5": getIsoDate("2021-11-01"),
			},
		},
		{
			outputCaseFormat: format.CaseUpperCamel,
			description:      "set by type mutator",
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         byTypeMutator,
					expectedSet: true,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         byTypeMutator,
					expectedSet: true,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         byTypeMutator,
					expectedSet: true,
				},
				"Prop4": {
					value:       "",
					index:       3,
					set:         byTypeMutator,
					expectedSet: true,
				},
				"Prop5": {
					value:       getIsoDate("2021-11-01"),
					index:       4,
					set:         byTypeMutator,
					expectedSet: true,
				},
			},
			isNil: false,
			asMap: map[string]interface{}{
				"Prop1": "abc",
				"Prop2": 3.5,
				"Prop3": 4,
				"Prop4": "",
				"Prop5": getIsoDate("2021-11-01"),
			},
		},
		{
			description:      "initialize and set",
			outputCaseFormat: format.CaseUpper,
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         byInitialize,
					expectedSet: true,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         byInitialize,
					expectedSet: true,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         byInitialize,
					expectedSet: true,
				},
				"Prop4": {
					value:       "",
					index:       3,
					set:         byInitialize,
					expectedSet: true,
				},
				"Prop5": {
					value:       getIsoDate("2021-11-01"),
					index:       4,
					set:         byInitialize,
					expectedSet: true,
				},
				"Prop6": {
					value:       "abc",
					index:       5,
					set:         byTypeMutator,
					expectedSet: true,
				},
				"Prop7": {
					value:       3.5,
					index:       6,
					set:         byValue,
					expectedSet: true,
				},
				"Prop8": {
					value:       4,
					index:       7,
					set:         byInterfaceMutator,
					expectedSet: true,
				},
				"Prop9": {
					value:       "",
					index:       8,
					set:         byTypeMutator,
					expectedSet: true,
				},
				"Prop10": {
					value:       getIsoDate("2021-11-01"),
					index:       9,
					set:         byValue,
					expectedSet: true,
				},
			},
			isNil: false,
			asMap: map[string]interface{}{
				"PROP1":  "abc",
				"PROP2":  3.5,
				"PROP3":  4,
				"PROP4":  "",
				"PROP5":  getIsoDate("2021-11-01"),
				"PROP6":  "abc",
				"PROP7":  3.5,
				"PROP8":  4,
				"PROP9":  "",
				"PROP10": getIsoDate("2021-11-01"),
			},
		},
		{
			description:      "nil object",
			outputCaseFormat: format.CaseLowerCamel,
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         noSet,
					expectedSet: false,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         noSet,
					expectedSet: false,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         noSet,
					expectedSet: false,
				},
			},
			asMap: map[string]interface{}{},
			isNil: true,
		},
	}

	for i, testCase := range testCases {
		fields, err := objectFields(testCase.fields)
		if !assert.Nil(t, err, testCase.description) {
			continue
		}
		provider, err := gtly.NewProvider(fmt.Sprintf("testCase%v", i), fields...)
		if !assert.Nil(t, err, testCase.description) {
			continue
		}
		initTestCase(testCase, provider)
		anObject := provider.NewObject()
		setObjectValues(testCase.fields, anObject)
		runObjectTestFor(t, testCase, anObject)
	}
}

func runObjectTestFor(t *testing.T, testCase ObjTestCase, anObject *gtly.Object) {
	checkObjectValues(t, testCase.fields, anObject, testCase.description)
	assert.Equal(t, testCase.isNil, anObject.IsNil(), testCase.description)
	assert.Equal(t, testCase.asMap, anObject.AsMap(), testCase.description)
}

func initTestCase(testCase ObjTestCase, provider *gtly.Provider) {
	fields := provider.Fields()
	for _, field := range fields {
		testField := testCase.fields[field.Name]
		testField.index = field.Index
		testCase.fields[field.Name] = testField
		if testCase.fields[field.Name].hidden {
			provider.Hide(field.Name)
		}
	}
	provider.OutputCaseFormat(format.CaseUpper, testCase.outputCaseFormat)
}

func checkObjectValues(t *testing.T, values map[string]ObjectTestField, object *gtly.Object, description string) {
	for k, v := range values {
		accessor := object.Proto().Accessor(k)
		if !assert.Equal(t, v.expectedSet, object.SetAt(v.index), description+" "+k) {
			//	fmt.Printf("%d %v %+v\n", v.index, k, object)
		}
		if v.set == noSet {
			continue
		}
		checkObjectByType(t, v.value, k, object, v.index, description)
		assert.Equal(t, v.value, object.Value(k), description)
		value := accessor.Value(object)
		assert.Equal(t, v.value, value, description)
		fieldValue, isValuePresent := object.ValueAt(v.index)
		assert.Equal(t, v.value, fieldValue, description)
		assert.True(t, isValuePresent, description)
	}
}

func checkObjectByType(t *testing.T, value interface{}, name string, object *gtly.Object, index int, description string) {
	accessor := object.Proto().Accessor(name)
	switch fieldValue := value.(type) {
	case int:
		assert.Equal(t, value, accessor.Int(object), description)
	case string:
		assert.Equal(t, value, accessor.String(object), description)
	case bool:
		assert.Equal(t, value, accessor.Bool(object), description)
	case float64:
		assert.Equal(t, value, accessor.Float64(object), description)
	case time.Time:
		assert.Equal(t, value, accessor.Time(object), description)
	default:
		panic(fmt.Errorf("not implemented Accessor for type %T", fieldValue))
	}
}

func setObjectValues(values map[string]ObjectTestField, object *gtly.Object) {
	_ = initializeObjectValues(values, object)
	for k, v := range values {
		mutator := object.Proto().Mutator(k)
		switch v.set {
		case noSet, byInitialize:
			continue
		case byInterfaceMutator:
			mutator.SetValue(object, v.value)
		case byTypeMutator:
			setObjectByTypeMutator(v, object, mutator)
		case byValue:
			object.SetValue(k, v.value)
		default:
			panic(fmt.Errorf("not implemented setting method: %v", v.set))
		}
	}

}

func objectFields(values map[string]ObjectTestField) ([]*gtly.Field, error) {
	var result = make([]*gtly.Field, 0)
	for k, v := range values {
		opt, err := gtly.ValueOpt(v.value)
		if err != nil {
			return nil, err
		}
		result = append(result, gtly.NewField(k, "", opt))
	}
	return result, nil
}

func initializeObjectValues(values map[string]ObjectTestField, object *gtly.Object) error {
	fieldValues := map[string]interface{}{}
	for k, value := range values {
		if value.set == byInitialize {
			fieldValues[k] = value.value
		}
	}
	if len(fieldValues) > 0 {
		return object.Set(fieldValues)
	}
	return nil
}

func setObjectByTypeMutator(v ObjectTestField, object *gtly.Object, mutator *gtly.Mutator) {
	switch fieldValue := v.value.(type) {
	case int:
		mutator.Int(object, fieldValue)
	case float64:
		mutator.Float64(object, fieldValue)
	case string:
		mutator.String(object, fieldValue)
	case bool:
		mutator.Bool(object, fieldValue)
	case time.Time:
		mutator.Time(object, fieldValue)
	default:
		panic(fmt.Errorf("not implemented Mutator for type %T", fieldValue))
	}
}

func getIsoDate(value string) time.Time {
	date, _ := time.Parse("YYYY-MM-DD", value)
	return date
}
