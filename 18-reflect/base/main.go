package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	age  int
	Sex  int
}

// reflect  TypeOf,ValueOf
func ReflectBase(i interface{}) {
	//argType := reflect.TypeOf(a)
	v := reflect.ValueOf(i)
	t := v.Type()

	fmt.Printf("a=%v\n", i)
	typeKind := v.Kind()
	switch typeKind {
	case reflect.Int:
		fmt.Printf("a is int\n")
	case reflect.Float64:
		fmt.Printf("a is Float64\n")
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fmt.Printf("name:%s type:%v value:%v \n", t.Field(i).Name, t.Field(i).Type, field.Interface())
		}
	}
}

// reflect set value
func ReflectSetValue(i interface{}) {
	value := reflect.ValueOf(i)
	t := value.Kind()
	switch t {
	case reflect.Int:
		value.SetInt(10000)
		break
	case reflect.Int64:
		value.SetInt(10000)
		break
	case reflect.Float64:
		value.SetFloat(8.8)
		break
	}
	fmt.Printf("i=%v\n", i)
}

func main() {
	a := 100
	ReflectBase(a)
	b := 1.0
	ReflectBase(b)
	var c int64 = 1
	ReflectSetValue(&c)

	var student Student
	ReflectBase(&student)
}
