package autoform

import (
	"fmt"
	"strconv"
)

// FormType is the interface define the type actions
type FormType interface {
	GenerateGolangPropertyScript(name string) string
	GenerateDbColumnScript(name string) string
}

type FormInfo struct {
}

// FormCommonType is common type supported in form
type FormCommonType struct {
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
	supportedFormTypes = make([]FormCommonType, 0, 10)
	supportedFormTypes = append(supportedFormTypes, FormCommonType{})
}

// GetSupportedCommonType returns all type for form
func GetSupportedCommonType() ([]FormCommonType, error) {
	return []FormCommonType{
		createIntType(),
		createDecimalType(),
		createFloatType(),
		createStringType(),
		createBoolType(),
		createAttachmentType()}, nil
}

// GetSupportedEnumType returns all enum type for form
func GetSupportedEnumType() ([]FormEnumType, error) {
	return nil, nil
}

// GetSupportedFormType returns all form type for form
func GetSupportedFormType() {

}

func createIntType() FormCommonType {
	return FormCommonType{
		Name:         "int",
		IsNeedLength: false,
		Length:       0,
		GolangPropertyTemplate: "%s		int32	`gorm:\"column:%s\"`",
		DbColumnTemplate: "%s		bigint not null,",
	}
}

func createDecimalType() FormCommonType {
	return FormCommonType{
		Name:         "decimal",
		IsNeedLength: false,
		Length:       0,
		GolangPropertyTemplate: "%s 	decimal.Decimal	`gorm:\"column:%s\"`",
		DbColumnTemplate: "%s		decimal not null,",
	}
}

func createFloatType() FormCommonType {
	return FormCommonType{
		Name:         "float",
		IsNeedLength: false,
		Length:       0,
		GolangPropertyTemplate: "%s		float	`gorm:\"column:%s\"`",
		DbColumnTemplate: "%s 	float not null,",
	}
}

func createStringType() FormCommonType {
	return FormCommonType{
		Name:         "string",
		IsNeedLength: true,
		Length:       0,
		GolangPropertyTemplate: "%s		string `gorm:\"column:%s\"`",
		DbColumnTemplate: "%s		varchar(%s) not null,",
	}
}

func createBoolType() FormCommonType {
	return FormCommonType{
		Name:         "bool",
		IsNeedLength: false,
		Length:       0,
		GolangPropertyTemplate: "%s		bool `gorm:\"column:%s\"`",
		DbColumnTemplate: "%s		tinyint not null,",
	}
}

func createAttachmentType() FormCommonType {
	return createIntType()
}
