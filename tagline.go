package tagline

import (
	"errors"
	"reflect"
	"strings"
)

var ErrTypeMustBePointer = errors.New("obj type must be a struct pointer")

// Process is a function that handle field with tag `tl` within a struct
func Process(obj interface{}) error {
	if obj == nil {
		return nil
	}

	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return ErrTypeMustBePointer
	}

	v = v.Elem()

	for i := 0; i < v.NumField(); i++ {
		valField := v.Field(i)
		typeField := v.Type().Field(i)
		tag := typeField.Tag
		tagVal := tag.Get("tl")
		processTag(valField, tagVal)
	}

	return nil
}

func processTag(val reflect.Value, tag string) {
	if tag == "upper" {
		if val.Kind() == reflect.String {
			val.SetString(strings.ToUpper(val.String()))
		}
	} else if tag == "lower" {
		if val.Kind() == reflect.String {
			val.SetString(strings.ToLower(val.String()))
		}
	}
}
