package structs

import (
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/helpers/array"
	"github.com/bpdlampung/banklampung-core-backend-go/helpers/date"
	"reflect"
	"strconv"
	"strings"
)

func Set(ptr interface{}, tag string) {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		panic("Not a pointer")
	}

	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		if defaultVal := t.Field(i).Tag.Get(tag); defaultVal != "-" {
			if err := setField(v.Field(i), defaultVal); err != nil {
				panic(fmt.Sprintf("%v :: %v", err.Error(), "structs.Set"))
			}

		}
	}
}

func setField(field reflect.Value, defaultVal string) error {

	if !field.CanSet() {
		return fmt.Errorf("Can't set value\n")
	}

	switch field.Kind() {

	case reflect.Int:
		if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
	}

	return nil
}

func GetTagNameFromStruct(object interface{}, tag string) string {
	val := reflect.ValueOf(object)

	// if its a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	// should double check we now have a structs (could still be anything)
	if val.Kind() != reflect.Struct {
		fmt.Println(val.String())
		fmt.Println(val.Kind())
		panic("unexpected type, must structs")
	}

	return toTagNameFromTag(val, tag)
}

func toTagNameFromTag(val reflect.Value, tag string) (tagName string) {
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tagField := typeField.Tag

		if valueField.Kind() == reflect.Struct {
			toTagNameFromTag(valueField, tag)
			continue
		}

		if len(tagField.Get(tag)) != 0 {
			tagName = tagField.Get(tag)
		}
	}

	return tagName
}

func GetStringValueFromStruct(object interface{}, tag, tagName string) string {
	val := reflect.ValueOf(object)

	// if its a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	// should double check we now have a structs (could still be anything)
	if val.Kind() != reflect.Struct {
		panic("unexpected type, must structs")
	}

	return toStringValueFromTag2(val, tag, tagName)
}

// Tag -> qtag
// TagName -> version, id, etc...
func toStringValueFromTag2(val reflect.Value, tag, tagName string) string {
	result := "-"

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tagField := typeField.Tag

		if valueField.Kind() == reflect.Struct {
			result = toStringValueFromTag2(valueField, tag, tagName)
		}

		//fmt.Println(tagField.Get(tag))
		if array.InArray(tagName, strings.Split(tagField.Get(tag), ",")) == tagName {
			//fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tagField.Get(tag))
			return fmt.Sprintf("%s", valueField.Interface())
		}
	}

	return result
}

func DoUpdateMongoEntity(object interface{}) interface{} {
	val := reflect.ValueOf(object)

	// if its a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	// should double check we now have a structs (could still be anything)
	if val.Kind() != reflect.Struct {
		panic("unexpected type, must structs")
	}

	return toUpdateMongoEntity(val)
}

func toUpdateMongoEntity(val reflect.Value) interface{} {
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		if valueField.Kind() == reflect.Struct {
			toUpdateMongoEntity(valueField)
		}

		if array.InArray("updated_date", strings.Split(tag.Get("bson"), ",")) == "updated_date" {
			valueField.Set(reflect.ValueOf(date.TimeNowUTC()))
		}

		if array.InArray("version", strings.Split(tag.Get("bson"), ",")) == "version" {
			valueField.SetUint(valueField.Uint() + 1)
		}
	}

	return val.Interface()
}

func DoCreateMongoEntity(object interface{}) interface{} {
	val := reflect.ValueOf(object)

	// if its a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	// should double check we now have a structs (could still be anything)
	if val.Kind() != reflect.Struct {
		panic("unexpected type, must structs")
	}

	return toCreateMongoEntity(val)
}

func toCreateMongoEntity(val reflect.Value) interface{} {
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		if valueField.Kind() == reflect.Struct {
			toCreateMongoEntity(valueField)
		}

		if array.InArray("updated_date", strings.Split(tag.Get("bson"), ",")) == "updated_date" {
			valueField.Set(reflect.ValueOf(date.TimeNowUTC()))
		}

		if array.InArray("created_date", strings.Split(tag.Get("bson"), ",")) == "created_date" {
			valueField.Set(reflect.ValueOf(date.TimeNowUTC()))
		}

		if array.InArray("version", strings.Split(tag.Get("bson"), ",")) == "version" {
			valueField.SetUint(1)
		}
	}

	return val.Interface()
}
