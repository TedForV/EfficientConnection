package autoform

import (
	jsoniter "github.com/json-iterator/go"
)

const (
	PRIMARY_KEY_SCRIPT = "primary key (id) ";
)

type dbCreator struct{}

func (dc *dbCreator) Validate(jsonObj jsoniter.Any) (bool, error) {
	return true, nil
}

func (dc *dbCreator) GenerateScript(jsoniter.Any) (string, error) {
	return "", nil
}
