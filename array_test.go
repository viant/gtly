package gtly_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/viant/gtly"
	"github.com/viant/toolbox/format"
	"math"
	"testing"
)

type ArrayFieldValue struct {
	values map[string]interface{}
	asMap  map[string]interface{}
	add    AddMethod
}

type ArrayTestCase struct {
	description  string
	outputFormat format.Case
	fields       []*gtly.Field
	values       []ArrayFieldValue
}

func TestArray(t *testing.T) {
	testCases := []ArrayTestCase{
		{
			description:  "as object",
			outputFormat: format.CaseLowerCamel,

			fields: []*gtly.Field{
				{
					Name:     "Prop1",
					DataType: gtly.FieldTypeString,
				},
				{
					Name:     "Prop2",
					DataType: gtly.FieldTypeInt,
				},
				{
					Name:     "Prop3",
					DataType: gtly.FieldTypeFloat64,
				},
			},
			values: []ArrayFieldValue{
				{
					values: map[string]interface{}{
						"Prop1": "abc",
						"Prop2": 1,
						"Prop3": 4.0,
					},
					asMap: map[string]interface{}{
						"prop1": "abc",
						"prop2": 1,
						"prop3": 4.0,
					},
					add: AsObject,
				},
				{
					values: map[string]interface{}{
						"Prop1": "cdef",
						"Prop2": 2,
						"Prop3": 8.0,
					},
					add: AsObject,
					asMap: map[string]interface{}{
						"prop1": "cdef",
						"prop2": 2,
						"prop3": 8.0,
					},
				},
			},
		},
		{
			description:  "as map",
			outputFormat: format.CaseLowerCamel,
			fields: []*gtly.Field{
				{
					Name:     "Prop1",
					DataType: gtly.FieldTypeString,
				},
				{
					Name:     "Prop2",
					DataType: gtly.FieldTypeInt,
				},
				{
					Name:     "Prop3",
					DataType: gtly.FieldTypeFloat64,
				},
			},
			values: []ArrayFieldValue{
				{
					values: map[string]interface{}{
						"Prop1": "abc",
						"Prop2": 1,
						"Prop3": 4.0,
					},
					add: AsMap,
					asMap: map[string]interface{}{
						"prop1": "abc",
						"prop2": 1,
						"prop3": 4.0,
					},
				},
				{
					values: map[string]interface{}{
						"Prop1": "cdef",
						"Prop2": 2,
						"Prop3": 8.0,
					},
					add: AsMap,
					asMap: map[string]interface{}{
						"prop1": "cdef",
						"prop2": 2,
						"prop3": 8.0,
					},
				},
			},
		},
	}

	for i, testCase := range testCases {

		provider, err := gtly.NewProvider(fmt.Sprintf("test case #%v", i), testCase.fields...)
		if !assert.Nil(t, err, testCase.description) {
			continue
		}
		err = provider.OutputCaseFormat(format.CaseUpperCamel, testCase.outputFormat)
		if !assert.Nil(t, err, testCase.description) {
			continue
		}
		slice := provider.NewArray()
		addToSlice(testCase, provider, slice)
		assert.Equal(t, slice.Size(), len(testCase.values), testCase.description)
		assert.Equal(t, provider.Proto, slice.Proto(), testCase.description)
		checkSliceObjects(t, testCase, slice, len(testCase.values))
		checkSliceObjects(t, testCase, slice, int(math.Ceil(float64(len(testCase.values))/2)))
		checkSliceRange(t, testCase, slice, len(testCase.values))
		checkSliceFirstElement(t, testCase, slice)
	}
}

func checkSliceRange(t *testing.T, testCase ArrayTestCase, slice *gtly.Array, n int) {
	counter := 0
	err := slice.Range(func(item interface{}) (bool, error) {
		assert.Equal(t, testCase.values[counter].asMap, item.(*gtly.Object).AsMap())
		counter++
		return counter != n, nil
	})
	assert.Equal(t, n, counter, testCase.description)
	assert.Nil(t, err, testCase.description)
}

func checkSliceObjects(t *testing.T, testCase ArrayTestCase, slice *gtly.Array, n int) {
	counter := 0
	err := slice.Objects(func(item *gtly.Object) (bool, error) {
		assert.Equal(t, testCase.values[counter].asMap, item.AsMap())
		counter++
		return counter != n, nil
	})
	assert.Equal(t, n, counter, testCase.description)
	assert.Nil(t, err, testCase.description)
}

func checkSliceFirstElement(t *testing.T, testCase ArrayTestCase, slice *gtly.Array) {
	if len(testCase.values) == 0 {
		assert.Equal(t, nil, slice.First(), testCase.description)
	} else {
		assert.Equal(t, testCase.values[0].asMap, slice.First().AsMap(), testCase.description)
	}
}

func addToSlice(testCase ArrayTestCase, provider *gtly.Provider, slice *gtly.Array) {
	for _, obj := range testCase.values {
		switch obj.add {
		case AsObject:
			anObject := provider.NewObject()
			initObjectValues(obj.values, anObject)
			slice.AddObject(anObject)
		case AsMap:
			slice.Add(obj.values)
		default:
			panic(fmt.Errorf("not implemented add method: %v", obj.add))
		}
	}
}

func initObjectValues(values map[string]interface{}, object *gtly.Object) {
	for k, v := range values {
		object.SetValue(k, v)
	}
}
