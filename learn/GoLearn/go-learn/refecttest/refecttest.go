package refecttest

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}

func Walk2(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func Walk3(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			Walk3(field.Interface(), fn)
		}
	}
}

func Walk4(x interface{}, fn func(input string)) {
	val := getValue(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk4(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func Walk5(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		Walk5(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}
}
