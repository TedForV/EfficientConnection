package autoform

const (
	// PrimaryKeyScript is primary key script
	PrimaryKeyScript = "primary key (id) "
)

type dbCreator struct{}

func (dc *dbCreator) Validate(form FormInfo) (bool, error) {
	return true, nil
}

func (dc *dbCreator) GenerateScript(form FormInfo) (string, error) {
	return "", nil
}
