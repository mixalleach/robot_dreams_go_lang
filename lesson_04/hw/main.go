package main

import (
	"fmt"
	"lesson04/hw/documentstore"
)

func main() {
	//1. Потрібно модифікувати пакет documentstore з 3-го заняття відповідно до опису структур та функцій
	//2. Данні кожної колекції зберігаються в структурі колекції (не в глобальній змінній)

	store := documentstore.NewStore()
	store.CreateCollection(
		"Products",
		&documentstore.CollectionConfig{PrimaryKey: "key"},
	)

	productsCollection, ok := store.GetCollection("Products")
	if !ok {
		fmt.Println("Products collection not found")
	}

	productsCollection.Put(
		documentstore.Document{
			Fields: map[string]documentstore.DocumentField{
				"key": {
					Value: "string_key",
				},
				"value": {
					Type:  documentstore.DocumentFieldTypeString,
					Value: "string_value",
				},
			},
		},
	)

	fmt.Printf("Product collection list: %+v\n", productsCollection.List())
}
