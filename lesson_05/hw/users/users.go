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

	marshalled, err := MarshalDocument(doc)
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

	err = UnmarshalDocument(doc, &user)
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

		err := UnmarshalDocument(&doc, &user)
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

func MarshalDocument(input any) (*documentstore.Document, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, errors.New("failed to marshal user")
	}

	var rawMap map[string]interface{}
	if err := json.Unmarshal(raw, &rawMap); err != nil {
		return nil, errors.New("failed to unmarshal user json")
	}

	fields := make(map[string]documentstore.DocumentField, len(rawMap))
	for k, v := range rawMap {
		fields[k] = documentstore.DocumentField{
			Type:  documentstore.GetFieldTypeByValue(v),
			Value: v,
		}
	}

	doc := documentstore.Document{
		Fields: fields,
	}

	return &doc, err
}

func UnmarshalDocument(doc *documentstore.Document, output any) error {
	raw := make(map[string]interface{})
	for k, v := range doc.Fields {
		raw[k] = v.Value
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return errors.New("failed to marshal document")
	}

	if err := json.Unmarshal(data, output); err != nil {
		return errors.New("failed to unmarshal document")
	}

	return nil
}
