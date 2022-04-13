// Copyright 2022 MaoLongLong. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package caller

import (
	"runtime"
	"testing"
)

func BenchmarkRuntimeCaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runtime.Caller(1)
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get(1)
	}
}

func BenchmarkGetParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Get(1)
		}
	})
}
