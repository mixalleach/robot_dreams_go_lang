package documentstore

import (
	"errors"
)

type Store struct {
	collection map[string]*Collection
}

func NewStore() *Store {
	return &Store{make(map[string]*Collection)}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (*Collection, error) {
	if _, ok := s.collection[name]; ok {
		return nil, errors.New("collection '" + name + "' already exists")
	}

	newCollection := Collection{
		Name:      name,
		documents: make(map[string]*Document),
		cfg:       *cfg,
	}

	s.collection[name] = &newCollection

	return &newCollection, nil
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	collection, ok := s.collection[name]
	if !ok {
		return nil, errors.New("collection '" + name + "' not found")
	}

	return collection, nil
}

func (s *Store) DeleteCollection(name string) (bool, error) {
	_, ok := s.collection[name]
	if !ok {
		return false, errors.New("collection '" + name + "' not found")
	}

	delete(s.collection, name)

	return true, nil
}
