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

type contextKey string

const testKey contextKey = "data"

var _ = Describe("DataContext", func() {
	It("saves data into context", func() {
		ctx := errors.AddToContext(context.Background(), "myKey", "myValue")
		data := errors.DataFromContext(ctx)
		Expect(data).NotTo(BeNil())
		Expect(data).To(HaveKeyWithValue("myKey", "myValue"))
	})
	It("data is not overwriteable", func() {
		ctx := errors.AddToContext(context.Background(), "myKey", "myValue")
		ctx = context.WithValue(ctx, testKey, "banana")
		data := errors.DataFromContext(ctx)
		Expect(data).NotTo(BeNil())
		Expect(data).To(HaveKeyWithValue("myKey", "myValue"))
	})
	It("supports non-string values", func() {
		ctx := errors.AddToContext(context.Background(), "count", 42)
		ctx = errors.AddToContext(ctx, "items", []string{"a", "b"})
		ctx = errors.AddToContext(ctx, "enabled", true)
		data := errors.DataFromContext(ctx)
		Expect(data).NotTo(BeNil())
		Expect(data["count"]).To(Equal(42))
		Expect(data["items"]).To(Equal([]string{"a", "b"}))
		Expect(data["enabled"]).To(Equal(true))
	})
})
