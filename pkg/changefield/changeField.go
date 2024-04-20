package changefield

import (
	"fmt"
	"reflect"
)

func ChangeField(object interface{}, fieldName string, newValue interface{}) error {
	val := reflect.ValueOf(object)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("obj должен быть структурой")
	}

	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("поле %s не найдено", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("поле %s не может быть изменено", fieldName)
	}

	newVal := reflect.ValueOf(newValue)
	if newVal.Type() != field.Type() {
		return fmt.Errorf("типы не совпадают: поле %s имеет тип %s, а значение имеет тип %s",
			fieldName,
			field.Type(),
			newVal.Type(),
		)
	}

	field.Set(newVal)

	return nil
}
