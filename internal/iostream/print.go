package iostream

import (
	"fmt"
	"reflect"
)

func PrintStruct(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	// Iterate through each field of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i).Interface()

		fmt.Printf("%s: %v\n", field.Name, fieldValue)
	}
}

func PrintObj(arrObj interface{}) {
	v := reflect.ValueOf(arrObj)

	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i).Interface()
			fmt.Println("Element", i)
			PrintStruct(elem)
		}
	} else {
		fmt.Println("Unsupported type:", reflect.TypeOf(arrObj))
	}
}
