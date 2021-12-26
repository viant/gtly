package json

import (
	"encoding/json"
	"github.com/francoispqt/gojay"
	"github.com/viant/assertly"
	"github.com/viant/gtly"
	"github.com/viant/toolbox/format"
	"reflect"
	"testing"
)

func TestObject_MarshalJSONObject(t *testing.T) {
	testCases := []struct {
		description      string
		fields           map[string]interface{}
		values           map[string]interface{}
		result           map[string]interface{}
		outputFormat     format.Case
		sourceCaseFormat format.Case
	}{
		{
			description: "marshal all fields",
			fields: map[string]interface{}{
				"Id":    0,
				"Name":  "",
				"Price": 0.5,
			},
			values: map[string]interface{}{
				"Id":    1,
				"Name":  "Foo",
				"Price": 10.5,
			},
			result: map[string]interface{}{
				"id":    1,
				"name":  "Foo",
				"price": 10.5,
			},
			outputFormat:     format.CaseLowerCamel,
			sourceCaseFormat: format.CaseUpper,
		},
		{
			description: "marshal fields which were set",
			fields: map[string]interface{}{
				"Id":    0,
				"Name":  "",
				"Price": 0.5,
			},
			values: map[string]interface{}{
				"Id":   1,
				"Name": "Foo",
			},
			result: map[string]interface{}{
				"id":   1,
				"name": "Foo",
			},
			outputFormat:     format.CaseLowerCamel,
			sourceCaseFormat: format.CaseUpper,
		},
	}

	for _, testCase := range testCases {
		provider := gtly.NewProvider(testCase.description)
		for k, v := range testCase.fields {
			provider.AddField(&gtly.Field{
				Name: k,
				Type: reflect.TypeOf(v),
			})
		}
		provider.OutputCaseFormat(testCase.sourceCaseFormat, testCase.outputFormat)
		anObject := provider.NewObject()
		for k, v := range testCase.values {
			anObject.SetValue(k, v)
		}
		object := Object{anObject}
		val, _ := gojay.MarshalJSONObject(object)
		newMap := new(map[string]interface{})
		json.Unmarshal(val, newMap)
		assertly.AssertValues(t, testCase.result, *newMap, testCase.description)
	}
}

// Benchmarks
var anObject *Object

func init() {
	provider := gtly.NewProvider("")
	provider.AddField(&gtly.Field{
		Name: "Id",
		Type: reflect.TypeOf(1),
	})
	provider.AddField(&gtly.Field{
		Name: "Name",
		Type: reflect.TypeOf(""),
	})
	provider.AddField(&gtly.Field{
		Name: "Price",
		Type: reflect.TypeOf(1.5),
	})
	anObject = &Object{provider.NewObject()}
	anObject.SetValue("Id", 10)
	anObject.SetValue("Name", "some name")
	anObject.SetValue("Price", 100.5)
}

func BenchmarkObject_MarshalJSONObject(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		gojay.MarshalJSONObject(anObject)
	}
}
