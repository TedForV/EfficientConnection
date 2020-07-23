package autoform

import (
	jsoniter "github.com/json-iterator/go"
)

// Creator define the actions for db_creator,model_creator and view_creator
type Creator interface {
	Validate(form FormInfo) (bool, error)
	GenerateScript(form FormInfo) (string, error)
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Generate is the main entry for the new form & flow
func Generate(formJSONStr string) error {
	// info, err := analyse(formJSONStr)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// analyse json string to original model for later processing
func analyse(formJSONStr string) (FormInfo, error) {
	var info FormInfo
	err := json.UnmarshalFromString(formJSONStr, &info)
	return info, err
}

func validate(form FormInfo, creators ...Creator) (bool, error) {
	var result bool
	var err error
	for _, v := range creators {
		if v == nil {
			continue
		}
		result, err = v.Validate(form)
		if !result {
			return result, err
		}
	}
	return true, nil
}





