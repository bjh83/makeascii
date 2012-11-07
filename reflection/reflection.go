package main

import(
	"os"
	"reflect"
	"strconv"
)

func main() {
	Print("Hello", 1, "\n")
}

func Print(value ... interface{}) {
	for _, i := range value {
		switch value := reflect.ValueOf(i); value.Kind() {
		case reflect.Int:
			os.Stdout.WriteString(strconv.FormatInt(value.Int(), 10))
		case reflect.String:
			os.Stdout.WriteString(value.String())
		}
	}
}
