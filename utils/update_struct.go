package utils

import "reflect"

func UpdateStruct[T any, U any](target *T, updateData U) error {
	vTarget := reflect.ValueOf(target).Elem()
	vUpdate := reflect.ValueOf(updateData)
	for i := 0; i < vUpdate.NumField(); i++ {
		field := vUpdate.Field(i)
		if field.IsValid() && !field.IsZero() {
			// Look for the same field name in the target struct
			targetField := vTarget.FieldByName(vUpdate.Type().Field(i).Name)
			if targetField.IsValid() && targetField.CanSet() {
				targetField.Set(field)
			}
		}
	}
	return nil
}
