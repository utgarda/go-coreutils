// Copyright 2010 Nigel Gourlay. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 yes - output a string repeatedly until killed
*/
package main

import (
	"flag"
	"os"
	"fmt"
)

var doHelp = flag.Bool("help", false, "display help and exit")
var doVersion = flag.Bool("version", false, "output version information and exit")

func main() {
	var s string = "y"
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Try '%s --help' for more information\n", os.Args[0])
	}

	flag.Parse()
	if *doHelp {
		fmt.Printf("Usage: %s [STRING]...\n"+
			"  or:  %s OPTION\n"+
			"Repeatedly output a line with all specified STRING(s), or 'y'.\n\n"+
			"      --help     display this help and exit\n"+
			"      --version  output version information and exit\n\n"+
			"Report yes bugs to ngourlay@gmail.com\n",
			os.Args[0], os.Args[0])
		os.Exit(0)
	}
	if *doVersion {
		fmt.Printf("yes (go coreutils) 0.1" +
			"Copyright (C) 2010 Nigel Gourlay. All rights reserved." +
			"Use of this source code is governed by a BSD-style" +
			"license that can be found at <http://golang.org/LICENSE>")
		os.Exit(0)
	}
	if flag.NArg() != 0 {
		s = flag.Arg(0)
		for i := 1; i < flag.NArg()+1; i++ {
			s += " " + flag.Arg(i)
		}
	}
	s += "\n"
	for {
		fmt.Printf(s)
	}
}
