package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	valid "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

func (v *ValidError) Error() string {
	return v.Message
}

type ValidErrors []*ValidError

func (v *ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func bindAndValid(c *gin.Context, err error) (bool, ValidErrors) {
	var errs ValidErrors
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(valid.ValidationErrors)
		if !ok {
			return false, errs
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}

func BindAndValid(c *gin.Context, v any) (bool, ValidErrors) {
	err := c.ShouldBindJSON(v)
	if err != nil && err.Error() == "EOF" {
		err = c.ShouldBind(v)
	}
	return bindAndValid(c, err)
}

func BindAndValidHeader(c *gin.Context, v any) (bool, ValidErrors) {
	err := c.ShouldBindHeader(v)
	return bindAndValid(c, err)
}
