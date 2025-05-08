package documentstore

import (
	"encoding/json"
	"errors"
)

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

type Document struct {
	Fields map[string]DocumentField
}

func MarshalDocument(input any) (*Document, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return nil, errors.New("failed to marshal user")
	}

	var rawMap map[string]interface{}
	if err := json.Unmarshal(raw, &rawMap); err != nil {
		return nil, errors.New("failed to unmarshal user json")
	}

	fields := make(map[string]DocumentField, len(rawMap))
	for k, v := range rawMap {
		fields[k] = DocumentField{
			Type:  GetFieldTypeByValue(v),
			Value: v,
		}
	}

	doc := Document{
		Fields: fields,
	}

	return &doc, err
}

func UnmarshalDocument(doc *Document, output any) error {
	raw := make(map[string]interface{})
	for k, v := range doc.Fields {
		raw[k] = v.Value
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return errors.New("failed to marshal document")
	}

	if err := json.Unmarshal(data, output); err != nil {
		return errors.New("failed to unmarshal document")
	}

	return nil
}

func GetFieldTypeByValue(val interface{}) DocumentFieldType {
	switch val.(type) {
	case string:
		return DocumentFieldTypeString
	case float64, int, int32, int64, float32:
		return DocumentFieldTypeNumber
	case bool:
		return DocumentFieldTypeBool
	case []interface{}:
		return DocumentFieldTypeArray
	case map[string]interface{}:
		return DocumentFieldTypeObject
	default:
		return "unknown"
	}
}
