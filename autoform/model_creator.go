package autoform

import (
	jsoniter "github.com/json-iterator/go"
)

type modelCreator struct {
}

func (mc *modelCreator) Validate(jsonObj jsoniter.Any) (bool, error) {
	 jsonObj.Keys()
}
