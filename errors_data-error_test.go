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
			map[string]string{
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
				map[string]string{
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
			map[string]string{
				"hello": "world",
			},
		)
		err = errors.AddDataToError(err, map[string]string{"hallo": "welt"})
		data := errors.DataFromError(err)
		Expect(data).To(HaveLen(2))
		Expect(data).To(HaveKeyWithValue("hello", "world"))
		Expect(data).To(HaveKeyWithValue("hallo", "welt"))
		Expect(err.Error()).To(Equal("banana"))
	})
})
