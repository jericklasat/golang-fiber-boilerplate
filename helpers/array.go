package helpers

import "reflect"

func InArray(val interface{}, array interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
			s := reflect.ValueOf(array)

			for i := 0; i < s.Len(); i++ {
					if reflect.DeepEqual(val, s.Index(i).Interface()) {
							exists = true
							return
					}
			}
	}

	return
}