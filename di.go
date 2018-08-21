package di

import (
	"fmt"
	"reflect"
)

// EnsureAllFieldsSet ensures all fields are set in an object.
// If this is not the case, it returns an error.
//
// The intent of this func is that if you have a constructor that
// takes its dependencies as a struct, you can validate them all in
// one shot with this method.
func EnsureAllFieldsSet(obj interface{}) error {
	val := reflect.ValueOf(obj)
	nf := val.NumField()
	var nv []int
	for i := 0; i < nf; i++ {
		fv := val.Field(i)
		if fv.IsNil() {
			nv = append(nv, i)
		}
	}
	if len(nv) > 0 {
		return fmt.Errorf("in %s, fields with no value: %v",
			reflect.TypeOf(obj), nv)
	}
	return nil
}

// PanicUnlessAllFieldsSet will panic if EnsureAllFieldsSet
// would return an error.
//
// The intent of this func is that if you have a constructor that
// takes its dependencies as a struct, you can validate them all in
// one shot with this method -- and not even have to bother to return
// an error.
func PanicUnlessAllFieldsSet(obj interface{}) {
	e := EnsureAllFieldsSet(obj)
	if e != nil {
		panic(e)
	}
}
