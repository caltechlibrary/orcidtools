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
