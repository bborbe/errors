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

var _ = Describe("Errors", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	Context("Wrap", func() {
		It("returns an error", func() {
			err := errors.Wrap(ctx, errors.New(ctx, "banana"), "failed")
			Expect(err).NotTo(BeNil())
		})
		It("returns nil for nil error", func() {
			err := errors.Wrap(ctx, nil, "failed")
			Expect(err).To(BeNil())
		})
	})
	Context("Wrapf", func() {
		It("returns an error", func() {
			err := errors.Wrapf(ctx, errors.New(ctx, "banana"), "failed")
			Expect(err).NotTo(BeNil())
		})
		It("returns nil for nil error", func() {
			err := errors.Wrapf(ctx, nil, "failed")
			Expect(err).To(BeNil())
		})
	})
	Context("errors.Is", func() {
		var err error
		var target error
		var is bool
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
})
