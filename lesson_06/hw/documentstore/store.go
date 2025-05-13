package documentstore

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

type Store struct {
	Collections map[string]*Collection `json:"collections"`
}

func NewStore() *Store {
	return &Store{make(map[string]*Collection)}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (*Collection, error) {
	if _, ok := s.Collections[name]; ok {
		return nil, errors.New("collection '" + name + "' already exists")
	}

	newCollection := Collection{
		Name:      name,
		Documents: make(map[string]*Document),
		Cfg:       *cfg,
	}

	s.Collections[name] = &newCollection

	slog.Default().Info(fmt.Sprintf("Collection '%s' created\n", name))

	return &newCollection, nil
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	collection, ok := s.Collections[name]
	if !ok {
		return nil, errors.New("collection '" + name + "' not found")
	}

	return collection, nil
}

func (s *Store) DeleteCollection(name string) (bool, error) {
	_, ok := s.Collections[name]
	if !ok {
		return false, errors.New("collection '" + name + "' not found")
	}

	delete(s.Collections, name)

	slog.Default().Info(fmt.Sprintf("Collection '%s' deleted\n", name))

	return true, nil
}

func NewStoreFromDump(dump []byte) (*Store, error) {
	var store Store

	if err := json.Unmarshal(dump, &store); err != nil {
		return nil, err
	}

	return &store, nil
}

func (s *Store) Dump() ([]byte, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewStoreFromFile(filename string) (*Store, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	store, err := NewStoreFromDump(data)

	return store, nil
}

func (s *Store) DumpToFile(filename string) error {
	dump, err := s.Dump()
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(dump)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
