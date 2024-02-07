package checkers

import (
	"log"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
)

func DobVal(fl validator.FieldLevel) bool {
	switch v := fl.Field(); v.Kind() {
	case reflect.String:
		dob := v.String()
		if _, err := time.Parse(time.DateOnly, dob); err != nil {
			log.Println(err)
			return false
		}
	}
	return true
}
