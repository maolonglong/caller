// Copyright 2022 MaoLongLong. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package caller

import "runtime"

var _lru = newLRUCache(512)

func Get(skip int) (pc uintptr, file string, line int, ok bool) {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(skip+1, rpc[:])
	if n < 1 {
		return
	}
	var frame runtime.Frame
	if cachedFrame, ok := _lru.load(rpc[0]); ok {
		frame = cachedFrame
	} else {
		frame, _ = runtime.CallersFrames(rpc).Next()
		_lru.store(rpc[0], frame)
	}
	return frame.PC, frame.File, frame.Line, frame.PC != 0
}
