package gtly

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyProducer_Produce(t *testing.T) {
	testCases := []struct {
		description string
		fields      []*Field
		values      []map[string]interface{}
		results     []string
	}{
		{
			description: "",
			fields: []*Field{
				{
					Name:     "Id",
					DataType: FieldTypeString,
				},
				{
					Name:     "Name",
					DataType: FieldTypeString,
				},
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
		provider, err := NewProvider(fmt.Sprintf("test#%v", i), testCase.fields...)
		if !assert.Nil(t, err, testCase.description) {
			continue
		}
		for index, value := range testCase.values {
			anObject := provider.NewObject()
			err = anObject.Set(value)
			assert.Nil(t, err, testCase.description)
			keyProducer := NewKeyProvider("Id")
			assert.Equal(t, testCase.results[index], keyProducer(anObject), testCase.description)
		}
	}
}

// Benchmarks
var anObject *Object
var keyProvider func(o *Object) interface{}

func init() {
	provider, err := NewProvider("", NewField("Id", FieldTypeInt), NewField("Name", FieldTypeString))
	if err != nil {
		panic(err)
	}
	anObject = provider.NewObject()
	anObject.Set([]interface{}{123, "John"})
	keyProvider = NewKeyProvider("Id")
}

func BenchmarkKeyProvider(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keyProvider(anObject)
	}
}
