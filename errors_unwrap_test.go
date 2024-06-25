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
			err = stderrors.New("banana")
		})
		It("returns correct error", func() {
			Expect(result).To(Equal(stderrors.New("banana")))
		})
	})
	Context("fmt wrapped", func() {
		BeforeEach(func() {
			err = fmt.Errorf("bla: %w", stderrors.New("banana"))
		})
		It("returns correct error", func() {
			Expect(result).To(Equal(stderrors.New("banana")))
		})
	})
	Context("join wrapped", func() {
		BeforeEach(func() {
			err = stderrors.Join(stderrors.New("banana"))
		})
		It("returns correct error", func() {
			Expect(result).To(Equal(stderrors.New("banana")))
		})
	})
})
