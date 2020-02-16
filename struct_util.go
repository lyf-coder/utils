/*
 * Copyright © 2020 - present. liyongfei <liyongfei@walktotop.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package utils

import (
	"reflect"
)

// CopyStruct - copy val from src struct to target struct
// target is &
func CopyStruct(src interface{}, target interface{}) {
	// get target Type and Value
	targetType := reflect.TypeOf(target).Elem()
	targetValue := reflect.ValueOf(target).Elem()
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)

	// get src val set to target field
	for i := 0; i < targetType.NumField(); i++ {
		// get target filedName
		fieldName := targetType.Field(i).Name
		_, haveName := srcType.FieldByName(fieldName)
		if haveName {
			// get src fieldName val
			fieldNameVal := srcValue.FieldByName(fieldName)
			// set val to target
			targetValue.FieldByName(fieldName).Set(fieldNameVal)
		}

	}
}
