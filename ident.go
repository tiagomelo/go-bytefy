// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imgbytefy

import (
	"go/token"
	"unicode"
)

// isValidGoIdent checks if the given name is a
// valid Go identifier and not a keyword.
func isValidGoIdent(name string) bool {
	if name == "" {
		return false
	}
	runes := []rune(name)
	if !unicode.IsLetter(runes[0]) && runes[0] != '_' {
		return false
	}
	for _, r := range runes[1:] {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return false
		}
	}
	return !token.Lookup(name).IsKeyword()
}
