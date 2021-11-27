package gtly_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/viant/gtly"
	"reflect"
	"testing"
	"time"
)

type setMethod string

const (
	NoSet              setMethod = "no set"
	ByInterfaceMutator setMethod = "by interface mutator"
	ByTypeMutator      setMethod = "by type mutator"
	ByInitialize       setMethod = "by initialize"
	ByValue            setMethod = "by interface"
)

type ObjTestCase struct {
	description      string
	fields           map[string]ObjectTestField
	isNil            bool
	outputCaseFormat string
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
			outputCaseFormat: gtly.CaseLowerCamel,
			description:      "set by value",
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         ByValue,
					expectedSet: true,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         ByValue,
					expectedSet: true,
					hidden:      true,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         ByValue,
					expectedSet: true,
				},
				"Prop4": {
					value:       true,
					index:       3,
					set:         ByValue,
					expectedSet: true,
				},
				"Prop5": {
					value:       getIsoDate("2021-11-01"),
					index:       4,
					set:         ByValue,
					expectedSet: true,
				},
				"Prop6": {
					value:       "",
					index:       5,
					set:         NoSet,
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
			outputCaseFormat: gtly.CaseLowerCamel,
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         ByInitialize,
					expectedSet: true,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         ByInitialize,
					expectedSet: true,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         ByInitialize,
					expectedSet: true,
				},
				"Prop4": {
					value:       "",
					index:       3,
					set:         ByInitialize,
					expectedSet: true,
				},
				"Prop5": {
					value:       getIsoDate("2021-11-01"),
					index:       4,
					set:         ByInitialize,
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
			outputCaseFormat: gtly.CaseLowerCamel,
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         ByInterfaceMutator,
					expectedSet: true,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         ByInterfaceMutator,
					expectedSet: true,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         ByInterfaceMutator,
					expectedSet: true,
				},
				"Prop4": {
					value:       "",
					index:       3,
					set:         ByInterfaceMutator,
					expectedSet: true,
				},
				"Prop5": {
					value:       getIsoDate("2021-11-01"),
					index:       4,
					set:         ByInterfaceMutator,
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
			outputCaseFormat: gtly.CaseUpperCamel,
			description:      "set by type mutator",
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         ByTypeMutator,
					expectedSet: true,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         ByTypeMutator,
					expectedSet: true,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         ByTypeMutator,
					expectedSet: true,
				},
				"Prop4": {
					value:       "",
					index:       3,
					set:         ByTypeMutator,
					expectedSet: true,
				},
				"Prop5": {
					value:       getIsoDate("2021-11-01"),
					index:       4,
					set:         ByTypeMutator,
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
			outputCaseFormat: gtly.CaseUpper,
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         ByInitialize,
					expectedSet: true,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         ByInitialize,
					expectedSet: true,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         ByInitialize,
					expectedSet: true,
				},
				"Prop4": {
					value:       "",
					index:       3,
					set:         ByInitialize,
					expectedSet: true,
				},
				"Prop5": {
					value:       getIsoDate("2021-11-01"),
					index:       4,
					set:         ByInitialize,
					expectedSet: true,
				},
				"Prop6": {
					value:       "abc",
					index:       5,
					set:         ByTypeMutator,
					expectedSet: true,
				},
				"Prop7": {
					value:       3.5,
					index:       6,
					set:         ByValue,
					expectedSet: true,
				},
				"Prop8": {
					value:       4,
					index:       7,
					set:         ByInterfaceMutator,
					expectedSet: true,
				},
				"Prop9": {
					value:       "",
					index:       8,
					set:         ByTypeMutator,
					expectedSet: true,
				},
				"Prop10": {
					value:       getIsoDate("2021-11-01"),
					index:       9,
					set:         ByValue,
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
			outputCaseFormat: gtly.CaseLowerCamel,
			fields: map[string]ObjectTestField{
				"Prop1": {
					value:       "abc",
					index:       0,
					set:         NoSet,
					expectedSet: false,
				},
				"Prop2": {
					value:       3.5,
					index:       1,
					set:         NoSet,
					expectedSet: false,
				},
				"Prop3": {
					value:       4,
					index:       2,
					set:         NoSet,
					expectedSet: false,
				},
			},
			asMap: map[string]interface{}{},
			isNil: true,
		},
	}

	for i, testCase := range testCases {
		provider := gtly.NewProvider(fmt.Sprintf("testCase%v", i))
		initObjectProviderFields(testCase, provider)
		anObject := provider.NewObject()
		setObjectValues(testCase.fields, anObject)
		runObjectTestFor(t, testCase, anObject)
	}
}

