// error.go contains the generic domain errors, specific application errors may reside in the usecase.

package domain

import (
	"database/sql"
	"fmt"
	"reflect"
)

func getTypeName(i interface{}) string {
	if t := reflect.TypeOf(i); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

//func main() {
//if err := NotFound(sql.ErrNoRows, new(User)); err != nil {
//fmt.Println(err)

//var nerr *NotFoundError
//if errors.As(err, &nerr) {
//fmt.Println(nerr.Name, nerr.Err)
//}
//}
//}

type NotFoundError struct {
	Err  error
	Name string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", getTypeName(e), e.Name)
}

// NotFound checks if the error is not found and returns the NotFoundError.
func NotFound(err error, i interface{}) error {
	if err == sql.ErrNoRows {
		return &NotFoundError{Err: err, Name: getTypeName(i)}
	}
	return err
}
