package Calc

import (
	"fmt"
	"main.go/tuuz/Calc"
	"reflect"
	"strconv"
	"strings"
)

func Chop(s string, character_mask string) string {
	return strings.TrimRight(s, character_mask)
}

func Any2String(any interface{}) string {
	var str string
	switch any.(type) {
	case string:
		str = any.(string)

	case int:
		tmp := any.(int)
		str = Calc.Int2String(tmp)

	case int32:
		tmp := int64(any.(int32))
		str = Calc.Int642String(tmp)

	case int64:
		tmp := any.(int64)
		str = Calc.Int642String(tmp)

	case float64:
		tmp := any.(float64)
		str = Calc.Float642String(tmp)

	case float32:
		tmp := float64(any.(float32))
		str = Calc.Float642String(tmp)

	case nil:
		str = ""

	default:
		fmt.Println(reflect.TypeOf(any))
		str = any.(string)
	}
	return str
}

func String2Int(str string) (int, error) {
	return strconv.Atoi(str)
}

func String2Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func String2Float64(str string) (float64, error) {
	float, err := strconv.ParseFloat(str, 64)
	return float, err
}
