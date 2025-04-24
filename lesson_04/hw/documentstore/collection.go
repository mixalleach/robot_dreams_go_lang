package documentstore

import "fmt"

type Collection struct {
	Name      string
	documents map[string]*Document
	cfg       CollectionConfig
}

type CollectionConfig struct {
	PrimaryKey string
}

func (s *Collection) Put(doc Document) {
	key, ok := doc.Fields[s.cfg.PrimaryKey]
	if !ok {
		fmt.Println("Field 'key' does not exist")
		return
	}

	if len(key.Value.(string)) == 0 {
		fmt.Println("Field 'key' is empty")
		return
	}

	k, ok := key.Value.(string)
	if !ok {
		fmt.Println("Field 'key' is not a string")
		return
	}
	s.documents[k] = &doc
}

func (s *Collection) Get(key string) (*Document, bool) {
	doc, ok := s.documents[key]
	if !ok {
		return nil, false
	}

	return doc, true
}

func (s *Collection) Delete(key string) bool {
	_, ok := s.documents[key]
	if !ok {
		return false
	}

	delete(s.documents, key)

	return true
}

func (s *Collection) List() []Document {
	documents := make([]Document, 0, len(s.documents))

	for _, doc := range s.documents {
		documents = append(documents, *doc)
	}

	return documents
}
