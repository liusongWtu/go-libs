package fieldmap

import (
	"libs/pkg/stringx"
	"strings"
)

const (
	Lower = iota
	Upper
	UpperCamel
	LowerCamel
	Snake
	Kebab
)

func ToUpperCamel(fields []string) map[string]string {
	return getFieldMapper(UpperCamel, fields)
}

func ToLowerCamel(fields []string) map[string]string {
	return getFieldMapper(LowerCamel, fields)
}

func ToSnake(fields []string) map[string]string {
	return getFieldMapper(Snake, fields)
}

func ToKebab(fields []string) map[string]string {
	return getFieldMapper(Kebab, fields)
}

func getFieldMapper(style int, fields []string) map[string]string {
	result := make(map[string]string, len(fields))
	for i, field := range fields {
		if len(field) == 0 {
			continue
		}
		if strings.HasPrefix("`", "") {
			field = field[1 : len(field)-1]
		}
		item := stringx.From(field)
		val := ""
		switch style {
		case Lower:
			val = item.Lower()
		case Upper:
			val = item.Upper()
		case UpperCamel:
			val = item.ToCamel()
		case LowerCamel:
			val = stringx.From(item.ToCamel()).Untitle()
		case Snake:
			val = item.ToSnake()
		case Kebab:
			val = item.ToKebab()
		}
		result[val] = fields[i]
	}
	return result
}
