package log

import (
	"reflect"
	"strconv"
	"regexp"
	"fmt"
)

var substitutionsRegex = regexp.MustCompile(":\\w[\\w]+[$,\\. ]");

func Substitute(original string, callback func(string, string, string) string) string {
	return substitutionsRegex.ReplaceAllStringFunc(original, func(in string) string {
		offset := 0
		separator := ""
		if last := in[len(in)-1:]; last == "," || last == "." || last == " " {
			offset = 1
			separator = last
		}

		return callback(in, in[1: len(in) - offset], separator)
	});
}

func SubstituteTypeHelper(value interface{}, nilf func() string, strf func(string) string, numf func(string) string) string {

	if value == nil {
		return nilf()
	}

	v := reflect.ValueOf(value)
	fmt.Printf("> %s %v\n", v.Kind().String(), value)
	switch v.Kind() {
	case reflect.Func:
		result := v.Call(nil)[0]
		return SubstituteTypeHelper(
			result.Interface(),
			nilf,
			strf,
			numf,
		)
	case reflect.Ptr:
		return "*ptr"
	case reflect.String:
		return strf(value.(string))
	case reflect.Int:
		return numf(strconv.FormatInt(int64(value.(int)), 10))
	case reflect.Int8:
		return numf(strconv.FormatInt(int64(value.(int8)), 10))
	case reflect.Int16:
		return numf(strconv.FormatInt(int64(value.(int16)), 10))
	case reflect.Int32:
		return numf(strconv.FormatInt(int64(value.(int32)), 10))
	case reflect.Int64:
		return numf(strconv.FormatInt(value.(int64), 10))
	case reflect.Float32:
		return numf(strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32))
	case reflect.Float64:
		return numf(strconv.FormatFloat(value.(float64), 'f', -1, 64))
	}

	return " oops! "
}