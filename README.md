# gtly - Dynamic data structure with go lang.

[![GoReportCard](https://goreportcard.com/badge/github.com/viant/gtly)](https://goreportcard.com/report/github.com/viant/gtly)
[![GoDoc](https://godoc.org/github.com/viant/gtly?status.svg)](https://godoc.org/github.com/viant/gtly)


This library is compatible with Go 1.15+

Please refer to [`CHANGELOG.md`](CHANGELOG.md) if you encounter breaking changes.

- [Motivation](#motivation)
- [Introduction](#introduction)
- [Usage](#usage)
   - [Object](#object)
   - [Array](#array)
   - [Map](#map)
   - [Multimap](#multimap)
- [Configuration Rule](#configuration-rule)
- [License](#license)


## Motivation

The goal of this project is to use dynamic data types without defining native GO structs with
minimum memory footprint. Alternative would be just using map, but that is way too inefficient.
Having dynamic type safe objects enables building generic solution for REST/Micro Service/ETL etc.
To build dynamic solution dynamic object should be easily transferable into common format like JSON, AVRO, ProtoBuf, etc....

## Introduction

Gtly complex data type use runtime struct based storage with [Proto](proto.go) reference to reduce memory footprint and to avoid reflection.
An proto instance is shared across all [Object](object.go) and [Collection](collection.go) of the same type.
Proto control mapping between field and field position withing object, or slice item.
What's more proto field define field meta data like DateLayout, OutputName controlled by proto CaseFormat dynamically. 
  

## Usage

#### Object

```go
package mypacakge

import (
	"fmt"
	"github.com/viant/gtly"
    "github.com/viant/toolbox/format"
	"github.com/viant/gtly/codec/json"
	"log"
	"time"
)

func NewObject_Usage() {
  fooProvider, err := gtly.NewProvider("foo",
    gtly.NewField("id", gtly.FieldTypeInt),
    gtly.NewField("firsName", gtly.FieldTypeString),
    gtly.NewField("description", gtly.FieldTypeString, gtly.OmitEmptyOpt(true)),
    gtly.NewField("updated", gtly.FieldTypeTime, gtly.DateLayoutOpt("2006-01-02T15:04:05Z07:00")),
    gtly.NewField("numbers", gtly.FieldTypeArray, gtly.ComponentTypeOpt(gtly.FieldTypeInt)),
  )
  if err != nil {
    log.Fatal(err)
  }
  foo1 := fooProvider.NewObject()
  foo1.SetValue("id", 1)
  foo1.SetValue("firsName", "Adam")
  foo1.SetValue("updated", time.Now())
  foo1.SetValue("numbers", []int{1, 2, 3})

  JSON, err := json.Marshal(foo1)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", JSON)
  fooProvider.OutputCaseFormat(format.CaseLower, format.CaseUpperUnderscore)
  JSON, _ = json.Marshal(foo1)
  fmt.Printf("%s\n", JSON)

  fooProvider.OutputCaseFormat(format.CaseLower, format.CaseLowerUnderscore)
  foo1.SetValue("active", true)
  foo1.SetValue("description", "some description")

  JSON, _ = json.Marshal(foo1)
  fmt.Printf("%s\n", JSON)
}
```

#### Array

```go
package mypacakge

import (
	"fmt"
	"github.com/viant/gtly"
	"github.com/viant/gtly/codec/json"
	"log"
	"time"
)

func NewArray_Usage() {
  fooProvider, err := gtly.NewProvider("foo",
    gtly.NewField("id", gtly.FieldTypeInt),
    gtly.NewField("firsName", gtly.FieldTypeString),
    gtly.NewField("income", gtly.FieldTypeFloat64),
    gtly.NewField("description", gtly.FieldTypeString, gtly.OmitEmptyOpt(true)),
    gtly.NewField("updated", gtly.FieldTypeTime, gtly.DateLayoutOpt(time.RFC3339)),
    gtly.NewField("numbers", gtly.FieldTypeArray, gtly.ComponentTypeOpt(gtly.FieldTypeInt)),
  )
  if err != nil {
    log.Fatal(err)
  }
  fooArray1 := fooProvider.NewArray()

  setId := fooProvider.Mutator("id")
  setIncome := fooProvider.Mutator("income")
  setFirstName := fooProvider.Mutator("firsName")

  for i := 0; i < 10; i++ {
    foo1 := fooProvider.NewObject()
    setId.Int(foo1, 1)
    setIncome.Float64(foo1, 64000.0*float64(1+(10/(i+1))))
    setFirstName.String(foo1, "Adam")
    fooArray1.AddObject(foo1)
  }

  now := time.Now()
  fooArray1.Add(map[string]interface{}{
    "id":       100,
    "firsName": "Tom",
    "updated":  now,
  })

  totalIncome := 0.0
  incomeField := fooProvider.Proto.Accessor("income")
  //Iterating collection

  err = fooArray1.Objects(func(object *gtly.Object) (bool, error) {
    fmt.Printf("id: %v\n", object.Value("id"))
    fmt.Printf("name: %v\n", object.Value("name"))
    value := incomeField.Float64(object)
    totalIncome += value
    return true, nil
  })
  fmt.Printf("income total: %v\n", totalIncome)
  if err != nil {
    log.Fatal(err)
  }
  JSON, err := json.Marshal(fooArray1)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s", JSON)
}

```

#### Map

```go
package mypacakge

import (
	"fmt"
	"github.com/viant/gtly"
	"github.com/viant/gtly/codec/json"
	"log"
	"time"
)

func NewMap_Usage() {
  fooProvider, err := gtly.NewProvider("foo",
    gtly.NewField("id", gtly.FieldTypeInt),
    gtly.NewField("firsName", gtly.FieldTypeString),
    gtly.NewField("description", gtly.FieldTypeString, gtly.OmitEmptyOpt(true)),
    gtly.NewField("updated", gtly.FieldTypeTime, gtly.DateLayoutOpt(time.RFC3339)),
    gtly.NewField("numbers", gtly.FieldTypeArray, gtly.ComponentTypeOpt(gtly.FieldTypeInt)),
  )
  if err != nil {
    log.Fatal(err)
  }

  //Creates a map keyed by id Field
  aMap := fooProvider.NewMap(gtly.NewKeyProvider("id"))
  for i := 0; i < 10; i++ {
    foo := fooProvider.NewObject()
    foo.SetValue("id", i)
    foo.SetValue("firsName", fmt.Sprintf("Name %v", i))
    aMap.AddObject(foo)
  }

  //Accessing map
  foo1 := aMap.Object("1")
  JSON, err := json.Marshal(foo1)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", JSON)

  //Iterating map
  err = aMap.Pairs(func(key interface{}, item *gtly.Object) (bool, error) {
    fmt.Printf("id: %v\n", item.Value("id"))
    fmt.Printf("name: %v\n", item.Value("name"))
    return true, nil
  })
  if err != nil {
    log.Fatal(err)
  }

  JSON, err = json.Marshal(aMap)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", JSON)
  //[{"id":1,"firsName":"Name 1","updated":null,"numbers":null},{"id":3,"firsName":"Name 3","updated":null,"numbers":null},{"id":4,"firsName":"Name 4","updated":null,"numbers":null},{"id":6,"firsName":"Name 6","updated":null,"numbers":null},{"id":8,"firsName":"Name 8","updated":null,"numbers":null},{"id":9,"firsName":"Name 9","updated":null,"numbers":null},{"id":0,"firsName":"Name 0","updated":null,"numbers":null},{"id":2,"firsName":"Name 2","updated":null,"numbers":null},{"id":5,"firsName":"Name 5","updated":null,"numbers":null},{"id":7,"firsName":"Name 7","updated":null,"numbers":null}]
}
```

#### MultiMap

```go
func NewMultiMap_Usage() {
  fooProvider, err := gtly.NewProvider("foo",
    gtly.NewField("id", gtly.FieldTypeInt),
    gtly.NewField("firsName", gtly.FieldTypeString),
    gtly.NewField("city", gtly.FieldTypeString, gtly.OmitEmptyOpt(true)),
    gtly.NewField("updated", gtly.FieldTypeTime, gtly.DateLayoutOpt(time.RFC3339)),
    gtly.NewField("numbers", gtly.FieldTypeArray, gtly.ComponentTypeOpt(gtly.FieldTypeInt)),
    )
  if err != nil {
     log.Fatal(err)
  }
  
  //Creates a multi map keyed by id Field
  aMap := fooProvider.NewMultimap(gtly.NewKeyProvider("city"))
  for i := 0; i < 10; i++ {
    foo := fooProvider.NewObject()
    foo.SetValue("id", i)
    foo.SetValue("firsName", fmt.Sprintf("Name %v", i))
    if i%2 == 0 {
        foo.SetValue("city", "Cracow")
    } else {
        foo.SetValue("city", "Warsaw")
    }
    aMap.AddObject(foo)
  }
  
  //Accessing map
  fooInWarsawSlice := aMap.Slice("Warsaw")
  JSON, err := json.Marshal(fooInWarsawSlice)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", JSON)
  //Prints [{"id":1,"firsName":"Name 1","city":"Warsaw","updated":null,"numbers":null},{"id":3,"firsName":"Name 3","city":"Warsaw","updated":null,"numbers":null},{"id":5,"firsName":"Name 5","city":"Warsaw","updated":null,"numbers":null},{"id":7,"firsName":"Name 7","city":"Warsaw","updated":null,"numbers":null},{"id":9,"firsName":"Name 9","city":"Warsaw","updated":null,"numbers":null}]
  
  //Iterating multi map
  err = aMap.Slices(func(key interface{}, value *gtly.Array) (bool, error) {
    fmt.Printf("%v -> %v\n", key, value.Size())
    return true, nil
  })
  if err != nil {
    log.Fatal(err)
  }
  
  JSON, err = json.Marshal(aMap)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", JSON)
  //[{"id":0,"firsName":"Name 0","city":"Cracow","updated":null,"numbers":null},{"id":2,"firsName":"Name 2","city":"Cracow","updated":null,"numbers":null},{"id":4,"firsName":"Name 4","city":"Cracow","updated":null,"numbers":null},{"id":6,"firsName":"Name 6","city":"Cracow","updated":null,"numbers":null},{"id":8,"firsName":"Name 8","city":"Cracow","updated":null,"numbers":null},{"id":1,"firsName":"Name 1","city":"Warsaw","updated":null,"numbers":null},{"id":3,"firsName":"Name 3","city":"Warsaw","updated":null,"numbers":null},{"id":5,"firsName":"Name 5","city":"Warsaw","updated":null,"numbers":null},{"id":7,"firsName":"Name 7","city":"Warsaw","updated":null,"numbers":null},{"id":9,"firsName":"Name 9","city":"Warsaw","updated":null,"numbers":null}]

}
```


## Contributing to gtly

Gtly is an open source project and contributors are welcome!

See [TODO](TODO.md) list

## License

The source code is made available under the terms of the Apache License, Version 2, as stated in the file `LICENSE`.

Individual files may be made available under their own specific license,
all compatible with Apache License, Version 2. Please see individual files for details.

<a name="Credits-and-Acknowledgements"></a>

## Credits and Acknowledgements

**Library Author:** Adrian Witas

