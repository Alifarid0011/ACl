package validation

type DefaultValidator struct{}

func (v *DefaultValidator) ValidateStruct(obj interface{}) error {
	return Validator().Struct(obj)
}

func (v *DefaultValidator) Engine() interface{} {
	return Validator()
}
