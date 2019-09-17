package struct2JSON

import (
	"reflect"
	"strconv"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func GenerateMapOfJSON(s interface{}) (string, error) {
	v := reflect.ValueOf(s)
	m := "{}"

	for i := 0; i < v.NumField(); i++ {
		s := strconv.Itoa(i)

		var err error
		m, err = sjson.Set(m, s, v.Type().Field(i).Name)
		if err != nil {
			return "", err
		}

		m, err = sjson.Set(m, v.Type().Field(i).Name, s)
		if err != nil {
			return "", err
		}
	}

	return m, nil
}

func CreateReducedJSONBasedOnMap(m string, o interface{}) (string, error) {
	v := reflect.ValueOf(o)
	json := "{}"

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i).Name
		field = gjson.Get(m, field).String()

		var err error
		json, err = sjson.Set(json, field, v.Field(i).Interface())
		if err != nil {
			return "", nil
		}
	}

	return json, nil
}
