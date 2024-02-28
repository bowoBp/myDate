package mapper

import "gorm.io/driver/postgres"

type (
	Mapper struct {
	}

	MapperUtility interface {
		ReplaceLabelErr(template error, params ...string) error
		ErrorIs(template error, targer error) bool
		TranslateSQLErr(mySqlErr *postgres.ErrMessage, methodName string) error
		CompareSliceOfErr(errs []error, target error) bool
		EvaluateError(label string, expectedErr []error, err error) error
		ParseServiceDurationFormat(d string) (string, error)
		SortingByStructField(vals interface{}, fieldName string, sorting SortingDirection) interface{}
		UniqueByStructField(vals interface{}, fieldName string) interface{}
	}
)
