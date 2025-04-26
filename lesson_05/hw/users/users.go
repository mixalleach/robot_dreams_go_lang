package users

import (
	"encoding/json"
	"errors"
	"lesson05/hw/documentstore"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	Coll *documentstore.Collection
}

func (s *Service) CreateUser(userID string, userName string) (*User, error) {
	doc := User{ID: userID, Name: userName}
	marshalled, _ := json.Marshal(doc)

	err := s.Coll.Put(
		documentstore.Document{
			Fields: map[string]documentstore.DocumentField{
				"key": {
					Value: doc.ID,
				},
				"value": {
					Type:  documentstore.DocumentFieldTypeString,
					Value: marshalled,
				},
			},
		},
	)

	return &doc, err
}

func (s *Service) ListUsers() ([]User, error) {
	docs := s.Coll.List()
	users := make([]User, 0, len(docs))

	for _, doc := range docs {
		var user User

		json.Unmarshal(doc.Fields["value"].Value.([]byte), &user)
		users = append(users, user)
	}

	return users, nil
}

func (s *Service) GetUser(userID string) (*User, error) {
	var user User
	doc, err := s.Coll.Get(userID)
	if err != nil {
		return nil, errors.New("user '" + userID + "' not found")
	}

	unmarshalErr := json.Unmarshal(doc.Fields["value"].Value.([]byte), &user)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &user, nil
}

func (s *Service) DeleteUser(userID string) error {
	err := s.Coll.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}
