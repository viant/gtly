package gtly

//Compacted represents compacted collection
type Compacted struct {
	Fields []*Field
	Data   [][]interface{}
}

//Update updates collection
func (c Compacted) Update(collection Collection) error {
	proto := collection.Proto()
	for i := range c.Fields {
		proto.AddField(c.Fields[i])
	}
	for i := range c.Data {
		object, err := proto.Object(c.Data[i])
		if err != nil {
			return err
		}
		collection.AddObject(object)

	}
	return nil
}


//TransformBinary transform binary data into string
func (c *Compacted) TransformBinary() {
	for i := range c.Data {
		for j, item := range c.Data[i] {
			switch val := item.(type) {
			case []byte:
				if len(val) == 0 {
					c.Data[i][j] = nil
					continue
				}
				c.Data[i][j] = string(val)
			}
		}
	}
}
