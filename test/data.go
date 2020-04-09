package test

//go:generate go run github.com/ryota0624/typed-json-struct -package_name=$GOPACKAGE -detail_name=DataDetail -dest_name json_$GOFILE

import (
	"errors"

	"github.com/ryota0624/typed-json-struct/typed"
)

type (
	DataJSON struct {
		Type   DataType       `json:"type"`
		Detail DataDetailJSON `json:"detail"`
	}

	DataDetail interface {
		typed.AnyDetail
	}

	DataType                  string
	DataDetailTypeConstructor struct{}

	DataDetail1 struct {
		StringField string
	}

	DataDetail2 struct {
		IntField int
	}
)

func (d DataDetail2) Type() typed.TypeEnum {
	return DataType2
}

func (d DataDetail1) Type() typed.TypeEnum {
	return DataType1
}

const (
	DataType1 DataType = "1"
	DataType2 DataType = "2"
)

func (d DataDetailTypeConstructor) FromString(str string) (typed.TypeEnum, error) {
	switch str {
	case DataType1.String():
		return DataType1, nil
	case DataType2.String():
		return DataType2, nil
	default:
		return nil, errors.New("not defined type received")
	}
}

func (d DataType) EmptyDetail() (interface{}, error) {
	switch d {
	case DataType1:
		return &DataDetail1{}, nil
	case DataType2:
		return &DataDetail2{}, nil
	default:
		return nil, errors.New("not found detail")
	}
}

func (d DataType) Constructor() typed.TypeEnumConstructor {
	return DataDetailTypeConstructor{}
}

func (d DataType) String() string {
	return string(d)
}

var (
	_ typed.TypeEnum            = DataType(0)
	_ typed.TypeEnumConstructor = DataDetailTypeConstructor{}
	_ DataDetail                = (*DataDetail1)(nil)
	_ DataDetail                = (*DataDetail2)(nil)
)
