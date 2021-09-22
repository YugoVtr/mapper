package mapper

import (
	"errors"
	"reflect"
)

var (
	errStruct = errors.New("need to be structs")
	errPtr    = errors.New("need to be pointer")
)

// Mapper traces fields in a structure source
// to fields of the same name in target structure.
//   - Source is a Struct.
//   - Target is a reference to Struct.
// If the target has an attribute of type pointer,
// the pointer needs to store a struct to be mapped.
// Fields with same name but different types are ignored.
func Mapper(source interface{}, target interface{}) (err error) {
	t := reflect.ValueOf(target)
	if t.Kind() != reflect.Ptr {
		return errPtr
	}

	rv := t.Elem()
	if rv.Kind() != reflect.Struct {
		return errStruct
	}

	rvs := reflect.ValueOf(source)
	if rvs.Kind() != reflect.Struct {
		return errStruct
	}

	for i := 0; i < rvs.NumField(); i++ {
		f := rvs.Field(i)
		vn := rvs.Type().Field(i).Name
		vt := rvs.Type().Field(i).Type.Kind()

		s := rv.FieldByName(vn)

		if s.Kind() != vt {
			continue
		}

		switch s.Kind() {
		case reflect.Ptr:
			if f.Elem().Kind() == reflect.Struct {
				err = Mapper(f.Elem().Interface(), s.Interface())
			}
		case reflect.Struct:
			err = Mapper(f.Interface(), s.Addr().Interface())
		case vt:
			if s.CanSet() {
				s.Set(f)
			}
		}
	}

	return err
}
