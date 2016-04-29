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
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

const (
	Version = "0.0.0"
)

var (
	showHelp    bool
	showVersion bool
)

func processTarBall(fname string) error {
	reader, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer reader.Close()

	fmt.Printf("DEBUG getting ready to read %s\n", fname)
	// Now setup to read the tar format
	tarReader := tar.NewReader(reader)

	//NOTE: I am only going to read the first 10 entries while I am sorting things out.
	//for {
	for i := 0; i < 100; i++ {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fmt.Printf("%d next tar filename: %s\n", i, header.Name)
		if strings.HasPrefix(header.Name, "./json/") {
			info := header.FileInfo()
			fmt.Printf("\t%+v\n", info)
		}
	}

	return nil
}

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help information")
	flag.BoolVar(&showVersion, "v", false, "display version information")
}

func main() {
	flag.Parse()
	appname := path.Base(os.Args[0])
	if showHelp == true {
		fmt.Printf(` USAGE: %s [OPTIONS] PUBLIC_RELEASE_DATA_TAR_FILENAME

 %s transformations an ORCID Public Release tar file into a database suitable for
 generating triples, key/value pairs, relational data or for feeding into a search engine.
 
 OPTIONS

`, appname, appname)
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("    -%s  (defaults to %s) %s\n", f.Name, f.Value, f.Usage)
		})
		fmt.Printf("\n Version %s\n", Version)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Printf(" Version %s\n", Version)
		os.Exit(0)
	}
	args := flag.Args()
	for _, fname := range args {
		log.Printf("Processing %s\n", fname)
		err := processTarBall(fname)
		if err != nil {
			log.Printf("%s, %s\n", fname, err)
		}
	}
}
