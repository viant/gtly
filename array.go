package gtly

//Array represents dynamic object slice
type Array struct {
	_data     []*Object
	_provider *Provider
}

//AddObject add elements to a slice
func (a *Array) AddObject(object *Object) {
	a._data = append(a._data, object)
}

//Size return slice size
func (a *Array) Size() int {
	return len(a._data)
}

//Proto returns slice proto
func (a *Array) Proto() *Proto {
	return a._provider.Proto
}

//Range calls handler with every slice element
func (a *Array) Range(handler func(item interface{}) (bool, error)) error {
	for _, v := range a._data {
		cont, err := handler(v)

		if !cont || err != nil {
			return err
		}
	}
	return nil
}

//Objects call handler for every object in this collection
func (a *Array) Objects(handler func(item *Object) (bool, error)) error {
	for _, v := range a._data {
		cont, err := handler(v)

		if !cont || err != nil {
			return err
		}
	}
	return nil
}

//Add adds object
func (a *Array) Add(value map[string]interface{}) error {
	item := a._provider.NewObject()
	err := item.Set(value)
	if err != nil {
		return err
	}
	a._data = append(a._data, item)
	return nil
}

//First returns the first element on the slice
func (a *Array) First() *Object {
	if a.Size() == 0 {
		return nil
	}
	return a._data[0]
}
