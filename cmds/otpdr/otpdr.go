//
// otpdr.go translates an ORCID API XML response on the file system to JSON
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
	"time"

	// Caltech Library Packages
	"github.com/boltdb/bolt"
	"github.com/caltechlibrary/ot"
)

var (
	showHelp    bool
	showVersion bool
)

func processTarBall(fname, dbname string) error {
	db, err := bolt.Open(dbname, 0600, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: false})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reader, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer reader.Close()

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
			//FIXME: decode the actually content and save to storage format.
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
	flag.Parse()
	if showHelp == true {
		fmt.Printf(`USAGE: %s [OPTIONS] TAR_FILENAME DB_FILENAME

%s process an Orcid Public Data release file and turns it into an
key/value database file suitable for further processing.

EXAMPLE

	%s ORCID_public_data_file_2015.tar pdr2015.boltDB

 OPTIONS

`, appname, appname, appname)
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("\t-%s\t(defaults to %s) %s\n", f.Name, f.Value, f.Usage)
		})
		fmt.Printf(`
 Version %s
`, ot.Version)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Printf(" Version %s\n", ot.Version)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) != 2 {
		fmt.Printf("Expecting a tar filename and db filename. Try %s -h for usage.\n", appname)
		os.Exit(1)
	}

	err := processTarBall(args[0], args[1])
	if err != nil {
		log.Printf("%s, %s\n", args[0], err)
	}
}
