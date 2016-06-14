package goutil

import (
	"github.com/ncbray/compilerutil/writer"
	"io"
	"path/filepath"
	"reflect"
	"strconv"
)

func isZero(element reflect.Value) bool {
	switch element.Kind() {
	case reflect.String, reflect.Slice:
		return element.Len() == 0
	case reflect.Bool:
		return !element.Bool()
	case reflect.Map, reflect.Interface:
		return element.IsNil()
	default:
		panic(element.Kind())
	}
}

func isConcrete(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Struct, reflect.String, reflect.Bool:
		return true
	case reflect.Interface:
		return false
	case reflect.Ptr:
		return isConcrete(t.Elem())
	default:
		panic(t.Kind())
	}
}

func dumpTypeName(t reflect.Type, out *writer.TabbedWriter) {
	switch t.Kind() {
	case reflect.Struct:
		_, pkg := filepath.Split(t.PkgPath())
		out.WriteString(pkg)
		out.WriteString(".")
		out.WriteString(t.Name())
	case reflect.Ptr:
		out.WriteString("*")
		dumpTypeName(t.Elem(), out)
	case reflect.String:
		out.WriteString(t.Name())
	default:
		panic(t.Kind())
	}
}

func dumpChild(element reflect.Value, can_infer_type bool, out *writer.TabbedWriter) {
	switch element.Kind() {
	case reflect.Ptr:
		v := element.Elem()
		t := v.Type()
		if !can_infer_type {
			out.WriteString("&")
			dumpTypeName(t, out)
		}
		out.WriteString("{")
		out.EndOfLine()

		out.Indent()
		for i := 0; i < v.NumField(); i++ {
			fieldInfo := t.Field(i)
			fieldValue := v.Field(i)
			if !isZero(fieldValue) {
				out.WriteString(fieldInfo.Name)
				out.WriteString(": ")
				dumpChild(fieldValue, false, out)
				out.WriteString(",")
				out.EndOfLine()
			}
		}
		out.Dedent()

		out.WriteString("}")
	case reflect.Slice:
		out.WriteString("[]")
		child_type := element.Type().Elem()
		dumpTypeName(child_type, out)
		out.WriteString("{")
		out.EndOfLine()

		can_infer_child_type := isConcrete(child_type)

		out.Indent()
		for i := 0; i < element.Len(); i++ {
			dumpChild(element.Index(i), can_infer_child_type, out)
			out.WriteString(",")
			out.EndOfLine()
		}
		out.Dedent()
		out.WriteString("}")
	case reflect.String:
		v := element.Interface()
		s, _ := v.(string)
		out.WriteString(strconv.Quote(s))
	case reflect.Float64:
		v := element.Float()
		out.WriteString(strconv.FormatFloat(v, 'g', -1, 64))
	case reflect.Bool:
		v := element.Interface()
		b, _ := v.(bool)
		out.WriteString(strconv.FormatBool(b))
	case reflect.Interface:
		dumpChild(element.Elem(), false, out)
	default:
		panic(element.Kind())
	}
}

func DumpTree(tree interface{}, out io.Writer) {
	tabbed := writer.MakeTabbedWriter("\t", out)
	dumpChild(reflect.ValueOf(tree), false, tabbed)
}
