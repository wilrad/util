package util

import (
	"fmt"
	"reflect"
	"strings"
)

// Methods prints the method set of the value x.
func Methods(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s %d\n", t, v.NumMethod())

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name,
			strings.TrimPrefix(methType.String(), "func"))
	}
}
