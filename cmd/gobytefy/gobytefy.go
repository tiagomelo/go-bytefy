// Copyright (c) 2025 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	imgbytefy "github.com/tiagomelo/go-bytefy"
)

type options struct {
	PackageName string `short:"p" long:"package" description:"Go package name" required:"true"`
	InputFile   string `short:"f" long:"file" description:"Input image file" required:"true"`
	OutputFile  string `short:"o" long:"output" description:"Output Go file" required:"true"`
	Identifier  string `short:"i" long:"id" description:"Identifier for the byte array (must be a valid Go identifier)" required:"true"`
}

func run(opts *options) error {
	if err := imgbytefy.Bytefy(opts.InputFile, opts.OutputFile, opts.PackageName, opts.Identifier); err != nil {
		return err
	}
	return nil
}

func main() {
	var opts options
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				fmt.Println(err)
				os.Exit(0)
			}
			fmt.Println(err)
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
	if err := run(&opts); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
