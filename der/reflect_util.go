package der

import (
	"reflect"
)

func isNil(v interface{}) bool {
	return (v == nil) || (reflect.ValueOf(v).IsNil())
}

func valueSetZero(v reflect.Value) {
	if v.CanSet() {
		zero := reflect.Zero(v.Type())
		v.Set(zero)
	}
}

func valueMake(v reflect.Value) {
	if t := v.Type(); t.Kind() == reflect.Ptr {
		nv := reflect.New(t.Elem())
		v.Set(nv)
	}
}
