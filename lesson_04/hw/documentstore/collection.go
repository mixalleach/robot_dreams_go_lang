package documentstore

import "fmt"

type Collection struct {
	Name             string
	Documents        map[string]*Document
	CollectionConfig CollectionConfig
}

type CollectionConfig struct {
	PrimaryKey string
}

func (s *Collection) Put(doc Document) {
	key, ok := doc.Fields["key"]
	if !ok {
		fmt.Println("Field 'key' does not exist")
		return
	}

	if len(key.Value.(string)) == 0 {
		fmt.Println("Field 'key' is empty")
		return
	}

	s.Documents[key.Value.(string)] = &doc
}

func (s *Collection) Get(key string) (*Document, bool) {
	doc, ok := s.Documents[key]
	if !ok {
		return nil, false
	}

	return doc, true
}

func (s *Collection) Delete(key string) bool {
	_, ok := s.Documents[key]
	if !ok {
		return false
	}

	delete(s.Documents, key)

	return true
}

func (s *Collection) List() []Document {
	documents := make([]Document, 0, len(s.Documents))

	for _, doc := range s.Documents {
		documents = append(documents, *doc)
	}

	return documents
}
