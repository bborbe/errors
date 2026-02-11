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
	It("derived contexts do not share data", func() {
		// Create parent context with initial data
		parent := errors.AddToContext(context.Background(), "parent", "value")

		// Derive two child contexts with different data
		child1 := errors.AddToContext(parent, "child", "one")
		child2 := errors.AddToContext(parent, "child", "two")

		// Each child should have its own "child" value
		data1 := errors.DataFromContext(child1)
		data2 := errors.DataFromContext(child2)

		Expect(data1["child"]).To(Equal("one"))
		Expect(data2["child"]).To(Equal("two"))

		// Both should have the parent data
		Expect(data1["parent"]).To(Equal("value"))
		Expect(data2["parent"]).To(Equal("value"))
	})
	It("is safe for concurrent use from multiple goroutines", func() {
		// This test verifies there are no data races when multiple goroutines
		// call AddToContext concurrently on contexts derived from the same parent.
		// Run with: go test -race ./...
		parent := errors.AddToContext(context.Background(), "shared", "parent")

		const numGoroutines = 100
		done := make(chan bool, numGoroutines)
		results := make(chan map[string]any, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				// Each goroutine derives its own context and adds unique data
				ctx := errors.AddToContext(parent, "goroutine", id)
				ctx = errors.AddToContext(ctx, "nested", id*10)

				// Verify the data is correct for this goroutine
				data := errors.DataFromContext(ctx)
				results <- data
			}(i)
		}

		// Wait for all goroutines to complete
		for i := 0; i < numGoroutines; i++ {
			<-done
		}
		close(results)

		// Verify each result has consistent data
		for data := range results {
			Expect(data["shared"]).To(Equal("parent"))
			goroutineID, ok := data["goroutine"].(int)
			Expect(ok).To(BeTrue())
			Expect(data["nested"]).To(Equal(goroutineID * 10))
		}
	})
})