func runObjectTestFor(t *testing.T, testCase ObjTestCase, anObject *gtly.Object) {
	checkObjectValues(t, testCase.fields, anObject, testCase.description)
	checkObjectValuesOutOfRange(t, testCase, anObject, testCase.description)
	assert.Equal(t, testCase.isNil, anObject.IsNil(), testCase.description)
	assert.Equal(t, testCase.asMap, anObject.AsMap(), testCase.description)
}

func initObjectProviderFields(testCase ObjTestCase, provider *gtly.Provider) {
	fields := make([]*gtly.Field, len(testCase.fields))
	// To keep given order after adding Field
	for k, field := range testCase.fields {
		fields[field.index] = &gtly.Field{Name: k, Type: reflect.TypeOf(field.value), Index: field.index}
	}

	for _, field := range fields {
		provider.AddField(field)
		if testCase.fields[field.Name].hidden {
			provider.Hide(field.Name)
		}
	}
	provider.OutputCaseFormat(gtly.CaseUpperCamel, testCase.outputCaseFormat)
}

func checkObjectValues(t *testing.T, values map[string]ObjectTestField, object *gtly.Object, description string) {
	for k, v := range values {
		assert.Equal(t, v.expectedSet, object.SetAt(v.index), description)
		if v.set == NoSet {
			continue
		}
		checkObjectByType(t, v.value, k, object, v.index, description)
		assert.Equal(t, v.value, object.Value(k), description)
		assert.Equal(t, v.value, object.Accessor(k)(), description)
		fieldValue, isValuePresent := object.ValueAt(v.index)
		assert.Equal(t, v.value, fieldValue, description)
		assert.True(t, isValuePresent, description)
	}
}

func checkObjectByType(t *testing.T, value interface{}, name string, object *gtly.Object, index int, description string) {
	switch fieldValue := value.(type) {
	case int:
		assert.Equal(t, value, object.IntAccessor(name)(), description)
		valueAt, isPresent := object.IntAt(index)
		assertObjectValueFound(t, fieldValue, valueAt, isPresent, description)
	case string:
		assert.Equal(t, value, object.StringAccessor(name)(), description)
		valueAt, isPresent := object.StringAt(index)
		assertObjectValueFound(t, value, valueAt, isPresent, description)
	case bool:
		assert.Equal(t, value, object.BoolAccessor(name)(), description)
		valueAt, isPresent := object.BoolAt(index)
		assertObjectValueFound(t, fieldValue, valueAt, isPresent, description)
	case float64:
		assert.Equal(t, value, object.FloatAccessor(name)(), description)
		valueAt, isPresent := object.FloatAt(index)
		assertObjectValueFound(t, fieldValue, valueAt, isPresent, description)
	case time.Time:
		assert.Equal(t, value, object.TimeAccessor(name)(), description)
		valueAt, isPresent := object.TimeAt(index)
		assertObjectValueFound(t, fieldValue, valueAt, isPresent, description)
	default:
		panic(fmt.Errorf("not implemented Accessor for type %T", fieldValue))
	}
}

func assertObjectValueFound(t *testing.T, expected interface{}, actual interface{}, isPresent bool, description string) {
	assert.Equal(t, expected, actual, description)
	assert.True(t, isPresent, description)
}

func setObjectValues(values map[string]ObjectTestField, object *gtly.Object) {
	initializeObjectValues(values, object)
	for k, v := range values {
		switch v.set {
		case NoSet, ByInitialize:
			continue
		case ByInterfaceMutator:
			object.Mutator(k)(v.value)
		case ByTypeMutator:
			setObjectByTypeMutator(v, object, k)
		case ByValue:
			object.SetValue(k, v.value)
		default:
			panic(fmt.Errorf("not implemented setting method: %v", v.set))
		}
	}
}

func initializeObjectValues(values map[string]ObjectTestField, object *gtly.Object) {
	fieldValues := map[string]interface{}{}
	for k, value := range values {
		if value.set == ByInitialize {
			fieldValues[k] = value.value
		}
	}
	if len(fieldValues) > 0 {
		object.Init(fieldValues)
	}
}

