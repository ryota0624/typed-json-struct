package test

import (
	"errors"

	"github.com/ryota0624/typed-json-struct/typed"
)

type (
	DataDetail interface {
		typed.AnyDetail
	}

	DataType string
	DataDetailTypeDef struct {}

	DataDetail1 struct {
		StringField string
	}

	DataDetail2 struct {
		IntField int
	}
)

const (
	DataType1 DataType = "1"
	DataType2 DataType = "2"
)

func (d DataDetailTypeDef) TypeFromString(str string) (typed.TypeLabel, error) {
	switch str {
	case DataType1.String():
		return DataType1,nil
	case DataType2.String():
		return DataType2,nil
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

func (d DataType) Def() typed.TypeLabelDef {
	return DataDetailTypeDef{}
}

func (d DataType) String() string {
	return string(d)
}

var (
	_ typed.TypeLabel = DataType(0)
	_ typed.TypeLabelDef = DataDetailTypeDef{}

)