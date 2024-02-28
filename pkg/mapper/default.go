package mapper

import "gorm.io/driver/postgres"

type (
	Mappers struct {
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

func (m Mappers) ReplaceLabelErr(template error, params ...string) error {
	//TODO implement me
	panic("implement me")
}

func (m Mappers) ErrorIs(template error, targer error) bool {
	//TODO implement me
	panic("implement me")
}

func (m Mappers) TranslateSQLErr(mySqlErr *postgres.ErrMessage, methodName string) error {
	//TODO implement me
	panic("implement me")
}

func (m Mappers) CompareSliceOfErr(errs []error, target error) bool {
	//TODO implement me
	panic("implement me")
}

func (m Mappers) EvaluateError(label string, expectedErr []error, err error) error {
	//TODO implement me
	panic("implement me")
}

func Default() MapperUtility {
	return &Mappers{}
}
