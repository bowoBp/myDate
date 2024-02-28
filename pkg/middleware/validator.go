package middleware

import (
	"encoding/json"
	"errors"
	"github.com/bowoBp/myDate/pkg/mapper"
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations_en "github.com/go-playground/validator/v10/translations/en"
	"log"
	"reflect"
	"strings"
)

type (
	Enigma struct {
		engine *validator.Validate
		trans  ut.Translator
	}
	EnigmaUtility interface {
		Validate(c *gin.Context, payload any) map[string][]string
		BindAndValidate(c *gin.Context, payload any) map[string][]string
	}
)

func NewEnigma() EnigmaUtility {
	engine := validator.New()
	engine.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	err := engine.RegisterValidation("nonzero-gt", validateNonZeroAndGreaterThan)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = engine.RegisterValidation("enum", validateEnumValue)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = engine.RegisterValidation("monthyearformat", validateMonthYearFormat)
	if err != nil {
		log.Println(err)
		return nil
	}

	en := en2.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	err = translations_en.RegisterDefaultTranslations(engine, trans)
	if err != nil {
		log.Println(err)
		log.Println("new validator: ", err)
		return nil
	}

	err = mapper.OverrideTranslation(engine, trans)
	if err != nil {
		log.Println(err)
		log.Println("new validator: ", err)
		return nil
	}

	return &Enigma{
		engine: engine,
		trans:  trans,
	}
}

func (v Enigma) Validate(c *gin.Context, payload any) map[string][]string {
	errs := make(map[string][]string)
	err := v.engine.StructCtx(c.Request.Context(), payload)
	if err != nil {
		log.Println(err)
		var errVals validator.ValidationErrors
		if errors.As(err, &errVals) {
			for i, _ := range errVals {
				errs[errVals[i].Field()] = []string{errVals[i].Translate(v.trans)}
			}

			return errs
		}
		return errs
	}

	return nil
}

func (v Enigma) BindAndValidate(c *gin.Context, payload any) map[string][]string {
	err := c.Bind(payload)
	if err != nil {
		log.Println(err)
		var errJSON *json.UnmarshalTypeError
		if errors.As(err, &errJSON) {
			field := errJSON.Field
			errVal := errJSON.Error()
			return map[string][]string{
				field: {errVal},
			}
		}
		return map[string][]string{
			"error": {err.Error()},
		}
	}
	return v.Validate(c, payload)
}
