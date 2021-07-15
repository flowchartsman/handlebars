package handlebars

import "reflect"

type virtualMethod func(ctx reflect.Value) (val reflect.Value, hasMethod bool)

func getVirtualMethod(name string) virtualMethod {
	switch name {
	case "length":
		return vmethodLength
	}
	return nil
}

func vmethodLength(ctx reflect.Value) (reflect.Value, bool) {
	switch ctx.Kind() {
	case reflect.Slice, reflect.Array, reflect.String, reflect.Map:
		return reflect.ValueOf(ctx.Len()), true
	}
	return zero, false
}
