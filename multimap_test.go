package gtly_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/gtly"
	"testing"
)

type MultimapTestCase struct {
	description string
	fields      []*gtly.Field
	multiValues []map[string]interface{}
	uniqueField string
	size        int
	slicesLen   map[interface{}]int
	isNil       bool
}

func TestMultimap(t *testing.T) {
	testCases := []MultimapTestCase{
		{
			description: "single values",
			uniqueField: "Prop1",
			size:        2,
			isNil:       false,
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
			multiValues: []map[string]interface{}{
				{
					"Prop1": "abc",
					"Prop2": 1,
					"Prop3": 2.0,
				},
				{
					"Prop1": "abcd",
					"Prop2": 1,
					"Prop3": 2.0,
				},
			},
			slicesLen: map[interface{}]int{
				"abc":  1,
				"abcd": 1,
			},
		},
	}

	for _, testCase := range testCases {
		provider, err := gtly.NewProvider("", testCase.fields...)
		if !assert.Nil(t, err, testCase.description) {
			continue
		}
		multiMap := provider.NewMultimap(gtly.NewKeyProvider(testCase.uniqueField))
		initMultimapValues(testCase, provider, multiMap)
		assert.Equal(t, testCase.size, multiMap.Size(), testCase.description)
		assert.Equal(t, testCase.isNil, multiMap.IsNil())
		assert.True(t, testCase.slicesLen[multiMap.First().Value(testCase.uniqueField)] != 0, testCase.description)
		checkMultimapRange(t, testCase, multiMap)
	}
}

func checkMultimapRange(t *testing.T, testCase MultimapTestCase, multiMap *gtly.Multimap) {
	counter := 0
	multiMap.Range(func(item interface{}) (bool, error) {
		counter++
		return true, nil
	})
	assert.Equal(t, counter, len(testCase.multiValues), testCase.description)
}

func initMultimapValues(testCase MultimapTestCase, provider *gtly.Provider, multiMap *gtly.Multimap) {
	for _, values := range testCase.multiValues {
		anObject := provider.NewObject()
		anObject.Set(values)
		multiMap.AddObject(anObject)
	}
}
