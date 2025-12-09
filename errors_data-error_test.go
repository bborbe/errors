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

var _ = Describe("DataError", func() {
	It("returns nothing for standard error", func() {
		err := errors.New(
			context.Background(),
			"banana",
		)
		data := errors.DataFromError(err)
		Expect(data).To(HaveLen(0))
		Expect(err.Error()).To(Equal("banana"))
	})
	It("returns data for DataError", func() {
		err := errors.AddDataToError(
			errors.New(context.Background(),
				"banana",
			),
			map[string]any{
				"hello": "world",
			},
		)
		data := errors.DataFromError(err)
		Expect(data).To(HaveLen(1))
		Expect(data).To(HaveKeyWithValue("hello", "world"))
		Expect(err.Error()).To(Equal("banana"))
	})
	It("returns data if DataError is wrapped", func() {
		err := errors.Wrap(
			context.Background(),
			errors.AddDataToError(
				errors.New(context.Background(),
					"banana",
				),
				map[string]any{
					"hello": "world",
				},
			),
			"foo bar",
		)
		data := errors.DataFromError(err)
		Expect(data).To(HaveLen(1))
		Expect(data).To(HaveKeyWithValue("hello", "world"))
		Expect(err.Error()).To(Equal("foo bar: banana"))
	})
	It("combines data", func() {
		err := errors.AddDataToError(
			errors.New(
				context.Background(),
				"banana",
			),
			map[string]any{
				"hello": "world",
			},
		)
		err = errors.AddDataToError(err, map[string]any{"hallo": "welt"})
		data := errors.DataFromError(err)
		Expect(data).To(HaveLen(2))
		Expect(data).To(HaveKeyWithValue("hello", "world"))
		Expect(data).To(HaveKeyWithValue("hallo", "welt"))
		Expect(err.Error()).To(Equal("banana"))
	})
	It("supports array values", func() {
		err := errors.AddDataToError(
			errors.New(context.Background(), "test"),
			map[string]any{
				"ids": []string{"a", "b", "c"},
			},
		)
		data := errors.DataFromError(err)
		ids, ok := data["ids"].([]string)
		Expect(ok).To(BeTrue())
		Expect(ids).To(Equal([]string{"a", "b", "c"}))
	})
	It("supports mixed type values", func() {
		err := errors.AddDataToError(
			errors.New(context.Background(), "test"),
			map[string]any{
				"field": "email",
				"count": 5,
				"valid": false,
			},
		)
		data := errors.DataFromError(err)
		Expect(data["field"]).To(Equal("email"))
		Expect(data["count"]).To(Equal(5))
		Expect(data["valid"]).To(Equal(false))
	})
	It("supports nested objects", func() {
		err := errors.AddDataToError(
			errors.New(context.Background(), "test"),
			map[string]any{
				"validation": map[string]any{
					"field":  "email",
					"reason": "invalid_format",
				},
			},
		)
		data := errors.DataFromError(err)
		validation, ok := data["validation"].(map[string]any)
		Expect(ok).To(BeTrue())
		Expect(validation["field"]).To(Equal("email"))
		Expect(validation["reason"]).To(Equal("invalid_format"))
	})
})
