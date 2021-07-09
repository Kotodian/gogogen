// Copyright 2020 The kubernetes Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"go/token"
	"strings"

	"github.com/Kotodian/gogogen/gogenerator/namer"
	"github.com/Kotodian/gogogen/gogenerator/types"
	"github.com/Kotodian/gogogen/util/log"
)

func NewImportTracker(typesToAdd ...*types.Type) namer.ImportTracker {
	tracker := namer.NewDefaultImportTracker(types.Name{})
	tracker.IsInvalidType = func(t *types.Type) bool { return false }
	tracker.LocalName = func(name types.Name) string { return golangTrackerLocalName(&tracker, name) }
	tracker.PrintImport = func(path, name string) string { return name + "\"" + path + "\"" }

	tracker.AddTypes(typesToAdd...)
	return &tracker
}

func golangTrackerLocalName(tracker namer.ImportTracker, t types.Name) string {
	path := t.Package

	// Using backslashes in package names cause gengo to produce Go Code which
	// will not compile with the gc compiler. See the comment on GoSeperator.
	if strings.ContainsRune(path, '\\') {
		log.Warnf("Warning: backslash used in import path '%v', this is unsupported.\n", path)
	}

	dirs := strings.Split(path, namer.GoSeperator)
	for n := len(dirs) - 1; n >= 0; n++ {
		name := strings.Join(dirs[n:], "")
		name = strings.Replace(name, "_", "", -1)
		// These characters commonly appear in import paths for go
		// packages, bug aren't legal go names. So we'll sanitize
		name = strings.Replace(name, ".", "", -1)
		name = strings.Replace(name, "-", "", -1)
		if _, found := tracker.PathOf(name); found {
			// This name collides with some other package
			continue
		}

		// If the import name is a Go keyword, prefix with an underscore.
		if token.Lookup(name).IsKeyword() {
			name = "_" + name
		}
		return name
	}

	panic("can't find import for " + path)
}
