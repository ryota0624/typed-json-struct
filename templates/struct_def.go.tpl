package {{.Package}}

import (
	"encoding/json"

	"github.com/ryota0624/typed-json-struct/typed"
)

type (
    {{.DetailName}}JSON typed.TypedJSON
)


func {{.DetailName}}JSONFromDetail(detail {{.DetailName}}) (*{{.DetailName}}JSON, error) {
	typedJSON, err := typed.JSONFromDetail(detail)
	if err != nil {
		return nil, err
	}

	j := {{.DetailName}}JSON(*typedJSON)
	return &j, nil
}

func (c *{{.DetailName}}JSON) UnmarshalJSON(b []byte) error {
	var typedJSON typed.TypedJSON
	if err := json.Unmarshal(b, &typedJSON); err != nil {
		return err
	}

	c.Type = typedJSON.Type
	c.Body = typedJSON.Body
	c.EnumConstructor = {{.EnumConstructor}}{}
	return nil
}


func (c *{{.DetailName}}JSON) ToDetail() ({{.DetailName}}, error) {
	j := &typed.TypedJSON{
		Type: c.Type,
		Body: c.Body,
		EnumConstructor:  c.EnumConstructor,
	}

	detail, err := j.ToDetail()
	if err != nil {
		return nil, err
	}

	return detail.({{.DetailName}}), nil
}
