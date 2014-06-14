// Copyright (c) 2014 The go-poblano AUTHORS
//
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package poblano

import (
	"fmt"
	"net/http"
)

type ErrFieldNotSet struct {
	fieldName string
}

func (err *ErrFieldNotSet) Error() string {
	return fmt.Sprintf("Required field '%s' is not set")
}

type ErrHTTP struct {
	*http.Response
}

func (err *ErrHTTP) Error() string {
	return fmt.Sprintf("%v %v -> %v", err.Request.Method, err.Request.URL, err.Status)
}
