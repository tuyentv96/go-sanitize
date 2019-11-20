package sanitize

import (
	"errors"
	"reflect"
	"strings"
)

func TrimSpace(input interface{}) error {
	inputVal := reflect.ValueOf(input)
	if inputVal.Kind() != reflect.Ptr {
		return errors.New("not a pointer")
	}

	inputType := reflect.Indirect(inputVal).Type()
	if inputType.Kind() != reflect.Struct {
		return nil
	}

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		val := reflect.Indirect(inputVal.Elem().FieldByName(field.Name))
		switch val.Kind() {
		case reflect.String:
			val.SetString(strings.TrimSpace(val.String()))
		case reflect.Struct:
			if val.CanAddr() && val.Addr().CanInterface() {
				TrimSpace(val.Addr().Interface())
			}
		case reflect.Slice:
			if val.CanInterface() {
				val := reflect.ValueOf(val.Interface())
				for i := 0; i < val.Len(); i++ {
					TrimSpace(val.Index(i).Addr().Interface())
				}
			}
		case reflect.Map:
			if val.CanInterface() {
				val := reflect.ValueOf(val.Interface())
				for _, key := range val.MapKeys() {
					mapValue := val.MapIndex(key)
					mapValuePtr := reflect.New(mapValue.Type())
					mapValuePtr.Elem().Set(mapValue)
					if mapValuePtr.Elem().CanAddr() {
						TrimSpace(mapValuePtr.Elem().Addr().Interface())
					}
					val.SetMapIndex(key, reflect.Indirect(mapValuePtr))
				}
			}
		}
	}

	return nil
}
