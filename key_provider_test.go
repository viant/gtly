package gtly

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestKeyProducer_Produce(t *testing.T) {
	testCases := []struct {
		description string
		fields      map[string]interface{}
		values      []map[string]interface{}
		results     []string
	}{
		{
			description: "",
			fields: map[string]interface{}{
				"Id":   "",
				"Name": "",
			},
			values: []map[string]interface{}{
				{
					"Id":   "123-123",
					"Name": "John",
				},
				{
					"Id":   "321-321",
					"Name": "John",
				},
			},
			results: []string{"123-123", "321-321"},
		},
	}

	for i, testCase := range testCases {
		provider := NewProvider(fmt.Sprintf("test#%v", i))
		initKeyProducerProvider(testCase, provider)
		for index, value := range testCase.values {
			anObject := provider.NewObject()
			anObject.Init(value)
			keyProducer := NewKeyProvider("Id")
			assert.Equal(t, testCase.results[index], keyProducer(anObject), testCase.description)
		}
	}
}

func initKeyProducerProvider(testCase struct {
	description string
	fields      map[string]interface{}
	values      []map[string]interface{}
	results     []string
}, provider *Provider) {
	for key, value := range testCase.fields {
		provider.AddField(&Field{
			Type: reflect.TypeOf(value),
			Name: key,
		})
	}
}

// Benchmarks
var anObject *Object
var keyProvider func(o *Object) interface{}

func init() {
	provider := NewProvider("")
	provider.AddField(&Field{
		Type: reflect.TypeOf(interface{}("")),
		Name: "Id",
	})
	provider.AddField(&Field{
		Type: reflect.TypeOf(interface{}("")),
		Name: "Name",
	})
	anObject = provider.NewObject()
	anObject.Init(map[string]interface{}{
		"Id":   "123",
		"Name": "John",
	})
	keyProvider = NewKeyProvider("Id")
}

func BenchmarkKeyProvider(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keyProvider(anObject)
	}
}
