package debug

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	FieldPage = "page"
	FieldID   = "id"
)

func Struct(i interface{}) string {
	typ := reflect.ValueOf(i)

	if typ.Kind() != reflect.Struct {
		typ = typ.Elem()
	}
	res := strings.Builder{}

	for i := 0; i < typ.NumField(); i++ {
		res.WriteString(fmt.Sprintf("%s => %v; ", typ.Type().Field(i).Name, typ.Field(i)))
	}

	return res.String()
}
