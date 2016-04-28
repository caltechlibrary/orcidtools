//
// orcidmessage.go translates an ORCID API XML response on the file system to JSON
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
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	// Caltech library packages
	"github.com/caltechlibrary/ot"
)

var (
	showHelp    bool
	showVersion bool
)

func processFile(fname string) {
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("Can't read %s, %s\n", fname, err)
	}
	data := new(ot.OrcidMessage)
	err = xml.Unmarshal(src, &data)
	if err != nil {
		fmt.Printf("Can't parse %s, %s\n", fname, err)
	}
	jsonSource, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Can't convert to JSON %s, %s", fname, err)
	}

	fmt.Printf("%s", jsonSource)
}

func main() {
	appname := os.Args[0]
	flag.Parse()
	if showHelp == true {
		fmt.Printf(`USAGE: %s [OPTIONS] XML_RESULTS_FILENAME

 OPTIONS
`, appname)
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("\t-%s\t(defaults to %s) %s\n", f.Name, f.Value, f.Usage)
		})
		fmt.Printf(`
 Version %s
`, ot.Version)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Println(" Version %s\n", Version)
		os.Exit(0)
	}

	args := flag.Args()
	for _, fname := range args {
		processFile(fname)
	}
}
