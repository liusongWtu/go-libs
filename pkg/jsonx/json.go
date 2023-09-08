package jsonx

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"
)

// ToSortedKeyJsonShallow 将json串转为json字符串
// 只支持一层json结构！！！！！
// float64不会丢失精度
func ToSortedKeyJsonShallow(data string, ignoreKeys map[string]struct{}) (string, error) {
	// dataJson, err := json.Marshal(data)
	// if err != nil {
	// 	return "", err
	// }
	var info map[string]any
	dd := json.NewDecoder(bytes.NewReader([]byte(data)))
	dd.UseNumber()
	err := dd.Decode(&info)
	if err != nil {
		return "", err
	}

	keys := make([]string, 0, len(info)-1)
	for key := range info {
		if _, ok := ignoreKeys[key]; ok {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)

	builder := strings.Builder{}
	builder.WriteString("{")
	for _, k := range keys {
		builder.WriteString("\"" + k + "\":")
		var val string
		if v, ok := info[k].(string); ok {
			val = "\"" + v + "\""
		}
		if v, ok := info[k].(json.Number); ok {
			val = string(v)
		}
		builder.WriteString(val)
		builder.WriteString(",")
	}
	targetString := builder.String()
	targetString = targetString[:len(targetString)-1] + "}"

	return targetString, nil
}

func JsonNumberIsString(jsonNumber json.Number, originalString string) bool {
	if len(jsonNumber) == 0 {
		return true
	}

	return strings.Contains(originalString, "\""+string(jsonNumber)+"\"")
}
