package json

import (
	"github.com/francoispqt/gojay"
	"github.com/pkg/errors"
	"github.com/viant/gtly"
	"github.com/viant/toolbox"
)

//Marshal convert type to JSON
func Marshal(v interface{}) ([]byte, error) {
	switch raw := v.(type) {
	case *gtly.Object:
		return gojay.Marshal(&Object{raw})
	case *gtly.Map:
		return gojay.Marshal(&Collection{Collection: raw})
	case *gtly.Multimap:
		return gojay.Marshal(&Collection{Collection: raw})
	case *gtly.Array:
		return gojay.Marshal(&Collection{Collection: raw})
	case []interface{}:
		return gojay.Marshal(&Slice{raw})
	default:
		if toolbox.IsSlice(raw) {
			return gojay.Marshal(&Slice{raw})
		}
		return nil, errors.Errorf("unsupported type: %T", v)
	}
}
