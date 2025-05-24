// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imgbytefy

import (
	"errors"
	"io"
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytefy(t *testing.T) {
	testCases := []struct {
		name          string
		fs            *mockFS
		tmplProcessor templateProcessor
		inputFile     string
		outputFile    string
		packageName   string
		identifier    string
		wantErr       string
	}{
		{
			name: "success",
			fs: &mockFS{
				openFiles: map[string]*mockFile{
					"img.png": {content: []byte{0xCA, 0xFE, 0xBA, 0xBE}},
				},
			},
			tmplProcessor: mockProcessor{exec: mockExecutor{}},
			inputFile:     "img.png",
			outputFile:    "mypkg/output.go",
			packageName:   "mypkg",
			identifier:    "ImageBytes",
		},
		{
			name:          "invalid identifier",
			fs:            &mockFS{},
			tmplProcessor: mockProcessor{},
			inputFile:     "img.png",
			outputFile:    "out.go",
			packageName:   "mypkg",
			identifier:    "123abc",
			wantErr:       "not a valid Go identifier",
		},
		{
			name: "open error",
			fs: &mockFS{
				openErr: errors.New("cannot open"),
			},
			tmplProcessor: mockProcessor{},
			inputFile:     "img.png",
			outputFile:    "out.go",
			packageName:   "mypkg",
			identifier:    "Valid",
			wantErr:       "failed to open input file",
		},
		{
			name: "read error",
			fs: &mockFS{
				openFiles: map[string]*mockFile{
					"img.png": {content: []byte{0x01}},
				},
				readErr: errors.New("read fail"),
			},
			tmplProcessor: mockProcessor{},
			inputFile:     "img.png",
			outputFile:    "out.go",
			packageName:   "mypkg",
			identifier:    "Valid",
			wantErr:       "error reading input",
		},
		{
			name: "template parse error",
			fs: &mockFS{
				openFiles: map[string]*mockFile{
					"img.png": {content: []byte{0x01}},
				},
			},
			tmplProcessor: mockProcessor{expectedErr: errors.New("some error")},
			inputFile:     "img.png",
			outputFile:    "out.go",
			packageName:   "mypkg",
			identifier:    "Valid",
			wantErr:       "template parsing error",
		},
		{
			name: "template execution error",
			fs: &mockFS{
				openFiles: map[string]*mockFile{
					"img.png": {content: []byte{0x01}},
				},
			},
			tmplProcessor: mockProcessor{exec: mockExecutor{expectedErr: errors.New("some error")}},
			inputFile:     "img.png",
			outputFile:    "out.go",
			packageName:   "mypkg",
			identifier:    "Valid",
			wantErr:       "template execution error",
		},
		{
			name: "write file error",
			fs: &mockFS{
				openFiles: map[string]*mockFile{
					"img.png": {content: []byte{0x01}},
				},
				writeErr: errors.New("write fail"),
			},
			tmplProcessor: mockProcessor{exec: mockExecutor{}},
			inputFile:     "img.png",
			outputFile:    "out.go",
			packageName:   "mypkg",
			identifier:    "Valid",
			wantErr:       "failed to write output file",
		},
		{
			name: "outputFile empty, auto-generate path",
			fs: &mockFS{
				openFiles: map[string]*mockFile{
					"foo/image.png": {content: []byte{0xaa, 0xbb}},
				},
			},
			tmplProcessor: mockProcessor{exec: mockExecutor{}},
			inputFile:     "foo/image.png",
			outputFile:    "", // <== triggers auto-generation
			packageName:   "generated",
			identifier:    "ImageData",
			wantErr:       "",
		},
		{
			name: "mkdirAll fails",
			fs: &mockFS{
				openFiles: map[string]*mockFile{
					"img.png": {content: []byte{0x01}},
				},
				mkdirErr: errors.New("mkdir failed"),
			},
			tmplProcessor: mockProcessor{exec: mockExecutor{}},
			inputFile:     "img.png",
			outputFile:    "pkg/image.go",
			packageName:   "pkg",
			identifier:    "Data",
			wantErr:       "failed to create output directory",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fsProvider = tc.fs
			templateProcessorProvider = tc.tmplProcessor

			err := Bytefy(tc.inputFile, tc.outputFile, tc.packageName, tc.identifier)
			if tc.wantErr == "" {
				assert.NoError(t, err, "expected no error")
			} else {
				assert.ErrorContains(t, err, tc.wantErr)
			}
		})
	}
}

type mockFile struct {
	content []byte
	offset  int
	readErr error
}

func (m *mockFile) Read(p []byte) (int, error) {
	if m.readErr != nil {
		return 0, m.readErr
	}
	if m.offset >= len(m.content) {
		return 0, io.EOF
	}
	n := copy(p, m.content[m.offset:])
	m.offset += n
	return n, nil
}

func (m *mockFile) Close() error { return nil }

type mockFS struct {
	openFiles map[string]*mockFile
	openErr   error
	readErr   error
	mkdirErr  error
	writeErr  error
}

func (m *mockFS) Open(name string) (io.ReadCloser, error) {
	if m.openErr != nil {
		return nil, m.openErr
	}
	if f, ok := m.openFiles[name]; ok {
		f.readErr = m.readErr
		return f, nil
	}
	return nil, errors.New("file not found")
}

func (m *mockFS) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return m.writeErr
}

func (m *mockFS) MkdirAll(path string, perm fs.FileMode) error {
	return m.mkdirErr
}

type mockExecutor struct {
	expectedErr error
}

func (m mockExecutor) Execute(w io.Writer, data interface{}) error {
	return m.expectedErr
}

type mockProcessor struct {
	expectedErr error
	exec        templateExecutor
}

func (m mockProcessor) Parse(name, text string) (templateExecutor, error) {
	return m.exec, m.expectedErr
}
