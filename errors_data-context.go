// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package errors

import (
	"context"
	"sync"
)

type dataCtxKeyType string

const dataCtxKey dataCtxKeyType = "data"

var mutex sync.Mutex

func AddContextDataToError(ctx context.Context, err error) error {
	return AddDataToError(err, DataFromContext(ctx))
}

func AddToContext(ctx context.Context, key string, value any) context.Context {
	v := ctx.Value(dataCtxKey)
	if v == nil {
		return context.WithValue(ctx, dataCtxKey, map[string]any{
			key: value,
		})
	}
	data, ok := v.(map[string]any)
	if ok {
		mutex.Lock()
		data[key] = value
		mutex.Unlock()
	}
	return ctx
}

func DataFromContext(ctx context.Context) map[string]any {
	value := ctx.Value(dataCtxKey)
	if value == nil {
		return nil
	}
	return value.(map[string]any)
}
