// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imgbytefy

import "testing"

func Test_isValidGoIdent(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "valid identifier",
			input: "validIdentifier",
			want:  true,
		},
		{
			name: "invalid identifier - blank input",
		},
		{
			name:  "invalid identifier - starts with digit",
			input: "1invalid",
		},
		{
			name:  "invalid identifier - contains special character",
			input: "invalid@name",
		},
		{
			name:  "invalid identifier - keyword",
			input: "func",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidGoIdent(tc.input)
			if got != tc.want {
				t.Errorf("IsValidGoIdent(%q) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}