func setObjectByTypeMutator(v ObjectTestField, object *gtly.Object, k string) {
	switch fieldValue := v.value.(type) {
	case int:
		object.IntMutator(k)(fieldValue)
	case float64:
		object.FloatMutator(k)(fieldValue)
	case string:
		object.Mutator(k)(fieldValue)
	case bool:
		object.BoolMutator(k)(fieldValue)
	case time.Time:
		object.TimeMutator(k)(fieldValue)
	default:
		panic(fmt.Errorf("not implemented Mutator for type %T", fieldValue))
	}
}

func checkObjectValuesOutOfRange(t *testing.T, testCase ObjTestCase, anObject *gtly.Object, description string) {
	fieldsLen := len(testCase.fields)
	assert.False(t, anObject.SetAt(fieldsLen), description)
	fieldValue, isValuePresent := anObject.ValueAt(fieldsLen)
	assertObjectValueNotFound(t, fieldValue, nil, isValuePresent, description)
	intTypeFieldValue, isValuePresent := anObject.IntAt(fieldsLen)
	assertObjectValueNotFound(t, intTypeFieldValue, 0, isValuePresent, description)
	stringTypeFieldValue, isValuePresent := anObject.StringAt(fieldsLen)
	assertObjectValueNotFound(t, stringTypeFieldValue, "", isValuePresent, description)
	boolTypeFieldValue, isValuePresent := anObject.BoolAt(fieldsLen)
	assertObjectValueNotFound(t, boolTypeFieldValue, false, isValuePresent, description)
	floatValue, isValuePresent := anObject.FloatAt(fieldsLen)
	assertObjectValueNotFound(t, floatValue, 0.0, isValuePresent, description)
	timeValue, isValuePresent := anObject.TimeAt(fieldsLen)
	assertObjectValueNotFound(t, timeValue, time.Time{}, isValuePresent, description)
}

func assertObjectValueNotFound(t *testing.T, fieldValue interface{}, nilValue interface{}, isValuePresent bool, description string) {
	assert.Equal(t, nilValue, fieldValue, description)
	assert.False(t, isValuePresent, description)
}

func getIsoDate(value string) time.Time {
	date, _ := time.Parse("YYYY-MM-DD", value)
	return date
}

// Benchmarks
var objBenchProvider *gtly.Provider
var objObjectValues map[string]interface{}
var objPreparedObj *gtly.Object

func init() {
	objBenchProvider = gtly.NewProvider("Bench")
	objBenchProvider.AddField(&gtly.Field{Name: "Prop1", Type: reflect.TypeOf(0), Index: 0})
	objBenchProvider.AddField(&gtly.Field{Name: "Prop2", Type: reflect.TypeOf(""), Index: 1})
	objBenchProvider.AddField(&gtly.Field{Name: "Prop3", Type: reflect.TypeOf(0.0), Index: 2})
	objObjectValues = map[string]interface{}{
		"Prop1": 123,
		"Prop2": "abc",
		"Prop3": 1.0,
	}
	objPreparedObj = objBenchProvider.NewObject()
	objPreparedObj.Init(objObjectValues)
}

func BenchmarkObject_IntAccessor(b *testing.B) {
	b.ReportAllocs()
	o := objBenchProvider.NewObject()
	prop1Accessor := o.IntAccessor("Prop1")
	prop1Mutator := o.IntMutator("Prop1")
	for i := 0; i < b.N; i++ {
		prop1Mutator(i)
		if prop1Accessor() != i {
			b.Fail()
		}
	}
}

func BenchmarkObject_Initialize(b *testing.B) {
	b.ReportAllocs()
	o := objBenchProvider.NewObject()
	for i := 0; i < b.N; i++ {
		o.Init(objObjectValues)
	}
}

func BenchmarkObject_AsMap(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		objPreparedObj.AsMap()
	}
}

func BenchmarkObject_ValueAt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		objPreparedObj.ValueAt(0)
	}
}

func BenchmarkObject_Value(b *testing.B) {
	for i := 0; i < b.N; i++ {
		objPreparedObj.Value("Prop1")
	}
}

func BenchmarkObject_StringAt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		objPreparedObj.StringAt(1)
	}
}

func BenchmarkObject_FloatAt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		objPreparedObj.FloatAt(2)
	}
}
