// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imgbytefy

import (
	"io"
	"io/fs"
	"os"
)

// fileSystem interface abstracts the file system operations. This allows
// for easier testing by mocking file system interactions.
type fileSystem interface {
	Open(name string) (io.ReadCloser, error)
	WriteFile(name string, data []byte, perm fs.FileMode) error
	MkdirAll(path string, perm fs.FileMode) error
}

// osFileSystem struct implements the fileSystem interface using
// the standard library's os package. This is the real implementation
// that interacts with the actual file system.
type osFileSystem struct{}

func (osFileSystem) Open(name string) (io.ReadCloser, error) {
	return os.Open(name)
}

func (osFileSystem) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (osFileSystem) MkdirAll(path string, perm fs.FileMode) error {
	return os.MkdirAll(path, perm)
}
