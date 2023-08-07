// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

import (
	stderrors "errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	liberrors "github.com/bborbe/errors"
)

var _ = Describe("Unwrap", func() {
	var err error
	var result error
	var banana error
	BeforeEach(func() {
		banana = stderrors.New("banana")
	})
	JustBeforeEach(func() {
		result = liberrors.Unwrap(err)
	})
	Context("nil", func() {
		BeforeEach(func() {
			err = nil
		})
		It("returns no error", func() {
			Expect(result).To(BeNil())
		})
	})
	Context("not wrapped", func() {
		BeforeEach(func() {
			err = banana
		})
		It("returns correct error", func() {
			Expect(result).To(Equal(banana))
		})
	})
	Context("fmt wrapped", func() {
		BeforeEach(func() {
			err = fmt.Errorf("bla: %w", banana)
		})
		It("returns correct error", func() {
			Expect(result).To(Equal(banana))
		})
	})
	Context("join wrapped", func() {
		BeforeEach(func() {
			err = stderrors.Join(banana)
		})
		It("returns correct error", func() {
			Expect(result).To(Equal(banana))
		})
	})
})
