package autoform

import (
	"fmt"
	"github.com/TedForV/EfficientConnection/commonfunc"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"strings"
)

var importPart = `import (
	"github.com/shopspring/decimal"
)

`


type modelCreator struct {
}

func (mc *modelCreator) Validate(jsonObj jsoniter.Any) (bool, error) {
	result, notInKeys := commonfunc.Contains(jsonObj.Keys(), true, "id", "name", "flow_id", "is_flow_end")
	if result {
		return true, nil
	}

	return false, errors.New(fmt.Sprintf("缺失必要的列：%s", strings.Join(notInKeys, ",")))

}

func (mc *modelCreator) GenerateScript(jsoniter.Any) (string, error) {
	return "", nil
}
