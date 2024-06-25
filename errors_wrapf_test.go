// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/errors"
)

var _ = Describe("Wrapf", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	It("returns an error", func() {
		err := errors.Wrapf(ctx, errors.New(ctx, "banana"), "failed")
		Expect(err).NotTo(BeNil())
	})
	It("returns nil for nil error", func() {
		err := errors.Wrapf(ctx, nil, "failed")
		Expect(err).To(BeNil())
	})
})
