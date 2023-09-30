package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	//Ma loi 404, 500, 403,...
	StatusCode int `json:"status_code"`
	//Loi goc
	RootErr error `json:"-"`
	//bao loi cho client
	Message string `json:"message"`
	//bao loi cho dev tu RootErr
	Log string `json:"log"`
	//custom thong bao loi nhieu ngon ngu nhu tieng viet, tieng anh
	Key string `json:"key"`
}

func NewFullErrorResponse(statusCode int, rootErr error, message string, log string, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    rootErr,
		Message:    message,
		Log:        log,
		Key:        key,
	}
}

// define error
func NewErrorResponse(rootErr error, message string, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest, //400
		RootErr:    rootErr,
		Message:    message,
		Key:        key,
	}
}

// authorize
func NewUnauthorizedError(rootErr error, message string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized, //401
		RootErr:    rootErr,
		Message:    message,
		Key:        key,
	}
}

// Error(Error(Error...))) 1 chuong trinh co the co nhieu loi
// RootError() lay ra loi goc
// day la design pattern decorator
func (e *AppError) RootError() error {
	//lay root error co 2 truong hop
	// 1 la boc 1 con tro error khac, neu nhu true thi se goi lai ham RootError de lay ra loi goc
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr

}
func (e *AppError) Error() string {
	return e.RootErr.Error() //neu nhu muon xay custom error thi chi can implement ham Error() string
}
func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewCustomError(errors.New(msg), msg, key)

}

// loi db
func ErrorDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with db", err.Error(), "DB_ERROR")
}

// loi validate
func ErrorValidate(err error) *AppError {
	return NewFullErrorResponse(http.StatusBadRequest, err, "Invalid input", err.Error(), "INVALID_INPUT")
}

// loi Internal- loi lien quan den server nhu runtime, panic, ...
func ErrorInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong", err.Error(), "INTERNAL_ERROR")
}

// loi khong list duoc
func ErrorCannotList(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToUpper(entity)),
		fmt.Sprintf("CANNOT_LIST_%s", strings.ToUpper(entity)),
	)
}

// loi delete
func ErrorCannotDelete(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToUpper(entity)),
		fmt.Sprintf("CANNOT_DELETE_%s", strings.ToUpper(entity)),
	)

}

// loi update
func ErrorCannotUpdate(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToUpper(entity)),
		fmt.Sprintf("CANNOT_UPDATE_%s", strings.ToUpper(entity)),
	)

}

// loi create
func ErrorCannotCreate(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot create %s", strings.ToUpper(entity)),
		fmt.Sprintf("CANNOT_CREATE_%s", strings.ToUpper(entity)),
	)

}

// loi khong tim thay
func ErrorCannotGet(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot find %s", strings.ToUpper(entity)),
		fmt.Sprintf("CANNOT_FIND_%s", strings.ToUpper(entity)),
	)
}

// loi khong tim thay RecordNotFound
var RecordNotFound = errors.New("record not found")
