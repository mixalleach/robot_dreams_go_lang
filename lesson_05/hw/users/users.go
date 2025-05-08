package users

import (
	"errors"
	"lesson05/hw/documentstore"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	coll *documentstore.Collection
}

func NewService(store *documentstore.Store) (*Service, error) {
	usersColl, err := store.CreateCollection(
		"Users",
		&documentstore.CollectionConfig{PrimaryKey: "id"},
	)

	return &Service{coll: usersColl}, err
}

func (s *Service) CreateUser(userID string, userName string) (*User, error) {
	doc := User{ID: userID, Name: userName}

	marshalled, err := documentstore.MarshalDocument(doc)
	if err != nil {
		return nil, err
	}

	err = s.coll.Put(*marshalled)

	return &doc, err
}

func (s *Service) GetUser(userID string) (*User, error) {
	var user User
	doc, err := s.coll.Get(userID)
	if err != nil {
		return nil, errors.New("user '" + userID + "' not found")
	}

	err = documentstore.UnmarshalDocument(doc, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *Service) ListUsers() ([]User, error) {
	docs := s.coll.List()
	users := make([]User, 0, len(docs))

	for _, doc := range docs {
		var user User

		err := documentstore.UnmarshalDocument(&doc, &user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *Service) DeleteUser(userID string) error {
	err := s.coll.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}
