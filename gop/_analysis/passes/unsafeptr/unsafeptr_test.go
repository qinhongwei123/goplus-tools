// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unsafeptr_test

import (
	"testing"

	"golang.org/x/tools/gop/analysis/analysistest"
	"golang.org/x/tools/gop/analysis/passes/unsafeptr"
	"golang.org/x/tools/internal/typeparams"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	pkgs := []string{"a"}
	if typeparams.Enabled {
		pkgs = append(pkgs, "typeparams")
	}
	analysistest.Run(t, testdata, unsafeptr.Analyzer, pkgs...)
}