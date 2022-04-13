// Copyright 2022 MaoLongLong. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package caller

import "fmt"

func ExampleGet() {
	_, file, line, _ := Get(1)
	fmt.Printf("%s:%d", file, line)

	// Output: go.chensl.me/caller/example_test.go:10
}
