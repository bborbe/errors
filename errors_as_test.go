// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

import (
	stderrors "errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	liberrors "github.com/bborbe/errors"
)

var _ = Describe("As", func() {
	var target error
	var err error
	var result bool
	JustBeforeEach(func() {
		result = liberrors.As(err, target)
	})
	Context("not matching", func() {
		BeforeEach(func() {
			err = stderrors.New("banana")
			target = &ErrorWithMessage{}
		})
		It("returns be false", func() {
			Expect(result).To(BeFalse())
			Expect(target.Error()).To(Equal(""))
		})
	})
	Context("matching", func() {
		BeforeEach(func() {
			err = ErrorWithMessage{
				Message: "banana",
			}
			target = &ErrorWithMessage{}
		})
		It("returns be true", func() {
			Expect(result).To(BeTrue())
			Expect(target.Error()).To(Equal("banana"))
		})
	})
	Context("matching with wrap", func() {
		BeforeEach(func() {
			err = ErrorWithUnwrap{
				Err: ErrorWithMessage{
					Message: "banana",
				},
			}
			target = &ErrorWithMessage{}
		})
		It("returns be true", func() {
			Expect(result).To(BeTrue())
			Expect(target.Error()).To(Equal("banana"))
		})
	})
	Context("matching with cause", func() {
		BeforeEach(func() {
			err = ErrorWithCause{
				Err: ErrorWithMessage{
					Message: "banana",
				},
			}
			target = &ErrorWithMessage{}
		})
		It("returns be true", func() {
			Expect(result).To(BeTrue())
			Expect(target.Error()).To(Equal("banana"))
		})
	})
})
