// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	liberrors "github.com/bborbe/errors"
)

var _ = Describe("Join", func() {
	var err error
	var input []error
	JustBeforeEach(func() {
		err = liberrors.Join(input...)
	})
	Context("no errors", func() {
		BeforeEach(func() {
			input = []error{}
		})
		It("returns error", func() {
			Expect(err).To(BeNil())
		})
	})
	Context("multiple errors", func() {
		BeforeEach(func() {
			input = []error{
				errors.New("a"),
				errors.New("b"),
				errors.New("c"),
			}
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
		})
		It("returns error", func() {
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("[\na\nb\nc\n]"))
		})
	})
})
