package json

import (
	"github.com/francoispqt/gojay"
	"github.com/viant/gtly"
)

//Collection JSON collection wrapper
type Collection struct {
	Collection gtly.Collection
}

//IsNil return true if collection is empty
func (c Collection) IsNil() bool {
	if c.Collection == nil {
		return true
	}
	return c.Collection.Size() == 0
}

//MarshalJSONArray converts collection  into JSON array
func (c Collection) MarshalJSONArray(enc *gojay.Encoder) {
	if c.Collection == nil {
		return
	}
	c.Collection.Objects(func(item *gtly.Object) (b bool, err error) {
		enc.AddObject(&Object{item})
		return true, nil
	})
}
