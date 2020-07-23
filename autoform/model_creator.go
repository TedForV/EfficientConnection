package autoform

import (
	"bytes"
	"fmt"
	"github.com/TedForV/EfficientConnection/commonfunc"
	"github.com/pkg/errors"
	"strings"
)

const (
	packageScript = `package formmodel
	
	`
	importPart = `import (
		github.com/shopspring/decimal
	)
	
	`
	importDecimalScript  = "github.com/shopspring/decimal"
	structScriptTemplate = `type %s struct {
		%s
	}`
)

type modelCreator struct {
}

func (mc *modelCreator) Validate(form FormInfo) (bool, error) {
	columns := form.Columns

	if columns == nil || len(columns) == 0 {
		return false, errors.New("无数据列定义数据")
	}

	formColumns := make([]string, 0, len(columns))
	existedColumn := make(map[string]interface{})

	for _, v := range columns {
		name := strings.ToUpper(strings.Trim(v.Name, " "))
		if _, ok := existedColumn[name]; ok {
			return false, errors.New(fmt.Sprintf("列（%s）重复", name))
		}
		if !isFormTypeExisted(v.ID) {
			return false, errors.New(fmt.Sprintf("列（%s）的类型（%d）错误", v.Name, v.ColumnType.ID))
		}
		existedColumn[name] = nil
		formColumns = append(formColumns, name)
	}

	result, notInKeys := commonfunc.Contains(formColumns, true, "id", "name", "need_flow", "flow_id", "is_flow_end")

	if result {
		return true, nil
	}

	return false, errors.New(fmt.Sprintf("缺失必要的列：%s", strings.Join(notInKeys, ",")))

}

func (mc *modelCreator) GenerateScript(form FormInfo, hasDecimalType bool) (string, error) {
	var fullBuf bytes.Buffer
	var columnBuf bytes.Buffer
	fullBuf.WriteString(packageScript)
	if hasDecimalType {
		fullBuf.WriteString(importPart)
	}

	for _, v := range form.Columns {
		script, err := generateColumnGoScript(v)
		if err != nil {
			return "", err
		}
		columnBuf.WriteString(script)
		columnBuf.WriteString("  ")
	}
	fullBuf.WriteString(fmt.Sprintf(structScriptTemplate, form.FormCode, columnBuf.String()))
	return fullBuf.String(), nil
}
