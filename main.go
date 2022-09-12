package main

import (
	"fmt"
	"reflect"
)

type Boy struct {
	Name string
	Wife Girl
}

type Girl struct {
	Name string
}

func main() {

	girl := &Girl{
		Name: "Mary",
	}
	boy := &Boy{
		Name: "Tom",
		Wife: *girl,
	}

	fieldName := "Wife"
	pointer := reflect.ValueOf(boy)
	field := reflect.Indirect(pointer).FieldByName(fieldName)
	fmt.Print(field.Interface().(Girl).Name)
}
