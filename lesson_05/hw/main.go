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

	userService := users.Service{usersColl}

	userService.CreateUser("1", "Bob")
	userService.CreateUser("2", "Ned")

	users, _ := userService.ListUsers()
	fmt.Printf("Users collection list: %+v\n", users)

	userService.DeleteUser("1")
	users2, _ := userService.ListUsers()
	fmt.Printf("Users collection list: %+v\n", users2)
}
