package validator

import (
	"reflect"
	"sync"

	v10 "github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Once     sync.Once
	Validate *v10.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{}
}

func (v *CustomValidator) ValidateStruct(obj any) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyInit()
		if err := v.Validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *CustomValidator) Engine() any {
	v.lazyInit()

	return v.Validate
}

func (v *CustomValidator) lazyInit() {
	v.Once.Do(func() {
		v.Validate = v10.New()
		v.Validate.SetTagName("binding")
	})
}

func kindOfData(data any) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}

	return valueType
}
