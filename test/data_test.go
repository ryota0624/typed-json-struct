package test

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDataDetailJSONConvert(t *testing.T) {
	detail := &DataDetail1{
		StringField: "Hoge",
	}
	detailJSON, err := DataDetailJSONFromDetail(detail)
	if err != nil {
		panic(err)
	}
	data := DataJSON{
		Type:   DataType1,
		Detail: *detailJSON,
	}

	detail2, err := data.Detail.ToDetail()
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(detail, detail2) {
		t.Errorf("%+v\n%+v\n", detail, detail2)
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	t.Logf("dataJSON: %s", dataJSON)

	var data2 DataJSON

	err = json.Unmarshal(dataJSON, &data2)
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(data, data2) {
		t.Errorf("%+v\n%+v\n", data, data2)
	}
}
