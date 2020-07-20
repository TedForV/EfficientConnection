package commonfunc

import (
	"strings"
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
