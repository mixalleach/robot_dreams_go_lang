package documentstore

import (
	"errors"
)

type Collection struct {
	Name      string
	documents map[string]*Document
	cfg       CollectionConfig
}

type CollectionConfig struct {
	PrimaryKey string
}

func (s *Collection) Put(doc Document) error {
	key, ok := doc.Fields[s.cfg.PrimaryKey]
	if !ok {
		return errors.New("field 'key' does not exist")
	}

	k, ok := key.Value.(string)
	if !ok {
		return errors.New("field 'key' is not a string")
	}

	if len(k) == 0 {
		return errors.New("field 'key' is empty string")
	}

	s.documents[k] = &doc

	return nil
}

func (s *Collection) Get(key string) (*Document, error) {
	doc, ok := s.documents[key]
	if !ok {
		return nil, errors.New("document '" + key + "' not found")
	}

	return doc, nil
}

func (s *Collection) Delete(key string) error {
	_, ok := s.documents[key]
	if !ok {
		return errors.New("document '" + key + "' not found")
	}

	delete(s.documents, key)

	return nil
}

func (s *Collection) List() []Document {
	documents := make([]Document, 0, len(s.documents))

	for _, doc := range s.documents {
		documents = append(documents, *doc)
	}

	return documents
}
