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
	collections map[string]*Collection
}

type dumpDocument struct {
	Fields map[string]DocumentField `json:"fields"`
}

type dumpCollection struct {
	Name      string           `json:"name"`
	Documents []dumpDocument   `json:"documents"`
	Cfg       CollectionConfig `json:"cfg"`
}

type dumpStore struct {
	Collections []dumpCollection `json:"collections"`
}

func NewStore() *Store {
	return &Store{make(map[string]*Collection)}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (*Collection, error) {
	if _, ok := s.collections[name]; ok {
		return nil, errors.New("collection '" + name + "' already exists")
	}

	newCollection := Collection{
		name:      name,
		documents: make(map[string]*Document),
		cfg:       *cfg,
	}

	s.collections[name] = &newCollection

	slog.Default().Info(fmt.Sprintf("Collection '%s' created\n", name))

	return &newCollection, nil
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	collection, ok := s.collections[name]
	if !ok {
		return nil, errors.New("collection '" + name + "' not found")
	}

	return collection, nil
}

func (s *Store) DeleteCollection(name string) (bool, error) {
	_, ok := s.collections[name]
	if !ok {
		return false, errors.New("collection '" + name + "' not found")
	}

	delete(s.collections, name)

	slog.Default().Info(fmt.Sprintf("Collection '%s' deleted\n", name))

	return true, nil
}

func NewStoreFromDump(dump []byte) (*Store, error) {
	var tmpStore dumpStore
	store := NewStore()

	if err := json.Unmarshal(dump, &tmpStore); err != nil {
		return nil, err
	}

	for _, tmpCollection := range tmpStore.Collections {
		newCollection, err := store.CreateCollection(tmpCollection.Name, &tmpCollection.Cfg)
		if err != nil {
			return nil, err
		}

		for _, tmpDocument := range tmpCollection.Documents {
			doc := Document{
				Fields: tmpDocument.Fields,
			}
			newCollection.Put(doc)
		}
	}

	return store, nil
}

func (s *Store) Dump() ([]byte, error) {
	tmpStore := dumpStore{
		Collections: make([]dumpCollection, 0, len(s.collections)),
	}

	for _, collection := range s.collections {
		tmpCollection := dumpCollection{
			Name:      collection.name,
			Documents: make([]dumpDocument, 0, len(collection.documents)),
			Cfg:       collection.cfg,
		}

		for _, doc := range collection.documents {
			tmpDocument := dumpDocument{
				Fields: doc.Fields,
			}
			tmpCollection.Documents = append(tmpCollection.Documents, tmpDocument)
		}

		tmpStore.Collections = append(tmpStore.Collections, tmpCollection)
	}

	data, err := json.Marshal(tmpStore)
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
