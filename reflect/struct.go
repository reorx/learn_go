package main

import (
	"fmt"
	"reflect"
)

// S ...
type S struct {
	Foo string
	Bar int
	Z   int
	A   int
}

func main() {
	{
		x := S{"foo", 2, 0, 0}

		t := reflect.TypeOf(x)
		v := reflect.ValueOf(x)

		fields := make([]interface{}, t.NumField())
		values := make([]interface{}, v.NumField())

		for i := 0; i < v.NumField(); i++ {
			fields[i] = t.Field(i).Name
			values[i] = v.Field(i).Interface()
		}

		fmt.Println("fields", fields)
		fmt.Println("values", values)
	}
	{
		type t struct {
			N int
		}
		var n = t{42}
		// N at start
		fmt.Println(n.N)
		// pointer to struct - addressable
		ps := reflect.ValueOf(&n)
		// struct
		s := ps.Elem()
		if s.Kind() == reflect.Struct {
			// exported field
			f := s.FieldByName("N")
			if f.IsValid() {
				// A Value can be changed only if it is
				// addressable and was not obtained by
				// the use of unexported struct fields.
				if f.CanSet() {
					// change value of N
					if f.Kind() == reflect.Int {
						x := int64(7)
						if !f.OverflowInt(x) {
							f.SetInt(x)
						}
					}
				}
			}
		}
		// N at end
		fmt.Println(n.N)
	}
	// {
	// 	t := reflect.TypeOf(S)
	// 	fmt.Println(t)
	// }
}
