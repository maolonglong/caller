// Copyright 2022 MaoLongLong. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package caller

import (
	"fmt"
	"io"
	"strings"
)

var _ io.Writer = (*StdLogWriter)(nil)

type StdLogWriter struct {
	Out io.Writer
}

func (w *StdLogWriter) Write(p []byte) (n int, err error) {
	_, file, line, _ := Get(4) // may only support some versions
	a := strings.Split(file, "/")
	l := len(a)
	prefix := fmt.Sprintf("[%s/%s:%d] ", a[l-2], a[l-1], line)
	n1, err := w.Out.Write([]byte(prefix))
	if err != nil {
		return
	}
	n2, err := w.Out.Write(p)
	return n1 + n2, err
}
