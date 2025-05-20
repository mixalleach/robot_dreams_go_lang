package main

import (
	"fmt"
	"lesson06/hw/documentstore"
	"lesson06/hw/users"
)

func main() {
	store := documentstore.NewStore()

	userService, err := users.NewService(store)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = userService.CreateUser("1", "Bob")
	if err != nil {
		fmt.Printf("cant create new user")
	}

	_, err = userService.CreateUser("2", "Ned")
	if err != nil {
		fmt.Printf("cant create new user")
	}

	usersList, _ := userService.ListUsers()
	fmt.Printf("users collection list: %+v\n", usersList)

	updatedUsersList, _ := userService.ListUsers()
	fmt.Printf("users collection list: %+v\n", updatedUsersList)

	user, err := userService.GetUser("1")
	if err != nil {
		fmt.Printf("cant get user\n")
	} else {
		fmt.Printf("user: %+v\n", user)
	}

	err = store.DumpToFile("dump.json")
	if err != nil {
		return
	}

	clonedStore, err := documentstore.NewStoreFromFile("dump.json")
	if err != nil {
		fmt.Printf("error while cloning store: %s\n", err)
	}

	clonedCollection, err := clonedStore.GetCollection("Users")
	if err != nil {
		fmt.Printf("error while cloning collection: %s\n", err)
	}

	fmt.Printf("clonned users collection list: %+v\n", clonedCollection.List())
}
