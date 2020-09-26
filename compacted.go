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
