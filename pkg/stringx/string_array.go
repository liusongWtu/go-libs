package stringx

import "encoding/json"

// string转换成[]string
// s=`["test","hi"]`
func StringToStringArray(s string) ([]string, error) {
	if len(s) < 3 {
		return []string{}, nil
	}
	var result []string
	err := json.Unmarshal(StringToBytes(s), &result)
	if err != nil {
		return []string{}, err
	}
	return result, nil
}

func StringArrayToString(s []string) (string, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	data, err := json.Marshal(s)
	if err != nil {
		return "[]", err
	}
	return BytesToString(data), nil
}
