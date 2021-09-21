package mapper

import (
	"errors"
	"reflect"
)

var (
	errStruct = errors.New("need struct")
)

func Mapper(source interface{}, target interface{}) error {
	rv := reflect.ValueOf(target).Elem()
	rvs := reflect.ValueOf(source)

	if rv.Kind() != reflect.Struct {
		return errStruct
	}

	for i := 0; i < rvs.NumField(); i++ {
		vn := rvs.Type().Field(i).Name
		vt := rvs.Type().Field(i).Type.Kind()

		s := rv.FieldByName(vn)

		if s.Kind() != vt {
			continue
		}

		if s.Kind() == reflect.Struct {
			Mapper(rvs.Field(i).Interface(), s.Addr().Interface())
			continue
		}

		if s.CanSet() {
			s.Set(rvs.Field(i))
		}
	}

	return nil
}
