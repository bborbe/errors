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

var _ = Describe("DataContext", func() {
	It("saves data into context", func() {
		ctx := errors.AddToContext(context.Background(), "myKey", "myValue")
		data := errors.DataFromContext(ctx)
		Expect(data).NotTo(BeNil())
		Expect(data).To(HaveKeyWithValue("myKey", "myValue"))
	})
	It("data is not overwriteable", func() {
		ctx := errors.AddToContext(context.Background(), "myKey", "myValue")
		ctx = context.WithValue(ctx, "data", "banana")
		data := errors.DataFromContext(ctx)
		Expect(data).NotTo(BeNil())
		Expect(data).To(HaveKeyWithValue("myKey", "myValue"))
	})
})
