package gtly

//Multimap represents generic multi map
type Multimap struct {
	_provider   *Provider
	_map        map[interface{}][]*Object
	keyProvider KeyProvider
}

//Proto returns multimap _proto
func (m *Multimap) Proto() *Proto {
	return m._provider.Proto
}

//Range calls handler with every slice element
func (m *Multimap) Range(handler func(item interface{}) (bool, error)) error {
	cont := true
	var err error
	for _, values := range m._map {
		for _, value := range values {
			cont, err = handler(value)
			if !cont || err != nil {
				return err
			}
		}
		if !cont || err != nil {
			return err
		}
	}
	return nil
}

//Objects call handler for every object in this collection
func (m *Multimap) Objects(handler func(item *Object) (bool, error)) error {
	cont := true
	var err error
	for _, values := range m._map {
		for _, value := range values {
			cont, err = handler(value)
			if !cont || err != nil {
				return err
			}
		}
		if !cont || err != nil {
			return err
		}
	}
	return nil
}

//First returns an element from multimap
func (m *Multimap) First() *Object {
	if m.Size() == 0 {
		return nil
	}

	for _, values := range m._map {
		for _, v := range values {
			return v
		}
	}
	return nil
}

//Add add item to a map
func (m *Multimap) Add(values map[string]interface{}) {
	object := m._provider.NewObject()
	object.Init(values)
	key := m.keyProvider(object)
	if _, ok := m._map[key]; !ok {
		m._map[key] = make([]*Object, 0)
	}
	m._map[key] = append(m._map[key], object)
}

//AddObject add object into multimap
func (m *Multimap) AddObject(o *Object) {
	key := m.keyProvider(o)
	if _, ok := m._map[key]; !ok {
		m._map[key] = make([]*Object, 0)
	}
	m._map[key] = append(m._map[key], o)
}

//Slices iterate over object slice, any update to objects are applied to the slice
func (m *Multimap) Slices(handler func(key interface{}, value *Array) (bool, error)) error {
	aMap := m._map
	for key, item := range aMap {
		slice := &Array{_provider: m._provider, _data: item}
		next, err := handler(key, slice)
		aMap[key] = slice._data
		if !next || err != nil {
			return err
		}
	}
	return nil
}

//Slice returns a slice for specified key or nil
func (m *Multimap) Slice(key string) *Array {
	data, ok := m._map[key]
	if !ok {
		return nil
	}
	return &Array{_provider: m._provider, _data: data}
}

//Size return slice size
func (m *Multimap) Size() int {
	return len(m._map)
}

//IsNil returns true if it's nil
func (m *Multimap) IsNil() bool {
	return len(m._map) == 0
}
