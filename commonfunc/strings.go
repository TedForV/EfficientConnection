package commonfunc

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"unicode"
)

// Contains for string contain keys
func Contains(source []string, caseInsensitive bool, keys ...string) (bool, []string) {
	if keys == nil || len(keys) == 0 {
		return false, keys
	} else if len(keys) == 1 {
		result := containsKey(source, caseInsensitive, keys[0])
		return result, nil
	} else {
		return containsKeys(source, caseInsensitive, keys)
	}
}

var underscore rune = rune('_')

// ConvertToStructPropertyName convert name fit the struct rules
func ConvertToStructPropertyName(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("字符串为空")
	}
	words := []rune(name)
	result := make([]rune, 0, len(words)*2)
	nextCapital := false
	for _, v := range words {
		if v == underscore {
			nextCapital = true
			continue
		}

		if !unicode.IsLetter(v) {
			return "", errors.New(fmt.Sprintf("包含非英文字符：%s", string(v)))
		}

		if nextCapital {
			result = append(result, unicode.ToUpper(v))
			nextCapital = false
		} else {
			result = append(result, v)
		}
	}
	if !unicode.IsUpper(result[0]) {
		result[0] = unicode.ToUpper(result[0])
	}
	return string(result), nil
}

// ConvertToMysqlColumnName convert name fit the mysql table column rule
func ConvertToMysqlColumnName(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("字符串为空")
	}
	words := []rune(name)
	result := make([]rune, 0, len(words)*2)
	for _, v := range words {
		if v == underscore {
			result = append(result, v)
			continue
		}
		if !unicode.IsLetter(v) {
			return "", errors.New(fmt.Sprintf("包含非英文字符：%s", string(v)))
		}
		if unicode.IsUpper(v) {
			if len(result) != 0 && result[len(result)-1] != underscore {
				result = append(result, underscore)
			}
			result = append(result, unicode.ToLower(v))
		} else {
			result = append(result, v)
		}
	}
	if result[0] == underscore {
		result = result[1:]
	}
	return string(result), nil
}

// ContainsKeys judge keys for contain method
func containsKeys(source []string, caseInsensitive bool, keys []string) (bool, []string) {
	var tempMap = make(map[string]interface{}, len(source))
	result := make([]string, 0, len(keys))
	if caseInsensitive {
		for _, v := range source {
			tempV := strings.ToUpper(strings.Trim(v, " "))
			if _, ok := tempMap[tempV]; !ok {
				tempMap[tempV] = nil
			}
		}
		for _, v := range keys {
			if _, ok := tempMap[strings.ToUpper(strings.Trim(v, " "))]; !ok {
				result = append(result, v)
			}
		}
		return len(result) == 0, result
	}

	for _, v := range source {
		tempV := strings.Trim(v, " ")
		if _, ok := tempMap[tempV]; !ok {
			tempMap[tempV] = nil
		}
	}
	for _, v := range keys {
		if _, ok := tempMap[strings.Trim(v, " ")]; !ok {
			result = append(result, v)
		}
	}
	return len(result) == 0, result
}

// ContainsKey judge only one key for contain method
func containsKey(source []string, caseInsensitive bool, key string) bool {
	if caseInsensitive {
		key = strings.ToUpper(strings.Trim(key, " "))
		for _, v := range source {
			if strings.ToUpper(strings.Trim(v, " ")) == key {
				return true
			}
		}
	} else {
		key := strings.Trim(key, " ")
		for _, v := range source {
			if strings.Trim(v, " ") == key {
				return true
			}
		}
	}
	return false
}
