// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

type ErrorWithUnwrapList struct {
	Errors []error
}

func (e ErrorWithUnwrapList) Unwrap() []error {
	return e.Errors
}

func (e ErrorWithUnwrapList) Error() string {
	return "ErrorWithUnwrapList"
}
