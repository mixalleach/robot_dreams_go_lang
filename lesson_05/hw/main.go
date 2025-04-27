package main

import (
	"fmt"
	"lesson05/hw/documentstore"
	"lesson05/hw/users"
)

func main() {
	store := documentstore.NewStore()
	usersColl, err := store.CreateCollection(
		"Users",
		&documentstore.CollectionConfig{PrimaryKey: "key"},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	userService := users.Service{Coll: usersColl}

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

	err = userService.DeleteUser("1")
	if err != nil {
		fmt.Printf("cant delete user")
	}

	updatedUsersList, _ := userService.ListUsers()
	fmt.Printf("users collection list: %+v\n", updatedUsersList)

	user, err := userService.GetUser("1")
	if err != nil {
		fmt.Printf("cant get user")
	} else {
		fmt.Printf("user: %+v\n", user)
	}
}
