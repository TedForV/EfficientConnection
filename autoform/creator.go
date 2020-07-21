package autoform

import (
	jsoniter "github.com/json-iterator/go"
)

// Creator define the actions for db_creator,model_creator and view_creator
type Creator interface {
	Validate(jsoniter.Any) (bool, error)
	GenerateScript(jsoniter.Any) (string, error)
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Analyse json string to original model for later processing
func Analyse(jsonStr string) {

}

