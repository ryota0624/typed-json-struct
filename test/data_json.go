package test

import (
	"encoding/json"

	"github.com/ryota0624/typed-json-struct/typed"
)

type (
	DataDetailJSON typed.TypedJSON
)

func DataDetailJSONFromDetail(detail DataDetail) (*DataDetailJSON, error) {
	typedJSON, err := typed.JSONFromDetail(detail)
	if err != nil {
		return nil, err
	}

	j := DataDetailJSON(*typedJSON)
	return &j, nil
}

func (c *DataDetailJSON) UnmarshalJSON(b []byte) error {
	var typedJSON typed.TypedJSON
	if err := json.Unmarshal(b, &typedJSON); err != nil {
		return err
	}

	c.Type = typedJSON.Type
	c.Body = typedJSON.Body
	c.EnumConstructor = DataDetailTypeConstructor{}
	return nil
}

func (c *DataDetailJSON) ToDetail() (DataDetail, error) {
	j := &typed.TypedJSON{
		Type:            c.Type,
		Body:            c.Body,
		EnumConstructor: c.EnumConstructor,
	}

	detail, err := j.ToDetail()
	if err != nil {
		return nil, err
	}

	return detail.(DataDetail), nil
}
