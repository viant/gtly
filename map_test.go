package gtly_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/viant/gtly"
	"github.com/viant/toolbox/format"
	"reflect"
	"testing"
)

type MapTestCase struct {
	description    string
	fields         map[string]interface{}
	uniqueField    string
	mapUniqueField string
	values         []MapFieldValues
	outputFormat   format.Case
}

type MapFieldValues struct {
	id     int
	values map[string]interface{}
	asMap  map[string]interface{}
	add    AddMethod
	key    string
}

func TestMap(t *testing.T) {
	testCases := []MapTestCase{
		{
			description:    "as object",
			outputFormat:   format.CaseLowerCamel,
			uniqueField:    "Id",
			mapUniqueField: "id",
			fields: map[string]interface{}{
				"Id":    0,
				"Prop1": "",
				"Prop2": 0,
				"Prop3": 0.0,
			},
			values: []MapFieldValues{
				{
					key: "obj-1",
					values: map[string]interface{}{
						"Id":    1,
						"Prop1": "abc",
						"Prop2": 10,
						"Prop3": 10.5,
					},
					add: AsObject,
					asMap: map[string]interface{}{

						"id":    1,
						"prop1": "abc",
						"prop2": 10,
						"prop3": 10.5,
					},
				},
			},
		},
	}

	for i, testCase := range testCases {
		provider := gtly.NewProvider(fmt.Sprintf("testCase#%v", i))
		initMapProvider(testCase, provider)
		aMap := provider.NewMap(gtly.NewKeyProvider(testCase.uniqueField))
		addMapObjects(testCase, provider, aMap)
		checkMapObjects(t, testCase, aMap, len(testCase.values))
		assert.Equal(t, len(testCase.values), aMap.Size())
		assert.Equal(t, provider.Proto, aMap.Proto(), testCase.description)
		checkMapFirstElement(t, testCase, aMap)
		checkMapRange(t, testCase, aMap, len(testCase.values))
		assert.Equal(t, testCase.values[0].asMap, aMap.Object(testCase.values[0].key).AsMap())
		assert.Nil(t, aMap.Object(""))
	}
}

func checkMapRange(t *testing.T, testCase MapTestCase, aMap *gtly.Map, n int) {
	counter := 0
	aMap.Range(func(item interface{}) (bool, error) {
		found := false
		for _, value := range testCase.values {
			if value.values[testCase.uniqueField] == item.(*gtly.Object).Value(testCase.uniqueField) {
				found = true
				break
			}
		}

		assert.True(t, found, testCase.description)
		counter++
		return counter < n, nil
	})
}

func checkMapObjects(t *testing.T, testCase MapTestCase, aMap *gtly.Map, n int) {
	counter := 0
	aMap.Objects(func(item *gtly.Object) (bool, error) {
		found := false
		for _, value := range testCase.values {
			if value.values[testCase.uniqueField] == item.Value(testCase.uniqueField) {
				found = true
				break
			}
		}

		assert.True(t, found)
		counter++
		return counter < n, nil
	})
}

func initMapProvider(testCase MapTestCase, provider *gtly.Provider) {
	for k, v := range testCase.fields {
		provider.AddField(&gtly.Field{
			Type: reflect.TypeOf(v),
			Name: k,
		})
	}
	provider.OutputCaseFormat(format.CaseUpperCamel, testCase.outputFormat)
}

func checkMapFirstElement(t *testing.T, testCase MapTestCase, aMap *gtly.Map) {
	if len(testCase.values) == 0 {
		assert.Nil(t, aMap.First(), testCase.description)
	} else {
		assert.NotNilf(t, aMap.First().AsMap(), testCase.description)
	}
}

func addMapObjects(testCase MapTestCase, provider *gtly.Provider, aMap *gtly.Map) {
	for _, value := range testCase.values {
		anObject := provider.NewObject()
		anObject.Init(value.values)
		aMap.PutObject(value.key, anObject)
	}
}
