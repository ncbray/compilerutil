package goutil

import (
	"encoding/csv"
	"io"
	"reflect"
	"strconv"
)

func ReadCSV(input io.Reader, output interface{}) error {
	output_ptr := reflect.ValueOf(output)
	if output_ptr.Kind() != reflect.Ptr {
		panic("output is not a pointer")
	}
	output_value := output_ptr.Elem()

	if output_value.Kind() != reflect.Slice {
		panic("output is not a slice")
	}
	element_ptr_type := output_value.Type().Elem()
	if element_ptr_type.Kind() != reflect.Ptr {
		panic("output is not a slice of pointers")
	}
	element_struct_type := element_ptr_type.Elem()
	if element_struct_type.Kind() != reflect.Struct {
		panic("output is not a slice of pointers to structs")
	}

	reader := csv.NewReader(input)
	heading, err := reader.Read()
	if err != nil {
		return err
	}

	fields := make([]reflect.StructField, len(heading))
	for i, field_name := range heading {
		field, ok := element_struct_type.FieldByName(field_name)
		if !ok {
			panic(strconv.Quote(field_name) + " not a field of " + element_struct_type.Name())
		}
		fields[i] = field
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		if len(line) > len(heading) {
			panic(line)
		}

		ref := reflect.New(element_struct_type)
		element := ref.Elem()

		for i, value := range line {
			if value == "" {
				continue
			}
			field := fields[i]
			fvalue := element.FieldByName(field.Name)

			switch field.Type.Kind() {
			case reflect.String:
				fvalue.SetString(value)
			case reflect.Int:
				value, err := strconv.ParseInt(value, 10, 0)
				if err != nil {
					panic(err)
				}
				fvalue.SetInt(value)
			case reflect.Bool:
				bvalue := false
				if len(value) > 0 {
					bvalue = true
				}
				fvalue.SetBool(bvalue)
			default:
				panic(field.Type.Kind())
			}
		}
		output_value.Set(reflect.Append(output_value, ref))
	}
	return nil
}
