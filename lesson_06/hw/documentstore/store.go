package documentstore

import (
	"encoding/json"
	"errors"
)

type Store struct {
	Collection map[string]*Collection `json:"collections"`
}

func NewStore() *Store {
	return &Store{make(map[string]*Collection)}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (*Collection, error) {
	if _, ok := s.Collection[name]; ok {
		return nil, errors.New("Collection '" + name + "' already exists")
	}

	newCollection := Collection{
		Name:      name,
		documents: make(map[string]*Document),
		cfg:       *cfg,
	}

	s.Collection[name] = &newCollection

	return &newCollection, nil
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	collection, ok := s.Collection[name]
	if !ok {
		return nil, errors.New("Collection '" + name + "' not found")
	}

	return collection, nil
}

func (s *Store) DeleteCollection(name string) (bool, error) {
	_, ok := s.Collection[name]
	if !ok {
		return false, errors.New("Collection '" + name + "' not found")
	}

	delete(s.Collection, name)

	return true, nil
}

func NewStoreFromDump(dump []byte) (*Store, error) {
	// Функція повинна створити та проініціалізувати новий `Store`
	// зі всіма колекціями да даними з вхідного дампу.

	return &Store{}, nil
}

func (s *Store) Dump() ([]byte, error) {
	// Методи повинен віддати дамп нашого стору в який включені дані про колекції та документ
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Значення яке повертає метод `store.Dump()` має без помилок оброблятись функцією `NewStoreFromDump`

func NewStoreFromFile(filename string) (*Store, error) {
	// Робить те ж саме що і функція `NewStoreFromDump`, але сам дамп має діставатись з файлу

	return &Store{}, nil
}

func (s *Store) DumpToFile(filename string) error {
	// Робить те ж саме що і метод  `Dump`, але записує у файл замість того щоб повертати сам дамп

	return nil
}
