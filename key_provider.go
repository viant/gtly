package gtly

//KeyProvider represents a key provider
type KeyProvider func(o *Object) interface{}

//NewKeyProvider creates a key provider
func NewKeyProvider(fieldName string) KeyProvider {
	uniqueKeyIndex := 0
	wasKeyProduced := false
	return func(o *Object) interface{} {
		if wasKeyProduced {
			value, _ := o.ValueAt(uniqueKeyIndex)
			return value
		}

		field := o.Field(fieldName)
		uniqueKeyIndex = field.Index
		wasKeyProduced = true
		return o.Value(fieldName)
	}
}
