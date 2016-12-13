/**
 * orcid is a command line utility for interacting with the ORCID API.
 * Currently it supports reading activity.
 *
 * @author R. S. Doiel, <rsdoiel@caltech.edu>
 *
 * Copyright (c) 2016, Caltech
 * All rights not granted herein are expressly reserved by Caltech.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 * * Redistributions of source code must retain the above copyright notice, this
 *   list of conditions and the following disclaimer.
 *
 * * Redistributions in binary form must reproduce the above copyright notice,
 *   this list of conditions and the following disclaimer in the documentation
 *   and/or other materials provided with the distribution.
 *
 * * Neither the name of ot nor the names of its
 *   contributors may be used to endorse or promote products derived from
 *   this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/ot"
)

var (
	usage = "USAGE: %s [OPTIONS] ORCID"

	description = `
SYSNOPIS

%s is a command line tool for harvesting ORCID data from the ORCID API.
See http://orcid.org/organizations/integrators for details. It requires
a client id and secret to access. This is set via environment variables
or the command line.

CONFIGURATION

+ ORCID_API_URL - set the URL for accessing the ORCID API (e.g. sandbox or members URL)
+ ORCID_CLIENT_ID - the client id for your registered ORCID app
+ ORCID_SECRET - the client secret needed to aquire an access token for the AP
`

	examples = `
EXAMPLES

Get an ORCID "works" from the sandbox for a given ORCID id.

    export ORCID_API_URL="https://pub.sandbox.orcid.org"
	export ORCID_CLIENT_ID="APP-01XX65MXBF79VJGF"
	export ORCID_CLIENT_SECRET="3a87028d-c84c-4d5f-8ad5-38a93181c9e1"
	%s -works 0000-0003-0900-6903

`

	// Standard Options
	showHelp    bool
	showLicense bool
	showVersion bool

	// Application Options
	showProfile      bool
	showWorks        bool
	showBio          bool
	showActivities   bool
	showAffiliations bool
	showFunding      bool

	// Required
	apiURL       string
	clientID     string
	clientSecret string
	orcidID      string
)

func init() {
	// Standard Options
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")

	// Application Options
	flag.BoolVar(&showProfile, "profile", false, "display profile")
	flag.BoolVar(&showWorks, "works", false, "display works")
	flag.BoolVar(&showBio, "bio", false, "display biography")
	flag.BoolVar(&showActivities, "activities", false, "display activity")
	flag.BoolVar(&showAffiliations, "affiliations", false, "display affiliations")
	flag.BoolVar(&showFunding, "funding", false, "display funding")

	flag.StringVar(&orcidID, "o", "", "use orcid id")
	flag.StringVar(&orcidID, "orcid", "", "use orcid id")
}

func main() {
	appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	cfg := cli.New(appName, "ORCID", fmt.Sprintf(ot.LicenseText, appName, ot.Version), ot.Version)
	cfg.UsageText = fmt.Sprintf(usage, appName)
	cfg.DescriptionText = fmt.Sprintf(description, appName)
	cfg.OptionsText = "OPTIONS\n"
	cfg.ExampleText = fmt.Sprintf(examples, appName)

	// Process flags and update the environment as needed.
	if showHelp == true {
		fmt.Println(cfg.Usage())
		os.Exit(0)
	}
	if showLicense == true {
		fmt.Println(cfg.License())
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Println(cfg.Version())
		os.Exit(0)
	}

	if len(args) > 0 {
		orcidID = args[0]
	}

	apiURL = cfg.CheckOption("api_url", cfg.MergeEnv("api_url", apiURL), true)
	clientID = cfg.CheckOption("client_id", cfg.MergeEnv("client_id", clientID), true)
	clientSecret = cfg.CheckOption("client_secret", cfg.MergeEnv("client_secret", clientSecret), true)
	orcidID = cfg.CheckOption("orcid_id", cfg.MergeEnv("orcid_id", orcidID), true)

	var requestType string

	if showProfile == true {
		requestType = "orcid-profile"
	}
	if showWorks == true {
		requestType = "orcid-works"
	}
	if showBio == true {
		requestType = "orcid-bio"
	}
	if showActivities == true {
		requestType = "acitivities"
	}
	if showAffiliations == true {
		requestType = "affiliations"
	}
	if showFunding == true {
		requestType = "funding"
	}
	if requestType == "" {
		fmt.Fprintf(os.Stderr, "Not sure what to do, see %s -help", appName)
		os.Exit(1)
	}

	// Setup the API access
	api, err := ot.New(apiURL, clientID, clientSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	src, err := api.Request("GET", fmt.Sprintf("/v1.2/%s/%s", orcidID, requestType), map[string]string{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", src)
}
