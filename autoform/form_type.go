package autoform

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

// FormType is the interface define the type actions
type FormType interface {
	GenerateGolangPropertyScript(name string) string
	GenerateDbColumnScript(name string) string
}

// FormInfo is json model
type FormInfo struct {
	FormName    string       `json:"formName"`
	FormCode    string       `json:"formCode"`
	DisplayText string       `json:"displayText"`
	NeedFlow    bool         `json:"needFlow"`
	Columns     []FormColumn `json:"columns"`
	Layout      FormLayout   `json:"layout"`
}

// FormColumn define the column definition in FormInfo
type FormColumn struct {
	ID          int            `jsong:"id"`
	Name        string         `json:"name"`
	DisplayText string         `json:"displayText"`
	ColumnType  FormColumnType `json:"type"`
}

// FormColumnType define the column type selection in FormInfo
type FormColumnType struct {
	ID  int `json:"id"`
	Len int `json:"len"`
}

// FormLayout define the layout for form columns
type FormLayout struct {
	Column  int     `json:"column"`
	Details [][]int `json:"details"`
}

// FormCommonType is common type supported in form
type FormCommonType struct {
	ID                     int
	Name                   string
	IsNeedLength           bool
	Length                 int
	GolangPropertyTemplate string
	DbColumnTemplate       string
}

// GenerateGolangPropertyScript generate golang script
func (fct *FormCommonType) GenerateGolangPropertyScript(name string) string {
	return fmt.Sprintf(fct.GolangPropertyTemplate, name)
}

// GenerateDbColumnScript generate db column script
func (fct *FormCommonType) GenerateDbColumnScript(name string) string {
	if fct.IsNeedLength {
		return fmt.Sprintf(fct.DbColumnTemplate, name)
	}
	return fmt.Sprintf(fct.DbColumnTemplate, name, strconv.Itoa(fct.Length))
}

// FormEnumType is enum type supported in form
type FormEnumType struct {
	Name  string
	Value []KeyValuePair
}

// KeyValuePair is key:value pair
type KeyValuePair struct {
	Key   int
	Value string
}

// FormRelateType define a relate table type
type FormRelateType struct {
	ID        int
	Name      string
	Selection []FormRelateTypeSelection
}

// FormRelateTypeSelection for selections for relate content
type FormRelateTypeSelection struct {
	ID          int
	Name        string
	Description string
}

var supportedFormTypes []FormCommonType

func init() {
	i := 0
	i = initSupportedCommonType(i)

}

var supportedCommonTypes map[int]FormCommonType
var supportedCommonTypesArr []FormCommonType

var decimailID int

// GetSupportedCommonType returns all type for form
func GetSupportedCommonType() ([]FormCommonType, error) {
	if supportedCommonTypesArr == nil {
		supportedCommonTypesArr = make([]FormCommonType, 0, len(supportedCommonTypes))
		for _, v := range supportedCommonTypes {
			supportedCommonTypesArr = append(supportedCommonTypesArr, v)
		}
	}

	return supportedCommonTypesArr, nil
}

func isFormTypeExisted(id int) bool {
	if _, ok := supportedCommonTypes[id]; ok {
		return true
	}
	return false
}

func generateColumnGoScript(column FormColumn) (string, error) {
	if v, ok := supportedCommonTypes[column.ColumnType.ID]; ok {
		return v.GenerateGolangPropertyScript(column.Name), nil
	}
	return "", errors.New("数据类型不支持")
}

// GetSupportedEnumType returns all enum type for form
func GetSupportedEnumType() ([]FormEnumType, error) {
	return nil, nil
}

// GetSupportedFormType returns all form type for form
func GetSupportedFormType() {

}

func initSupportedCommonType(i int) int {
	supportedCommonTypes = make(map[int]FormCommonType)
	i++
	supportedCommonTypes[i] = createIntType(i)
	i++
	supportedCommonTypes[i] = createDecimalType(i)
	decimailID = i
	i++
	supportedCommonTypes[i] = createFloatType(i)
	i++
	supportedCommonTypes[i] = createStringType(i)
	i++
	supportedCommonTypes[i] = createBoolType(i)
	i++
	supportedCommonTypes[i] = createAttachmentType(i)
	return i
}

func createIntType(id int) FormCommonType {
	return FormCommonType{
		ID:           id,
		Name:         "int",
		IsNeedLength: false,
		Length:       0,
		GolangPropertyTemplate: "%s		int32	`gorm:\"column:%s\"`",
		DbColumnTemplate: "%s		bigint not null,",
	}
}

func createDecimalType(id int) FormCommonType {
	return FormCommonType{
		ID:           id,
		Name:         "decimal",
		IsNeedLength: false,
		Length:       0,
		GolangPropertyTemplate: "%s 	decimal.Decimal	`gorm:\"column:%s\"`",
		DbColumnTemplate: "%s		decimal not null,",
	}
}

func createFloatType(id int) FormCommonType {
	return FormCommonType{
		ID:           id,
		Name:         "float",
		IsNeedLength: false,
		Length:       0,
		GolangPropertyTemplate: "%s		float	`gorm:\"column:%s\"`",
		DbColumnTemplate: "%s 	float not null,",
	}
}

func createStringType(id int) FormCommonType {
	return FormCommonType{
		ID:           id,
		Name:         "string",
		IsNeedLength: true,
		Length:       0,
		GolangPropertyTemplate: "%s		string `gorm:\"column:%s\"`",
		DbColumnTemplate: "%s		varchar(%s) not null,",
	}
}

func createBoolType(id int) FormCommonType {
	return FormCommonType{
		ID:           id,
		Name:         "bool",
		IsNeedLength: false,
		Length:       0,
		GolangPropertyTemplate: "%s		bool `gorm:\"column:%s\"`",
		DbColumnTemplate: "%s		tinyint not null,",
	}
}

func createAttachmentType(id int) FormCommonType {
	return createIntType(id)
}
