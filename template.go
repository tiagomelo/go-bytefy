// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package imgbytefy

import (
	"io"
	"text/template"
)

// templateExecutor interface abstracts the execution
// of a parsed template.
type templateExecutor interface {
	Execute(wr io.Writer, data interface{}) error
}

// textTemplateExecutor is a concrete implementation of
// the templateExecutor interface.
type textTemplateExecutor struct {
	tmpl *template.Template
}

// templateProcessor interface abstracts the parsing of a template.
type templateProcessor interface {
	Parse(name, text string) (templateExecutor, error)
}

// textTemplateProcessor is a concrete implementation of
// the templateProcessor interface.
type textTemplateProcessor struct{}

// Parse parses the provided template text and returns
// a templateExecutor.
func (textTemplateProcessor) Parse(name, text string) (templateExecutor, error) {
	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		return nil, err
	}
	return textTemplateExecutor{tmpl}, nil
}

// Execute executes the parsed template with the provided
// data and writes the output to the writer.
func (r textTemplateExecutor) Execute(wr io.Writer, data interface{}) error {
	return r.tmpl.Execute(wr, data)
}
