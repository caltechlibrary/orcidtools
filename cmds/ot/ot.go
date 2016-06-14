//
// ot.go translates an ORCID API XML response on the file system to JSON
// on the filesystem. It was create to debug some of the response parsing for XML.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	// 3rd Party Libraries
	"github.com/robertkrimen/otto"

	// Caltech library packages
	"github.com/caltechlibrary/ostdlib"
	"github.com/caltechlibrary/ot"
)

var (
	showHelp      bool
	showVersion   bool
	jsInteractive bool
)

type Expr struct {
	ORCID string `json:"orcid,omitempty"`
	Path  string `json:"path,omitempty"`
}

func processExpression(api *ot.OrcidAPI, src []byte) {
	var (
		data *ot.OrcidMessage
		err  error
	)
	expr := new(Expr)

	err = json.Unmarshal(src, &expr)
	if err != nil {
		log.Fatalf("Do not know how to process %s", src)
	}
	switch {
	case strings.HasPrefix(expr.Path, "/orcid-bio"):
		data, err = api.GetBio(expr.ORCID)
	case strings.HasPrefix(expr.Path, "/orcid-works"):
		data, err = api.GetWorks(expr.ORCID)
	case strings.HasPrefix(expr.Path, "/orcid-profile"):
		data, err = api.GetProfile(expr.ORCID)
	default:
		data, err = api.Get(expr.Path, expr.ORCID)
	}
	if err != nil {
		log.Fatalf("%s", err)
	}
	src, _ = json.Marshal(data)
	fmt.Printf("%s", src)
}

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help information")
	flag.BoolVar(&showVersion, "v", false, "display version information")
	flag.BoolVar(&jsInteractive, "i", false, "run an interactive JavaScript REPL")
}

func main() {
	appname := os.Args[0]
	flag.Parse()
	if showHelp == true {
		fmt.Printf(`USAGE: %s [OPTIONS] JSON_EXPRESSION|JS_FILENAME

orcid message connects to the orcid API and submits a request based on the
JSON EXPRESSION provided or runs the JavaScript file provides.

The JSON expression has two fields ORCID and path
where path can be one of three end points.

	/orcid-bio/
	/orcid-works/
	/orcid-profile/

The expression or JavaScript file is run in the order listed in the command line.

EXAMPLES

	%s '{"orcid": "0000-0002-2389-8429", "path": "/orcid-bio/"}'
	%s '{"orcid": "0000-0002-2389-8429", "path": "/orcid-works/"}'
	%s '{"orcid": "0000-0002-2389-8429", "path": "/orcid-profile/"}'

	%s my-script.js

 OPTIONS
`, appname, appname, appname, appname)
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("\t-%s\t(defaults to %s) %s\n", f.Name, f.Value, f.Usage)
		})
		fmt.Printf(`
 Version %s
`, ot.Version)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Println(" Version %s\n", ot.Version)
		os.Exit(0)
	}

	api := ot.New()
	if api == nil {
		log.Fatalf("Environment not setup. Try %s -h for usage", appname)
	}
	_, err := api.Login()
	if err != nil {
		log.Fatalf("%s", err)
	}

	vm := otto.New()
	js := ostdlib.New(vm)
	js.AddExtensions()
	api.AddExtensions(js)

	args := flag.Args()
	for _, expr := range args {
		if strings.HasSuffix(expr, ".js") == true {
			js.Run(expr)
		} else {
			processExpression(api, []byte(expr))
		}
	}
	if jsInteractive == true {
		js.AddHelp()
		api.AddHelp(js)
		js.AddAutoComplete()
		js.PrintDefaultWelcome()
		js.Repl()
	}
}
