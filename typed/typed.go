package typed

import (
	"encoding/json"
	"fmt"
)

type (
	TypeEnum interface {
		EmptyDetail() (interface{}, error)
		Constructor() TypeEnumConstructor
		fmt.Stringer
	}

	TypeEnumConstructor interface {
		FromString(string) (TypeEnum, error)
	}

	AnyDetail interface {
		Type() TypeEnum
	}

	TypedJSON struct {
		Type            string              `json:"type"`
		Body            json.RawMessage     `json:"body"`
		EnumConstructor TypeEnumConstructor `json:"-"`
	}
)

func (j *TypedJSON) ToDetail() (interface{}, error) {
	typ, err := j.EnumConstructor.FromString(j.Type)
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
		Type:            detail.Type().String(),
		Body:            detailJSON,
		EnumConstructor: detail.Type().Constructor(),
	}, nil
}
