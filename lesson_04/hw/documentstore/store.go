package documentstore

type Store struct {
	collection map[string]*Collection
}

func NewStore() *Store {
	return &Store{make(map[string]*Collection)}
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	// Створюємо нову колекцію і повертаємо `true` якщо колекція була створена
	// Якщо ж колекція вже створена то повертаємо `false` та nil

	if _, ok := s.collection[name]; ok {
		return false, nil
	}

	newCollection := Collection{
		Name:      name,
		documents: make(map[string]*Document),
		cfg:       *cfg,
	}

	s.collection[name] = &newCollection

	return true, &newCollection
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	collection, ok := s.collection[name]

	if !ok {
		return nil, false
	}

	return collection, true
}

func (s *Store) DeleteCollection(name string) bool {
	_, ok := s.collection[name]
	if !ok {
		return false
	}

	delete(s.collection, name)

	return true
}
