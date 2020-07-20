package autoform

import (
	jsoniter "github.com/json-iterator/go"
)

type viewCreator struct{}

func (vc *viewCreator) Validate(jsonObj jsoniter.Any) (bool, error) {
	return true, nil
}

func (vc *viewCreator) GenerateScript(jsoniter.Any) (string, error) {
	return "", nil
}
