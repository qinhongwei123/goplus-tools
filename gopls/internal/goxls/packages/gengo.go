// Copyright 2023 The GoPlus Authors (goplus.org). All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packages

import (
	"context"
	"log"
	"path/filepath"
	"strings"

	"github.com/goplus/gop/env"
	"golang.org/x/tools/gopls/internal/goxls/langserver"
)

var (
	gopInstalled = env.Installed()
)

func GenGo(patternIn ...string) (patternOut []string, err error) {
	if !gopInstalled {
		return patternIn, nil
	}
	pattern, patternOut := buildPattern(patternIn)
	log.Println("GenGo:", pattern, "in:", patternIn, "out:", patternOut)
	langserver.GenGo(context.Background(), pattern...)
	return
}

type none = struct{}

func buildPattern(pattern []string) (gopPattern []string, allPattern []string) {
	const filePrefix = "file="
	gopPattern = make([]string, 0, len(pattern))
	allPattern = make([]string, 0, len(pattern))
	dirs := make(map[string]none)
	for _, v := range pattern {
		if strings.HasPrefix(v, filePrefix) {
			file := v[len(filePrefix):]
			dir := filepath.Dir(file)
			if strings.HasSuffix(file, ".go") { // skip go file
				allPattern = append(allPattern, v)
			} else {
				dirs[dir] = none{}
			}
			continue
		}
		allPattern = append(allPattern, v)
		if pos := strings.Index(v, "/"); pos >= 0 {
			if pos > 0 {
				domain := v[:pos]
				if !strings.Contains(domain, ".") || domain == "golang.org" { // std or golang.org
					continue
				}
			}
			gopPattern = append(gopPattern, v)
		}
	}
	for dir := range dirs {
		gopPattern = append(gopPattern, dir)
		allPattern = append(allPattern, dir)
	}
	return
}
