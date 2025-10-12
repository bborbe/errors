// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

type ErrorWithUnwrap struct {
	Err error
}

func (e ErrorWithUnwrap) Unwrap() error {
	return e.Err
}

func (e ErrorWithUnwrap) Error() string {
	return "ErrorWithUnwrap"
}
