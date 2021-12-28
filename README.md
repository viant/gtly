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
	"github.com/viant/gtly/codec/json"
	"log"
	"time"
)

func NewObject_Usage() {
	fooProvider := gtly.NewProvider("foo",  //create foo type 
		gtly.NewField("id", gtly.FieldTypeInt),
		gtly.NewField("firsName", gtly.FieldTypeString),
		gtly.NewField("description", gtly.FieldTypeString, gtly.OmitEmptyOpt(true)),
		gtly.NewField("updated", gtly.FieldTypeTime, gtly.DateLayoutOpt("2006-01-02T15:04:05Z07:00")),
		gtly.NewField("numbers", gtly.FieldTypeArray, gtly.ComponentTypeOpt(gtly.FieldTypeInt)),
	)
    //create an instance of foo type
	foo1 := fooProvider.NewObject()
	foo1.SetInt("id", 1)
	foo1.SetString("firsName", "Adam")
	foo1.SetTime("updated", time.Now())
	foo1.SetValue("numbers", []int{1, 2, 3})
	JSON, err := json.Marshal(foo1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", JSON)
    //Prints: {"id":1,"firsName":"Adam","updated":"2020-09-25T20:43:40-07:00","numbers":[1,2,3]}
    //change foo type output to upper underscore
	fooProvider.OutputCaseFormat(gtly.CaseLowerCamel, gtly.CaseUpperUnderscore)
	JSON, _ = json.Marshal(foo1)
	fmt.Printf("%s\n", JSON)
    //Prints: {"ID":1,"FIRS_NAME":"Adam","UPDATED":"2020-09-25T20:43:40-07:00","NUMBERS":[1,2,3]}
    //change foo type output to lower underscore
	fooProvider.OutputCaseFormat(gtly.CaseLowerCamel, gtly.CaseLowerUnderscore)
    foo1.SetBool("active", true) //add dynamically new field
	foo1.SetString("description", "some description") // set value for existing field
	JSON, _ = json.Marshal(foo1)
	fmt.Printf("%s\n", JSON)
    //Prints: {"id":1,"firs_name":"Adam","description":"some description","updated":"2020-09-25T20:43:40-07:00","numbers":[1,2,3],"active":true}
   //create another instance of foo type
    foo2 := fooProvider.NewObject()
    JSON, _ = json.Marshal(foo2)
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
    fooProvider := gtly.NewProvider("foo",
        gtly.NewField("id", gtly.FieldTypeInt),
        gtly.NewField("firsName", gtly.FieldTypeString),
        gtly.NewField("income", gtly.FieldTypeFloat),
        gtly.NewField("description", gtly.FieldTypeString, gtly.OmitEmptyOpt(true)),
        gtly.NewField("updated", gtly.FieldTypeTime, gtly.DateLayoutOpt(time.RFC3339)),
        gtly.NewField("numbers", gtly.FieldTypeArray, gtly.ComponentTypeOpt(gtly.FieldTypeInt)),
    )
    fooArray1 := fooProvider.NewArray()
    for i:=0;i<10;i++ {
        foo1 := fooProvider.NewObject()
        foo1.SetInt("id", 1)
        foo1.SetFloat("income", 64000.0*float64(1+(10/(i+1))))
        foo1.SetString("firsName", "Adam")
        fooArray1.AddObject(foo1)
    }
    now := time.Now()
    fooArray1.Add(map[string]interface{}{
        "id":      100,
        "firsName":    "Tom",
        "updated": now,
    })
    totalIncome := 0.0
    incomeField := fooProvider.Field("income")
    //Iterating collection
    err := fooArray1.Objects(func(object *gtly.Object) (bool, error) {
        fmt.Printf("id: %v\n",  object.Int("id"))
        fmt.Printf("name: %v\n",  object.String("name"))
        totalIncome +=  object.FloatAt(incomeField.Index)
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
	fooProvider := gtly.NewProvider("foo",
		gtly.NewField("id", gtly.FieldTypeInt),
		gtly.NewField("firsName", gtly.FieldTypeString),
		gtly.NewField("description", gtly.FieldTypeString, gtly.OmitEmptyOpt(true)),
		gtly.NewField("updated", gtly.FieldTypeTime, gtly.DateLayoutOpt(time.RFC3339)),
		gtly.NewField("numbers", gtly.FieldTypeArray, gtly.ComponentTypeOpt(gtly.FieldTypeInt)),
	)
	//Creates a map keyed by id field
	aMap := fooProvider.NewMap(gtly.NewIndex([]string{"id"}))
	for i := 0; i < 10; i++ {
		foo := fooProvider.NewObject()
		foo.SetInt("id", i)
		foo.SetString("firsName", fmt.Sprintf("Name %v", i))
		aMap.AddObject(foo)
	}
	//Accessing map
	foo1 := aMap.Object("1")
	JSON, err := json.Marshal(foo1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", JSON)
    //Prints  {"id":1,"firsName":"Name 1","updated":null,"numbers":null}
	//Iterating map
    aMap.Pairs(func(key string, item *gtly.Object) (bool, error) {
        fmt.Printf("id: %v\n",  item.Int("id"))
        fmt.Printf("name: %v\n",  item.String("name"))
        return true, nil
    })
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
	fooProvider := gtly.NewProvider("foo",
		gtly.NewField("id", gtly.FieldTypeInt),
		gtly.NewField("firsName", gtly.FieldTypeString),
		gtly.NewField("city", gtly.FieldTypeString, gtly.OmitEmptyOpt(true)),
		gtly.NewField("updated", gtly.FieldTypeTime, gtly.DateLayoutOpt(time.RFC3339)),
		gtly.NewField("numbers", gtly.FieldTypeArray, gtly.ComponentTypeOpt(gtly.FieldTypeInt)),
	)
	//Creates a multi map keyed by id field
	aMap := fooProvider.NewMultimap(gtly.NewIndex([]string{"city"}))
	for i := 0; i < 10; i++ {
		foo := fooProvider.NewObject()
		foo.SetInt("id", i)
		foo.SetString("firsName", fmt.Sprintf("Name %v", i))
		if i % 2 ==0 {
			foo.SetString("city", "Cracow")
		} else {
			foo.SetString("city", "Warsaw")
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
	err = aMap.Slices(func(key string, value *gtly.Array) (bool, error) {
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

