package main

import (
	"fmt"
	"lesson03/hw/documentstore"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	stringDoc := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Value: "string_key",
			},
			"value": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "string_value",
			},
		},
	}

	invalidStringDoc := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Value: 123,
			},
			"value": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "string_value",
			},
		},
	}

	numberDoc := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Value: "number_key",
			},
			"value": {
				Type:  documentstore.DocumentFieldTypeNumber,
				Value: 1,
			},
		},
	}

	boolDoc := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Value: "bool_key",
			},
			"value": {
				Type:  documentstore.DocumentFieldTypeBool,
				Value: true,
			},
		},
	}

	arrayDoc := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Value: "array_key",
			},
			"value": {
				Type:  documentstore.DocumentFieldTypeArray,
				Value: [...]int{1, 2, 3, 4},
			},
		},
	}

	objectDoc := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Value: "object_key",
			},
			"value": {
				Type: documentstore.DocumentFieldTypeObject,
				Value: Person{
					Name: "Bob",
					Age:  31,
				},
			},
		},
	}

	documentstore.Put(stringDoc)
	documentstore.Put(invalidStringDoc)
	documentstore.Put(numberDoc)
	documentstore.Put(boolDoc)
	documentstore.Put(arrayDoc)
	documentstore.Put(objectDoc)

	fmt.Printf("%+v\n", documentstore.List())
	fmt.Printf("%+v\n", documentstore.Delete("string_key"))
	fmt.Printf("%+v\n", documentstore.List())

	gotStringDoc, ok := documentstore.Get("string_key")
	if ok {
		fmt.Printf("%+v\n", gotStringDoc)
	}

	gotArrayDoc, ok := documentstore.Get("array_key")
	if ok {
		fmt.Printf("%+v\n", gotArrayDoc)
	}

	gotObjectDoc, ok := documentstore.Get("object_key")
	if ok {
		fmt.Printf("%+v\n", gotObjectDoc)
	}
}
