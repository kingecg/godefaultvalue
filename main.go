// Package godefaultvalue is for mark default value and get default for struct
package godefaultvalue

import (
	"encoding/json"
	"reflect"
)

var DefaultValueMap = map[string]map[string]interface{}{}

type Godefault[T any] struct {
}

// Function GetDefault is used to get default value by fieldName
func (g *Godefault[T]) GetDefault(fieldName string) (interface{}, bool) {
	tName := reflect.TypeOf(*g).Name()
	vmap, ok := DefaultValueMap[tName]
	if !ok {
		DefaultValueMap[tName] = GetDefaultValue(reflect.TypeOf((*T)(nil)).Elem())
		vmap = DefaultValueMap[tName]
	}
	ret, ok := vmap[fieldName]
	return ret, ok
}

// Function GetValue is used to get p.[filedName], if get zero, return default value
func (g *Godefault[T]) GetValue(p *T, fieldName string) (interface{}, bool) {
	v := reflect.ValueOf(p).FieldByName(fieldName)
	if v.IsZero() {
		return g.GetDefault(fieldName)
	}
	return v, true
}

func GetDefaultValue(t reflect.Type) map[string]interface{} {
	ret := map[string]interface{}{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("defaultV")
		if field.Type.Kind() != reflect.Struct && field.Type.Kind() != reflect.Ptr {
			x := reflect.Zero(field.Type).Interface()
			err := json.Unmarshal([]byte(tag), &x)
			if err == nil {
				ret[field.Name] = x
			}
		}
	}
	return ret
}
