// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ctrlflow_test

import (
	"go/ast"
	"testing"

	"golang.org/x/tools/gop/analysis/analysistest"
	"golang.org/x/tools/gop/analysis/passes/ctrlflow"
	"golang.org/x/tools/internal/typeparams"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()

	// load testdata/src/a/a.go
	tests := []string{"a"}
	if typeparams.Enabled {
		// and testdata/src/typeparams/typeparams.go when possible
		tests = append(tests, "typeparams")
	}
	results := analysistest.Run(t, testdata, ctrlflow.Analyzer, tests...)

	// Perform a minimal smoke test on
	// the result (CFG) computed by ctrlflow.
	for _, result := range results {
		cfgs := result.Result.(*ctrlflow.CFGs)

		for _, decl := range result.Pass.Files[0].Decls {
			if decl, ok := decl.(*ast.FuncDecl); ok && decl.Body != nil {
				if cfgs.FuncDecl(decl) == nil {
					t.Errorf("%s: no CFG for func %s",
						result.Pass.Fset.Position(decl.Pos()), decl.Name.Name)
				}
			}
		}
	}
}
