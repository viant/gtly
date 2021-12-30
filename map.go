package gtly

//Map represents generic map
type Map struct {
	_provider   *Provider
	_map        map[interface{}]*Object
	keyProvider KeyProvider
}

//Proto returns map proto
func (m *Map) Proto() *Proto {
	return m._provider.Proto
}

//Range calls handler with every slice element
func (m *Map) Range(handler func(item interface{}) (bool, error)) error {
	for _, v := range m._map {
		cont, err := handler(v)

		if !cont || err != nil {
			return err
		}
	}
	return nil
}

//Objects call handler for every object in this collection
func (m *Map) Objects(handler func(item *Object) (bool, error)) error {
	for _, v := range m._map {
		cont, err := handler(v)

		if !cont || err != nil {
			return err
		}
	}
	return nil
}

//PutObject add object to the map
func (m *Map) PutObject(key interface{}, object *Object) {
	m._map[key] = object
}

//AddObject adds object
func (m *Map) AddObject(obj *Object) {
	m._map[m.keyProvider(obj)] = obj
}

//Add adds object from a map
func (m *Map) Add(values map[string]interface{}) error {
	item := m._provider.NewObject()
	err := item.Set(values)
	if err != nil {
		return err
	}
	m._map[m.keyProvider(item)] = item
	return nil
}

//First return the first map elements
func (m *Map) First() *Object {
	if m.Size() == 0 {
		return nil
	}

	for _, v := range m._map {
		return v
	}
	return nil
}

//Size return map size
func (m *Map) Size() int {
	return len(m._map)
}

//Object returns an object for specified key or nil
func (m *Map) Object(key interface{}) *Object {
	data, ok := m._map[key]
	if !ok {
		return nil
	}
	return data
}

//Pairs iterate over object slice, any update to objects are applied to the slice
func (m *Map) Pairs(handler func(key interface{}, item *Object) (bool, error)) error {
	for key, item := range m._map {
		next, err := handler(key, item)
		if !next || err != nil {
			return err
		}
	}
	return nil
}
