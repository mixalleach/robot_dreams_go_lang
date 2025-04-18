package documentstore

type Store struct {
	Collection map[string]*Collection
}

func NewStore() *Store {
	return &Store{make(map[string]*Collection)}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	// Створюємо нову колекцію і повертаємо `true` якщо колекція була створена
	// Якщо ж колекція вже створена то повертаємо `false` та nil

	if _, ok := s.Collection[name]; ok {
		return false, nil
	}

	newCollection := Collection{
		Name:             name,
		Documents:        make(map[string]*Document),
		CollectionConfig: *cfg,
	}

	s.Collection[name] = &newCollection

	return true, &newCollection
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	collection, ok := s.Collection[name]

	if !ok {
		return nil, false
	}

	return collection, true
}

func (s *Store) DeleteCollection(name string) bool {
	_, ok := s.Collection[name]
	if !ok {
		return false
	}

	delete(s.Collection, name)

	return true
}
