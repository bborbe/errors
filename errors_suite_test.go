// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

type ErrorWithMessage struct {
	Message string
}

func (e ErrorWithMessage) Error() string {
	return e.Message
}

type ErrorWithCause struct {
	Err error
}

func (e ErrorWithCause) Cause() error {
	return e.Err
}

func (e ErrorWithCause) Error() string {
	return "ErrorWithCause"
}

type ErrorWithUnwrap struct {
	Err error
}

func (e ErrorWithUnwrap) Unwrap() error {
	return e.Err
}

func (e ErrorWithUnwrap) Error() string {
	return "ErrorWithUnwrap"
}

type ErrorWithUnwrapList struct {
	Errors []error
}

func (e ErrorWithUnwrapList) Unwrap() []error {
	return e.Errors
}

func (e ErrorWithUnwrapList) Error() string {
	return "ErrorWithUnwrapList"
}

//go:generate go run -mod=vendor github.com/maxbrunsfeld/counterfeiter/v6 -generate
func TestSuite(t *testing.T) {
	time.Local = time.UTC
	format.TruncatedDiff = false
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}
