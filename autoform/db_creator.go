package autoform

import (
	jsoniter "github.com/json-iterator/go"
)

type dbCreator struct{}

func (dc *dbCreator) Validate(jsonObj jsoniter.Any) (bool, error) {
	return true, nil
}

func (dc *dbCreator) GenerateScript(jsoniter.Any) (string, error) {
	return "", nil
}
