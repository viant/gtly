package gtly

//Array represents dynamic object slice
type Array struct {
	_data     []*Object
	_provider *Provider
}

//AddObject add elements to a slice
func (s *Array) AddObject(object *Object) {
	s._data = append(s._data, object)
}

//Size return slice size
func (s *Array) Size() int {
	return len(s._data)
}

//Proto returns slice _proto
func (s *Array) Proto() *Proto {
	return s._provider.Proto
}

//Range calls handler with every slice element
func (m *Array) Range(handler func(item interface{}) (bool, error)) error {
	for _, v := range m._data {
		cont, err := handler(v)

		if !cont || err != nil {
			return err
		}
	}
	return nil
}

//Objects call handler for every object in this collection
func (m *Array) Objects(handler func(item *Object) (bool, error)) error {
	for _, v := range m._data {
		cont, err := handler(v)

		if !cont || err != nil {
			return err
		}
	}
	return nil
}

//AddObject adds object
func (s *Array) Add(value map[string]interface{}) {
	anObject := s._provider.NewObject()
	anObject.Init(value)
	s._data = append(s._data, anObject)
}

//First returns first element on the slice
func (s *Array) First() *Object {
	if s.Size() == 0 {
		return nil
	} else {
		return s._data[0]
	}
}
