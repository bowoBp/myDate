package mapper

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
	"strings"
)

func OverrideTranslation(engine *validator.Validate, trans ut.Translator) error {
	err := engine.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} tidak boleh kosong", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})
	if err != nil {
		log.Println(err)
		return err
	}

	err = engine.RegisterTranslation("excludesall", trans, func(ut ut.Translator) error {
		return ut.Add("excludesall", "{0} tidak boleh memiliki {1}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		charStr := fe.Param()
		charStr = strings.ReplaceAll(charStr, "", ",")
		charStr = strings.ReplaceAll(charStr, " ", "<spasi>")
		t, _ := ut.T("excludesall", fe.Field(), charStr)
		return t
	})
	if err != nil {
		log.Println(err)
		return err
	}

	err = engine.RegisterTranslation("enum", trans, func(ut ut.Translator) error {
		return ut.Add("enum", "pada {0} input yang diperbolehkan adalah {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		enums := fe.Param()
		t, _ := ut.T("enum", fe.Field(), enums)
		return t
	})
	if err != nil {
		log.Println(err)
		return err
	}

	err = engine.RegisterTranslation("monthyearformat", trans, func(ut ut.Translator) error {
		return ut.Add("monthyearformat", "{0} format harus mm-yyyy", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("monthyearformat", fe.Field())
		return t
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
