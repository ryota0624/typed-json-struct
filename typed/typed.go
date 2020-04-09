package typed

import (
	"encoding/json"
	"fmt"
)

type (
	TypeLabel interface {
		EmptyDetail() (interface{}, error)
		Def() TypeLabelDef
		fmt.Stringer
	}

	TypeLabelDef interface {
		TypeFromString(string) (TypeLabel, error)
	}

	AnyDetail interface {
		Type() TypeLabel
	}

	TypedJSON struct {
		Type string          `json:"type"`
		Body json.RawMessage `json:"body"`
		Def  TypeLabelDef
	}
)

func (j *TypedJSON) ToDetail() (interface{}, error) {
	typ, err := j.Def.TypeFromString(j.Type)
	if err != nil {
		return nil, err
	}

	detail, err := typ.EmptyDetail()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(j.Body, detail)

	return detail, err
}

func JSONFromDetail(detail AnyDetail) (*TypedJSON, error) {
	detailJSON, err := json.Marshal(detail)
	if err != nil {
		return nil, err
	}
	return &TypedJSON{
		Type: detail.Type().String(),
		Body: detailJSON,
		Def:  detail.Type().Def(),
	}, nil
}
