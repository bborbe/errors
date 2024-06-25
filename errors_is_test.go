// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

import (
	"context"
	stderrors "errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/errors"
)

var _ = Describe("Is", func() {
	var err error
	var target error
	var is bool
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	JustBeforeEach(func() {
		is = stderrors.Is(err, target)
	})
	Context("same", func() {
		BeforeEach(func() {
			err = context.Canceled
			target = context.Canceled
		})
		It("returns true", func() {
			Expect(is).To(BeTrue())
		})
	})
	Context("wrapped", func() {
		BeforeEach(func() {
			err = errors.Wrapf(ctx, context.Canceled, "banana")
			target = context.Canceled
		})
		It("returns true", func() {
			Expect(is).To(BeTrue())
		})
	})
})
