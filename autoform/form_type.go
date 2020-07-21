package autoform

// FormType is the interface define the type actions
type FormType interface {
	GenerateGolangPropertyScript()
	GenerateDbColumnScript()
}

// FormCommonType is common type supported in form
type FormCommonType struct {
	Name                   string
	IsNeedLength           bool
	Length                 int
	GolangPropertyTemplate string
	DbColumnTemplate       string
}

// FormRelateType define a relate table type
type FormRelateType struct {
	Name      string
	Selection []FormRelateTypeSelection
}

// FormRelateTypeSelection for selections for relate content
type FormRelateTypeSelection struct {
	id          int
	Name        string
	Description string
}

var supportedFormTypes []FormCommonType

func init() {
	supportedFormTypes = make([]FormCommonType, 0, 10)
	supportedFormTypes = append(supportedFormTypes, FormCommonType{})
}

// GetSupportedFormType returns all type for form
func GetSupportedFormType() {

}

func createIntType() FormCommonType {
	return FormCommonType{
		Name:         "int",
		IsNeedLength: false,
		Length:       0,
		GolangPropertyTemplate: "%s		int32",
		DbColumnTemplate: "%s		bigint not null,",
	}
}

//TODO: missing golang property template,missing decimal import
func createDecimalType() FormCommonType {
	return FormCommonType{
		Name:                   "decimal",
		IsNeedLength:           false,
		Length:                 0,
		GolangPropertyTemplate: "",
		DbColumnTemplate: "%s		decimal not null,",
	}
}
