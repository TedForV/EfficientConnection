package autoform

type viewCreator struct{}

func (vc *viewCreator) Validate(form FormInfo) (bool, error) {
	return true, nil
}

func (vc *viewCreator) GenerateScript(form FormInfo) (string, error) {
	return "", nil
}
