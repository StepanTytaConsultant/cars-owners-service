package urlvals

import (
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const (
	searchTag = "search"
	filterTag = "filter"
	pageTag   = "page"
)

func PopulateParams(values url.Values, i interface{}) error {
	t := reflect.ValueOf(i)
	if t.Kind() != reflect.Pointer {
		return errors.New(fmt.Sprintf("interface type must be pointer, got: %v", t.Kind()))
	}

	el := t.Elem()
	if el.Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("interface type must be pointer to a struct, got: %v", el.Kind()))
	}

	for k, v := range values {
		if strings.HasPrefix(k, searchTag) {
			if err := parseTag(el, searchTag, k, v); err != nil {
				return errors.Wrap(err, "failed to parse search tag")
			}
		} else if strings.HasPrefix(k, filterTag) {
			if err := parseTag(el, filterTag, k, v); err != nil {
				return errors.Wrap(err, "failed to parse filter tag")
			}
		} else if strings.HasPrefix(k, pageTag) {
			if err := parseTag(el, pageTag, k, v); err != nil {
				return errors.Wrap(err, "failed to parse page tag")
			}
		}
	}
	return nil
}

func parseTag(el reflect.Value, tag, key string, values []string) error {
	if !el.IsValid() {
		return nil
	}

	if el.Kind() == reflect.Ptr {
		el = el.Elem()
	}

	if el.Kind() == reflect.Struct {
		for i := 0; i < el.NumField(); i++ {
			if err := parseTag(el.Field(i), tag, key, values); err != nil {
				return errors.Wrap(err, "failed to parse tag")
			}
		}

		fieldName := findField(el.Type(), tag, getKey(tag, key))
		if fieldName == "" {
			return nil
		}

		f := el.FieldByName(fieldName)
		if err := setField(f, values); err != nil {
			return errors.Wrap(err, "failed to set field")
		}
	}

	return nil
}

func setField(v reflect.Value, values []string) error {
	if v.IsValid() {
		switch v.Kind() {
		case reflect.Pointer:
			v.Set(reflect.New(v.Type().Elem()))
			return setField(v.Elem(), values)
		case reflect.String:
			if len(values) == 1 {
				v.SetString(values[0])
			}
		case reflect.Int32:
			if len(values) == 1 {
				val, err := strconv.ParseInt(values[0], 10, 64)
				if err != nil {
					return errors.Wrap(err, "failed to parse int")
				}
				v.SetInt(val)
			}
		case reflect.Slice:
			v.Set(reflect.ValueOf(values))
		default:
			panic("Unimplemented type tag")
		}
		return nil
	}
	return errors.New("field is invalid, and cannot be set")
}

func findField(t reflect.Type, tagName, key string) string {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get(tagName)

		if tag == key {
			return field.Name
		}
	}
	return ""
}
