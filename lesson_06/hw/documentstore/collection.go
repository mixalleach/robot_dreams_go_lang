package documentstore

import (
	"errors"
	"fmt"
	"log/slog"
)

type Collection struct {
	Name      string               `json:"name"`
	Documents map[string]*Document `json:"documents"`
	Cfg       CollectionConfig     `json:"cfg"`
}

type CollectionConfig struct {
	PrimaryKey string
}

func (s *Collection) Put(doc Document) error {
	key, ok := doc.Fields[s.Cfg.PrimaryKey]
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

	s.Documents[k] = &doc

	slog.Default().Info(fmt.Sprintf("Document '%s' added to collection '%s'\n", k, s.Name))

	return nil
}

func (s *Collection) Get(key string) (*Document, error) {
	doc, ok := s.Documents[key]
	if !ok {
		return nil, errors.New("document '" + key + "' not found")
	}

	return doc, nil
}

func (s *Collection) Delete(key string) error {
	_, ok := s.Documents[key]
	if !ok {
		return errors.New("document '" + key + "' not found")
	}

	delete(s.Documents, key)

	slog.Default().Info(fmt.Sprintf("Document '%s' deleted from collection '%s'\n", key, s.Name))

	return nil
}

func (s *Collection) List() []Document {
	documents := make([]Document, 0, len(s.Documents))

	for _, doc := range s.Documents {
		documents = append(documents, *doc)
	}

	return documents
}
