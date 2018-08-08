/**
 * orcid is a command line utility for interacting with the ORCID API.
 * Currently it supports reading activity.
 *
 * @author R. S. Doiel, <rsdoiel@caltech.edu>
 *
 * Copyright (c) 2018, Caltech
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
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/cli"
	"github.com/caltechlibrary/dotpath"
	ot "github.com/caltechlibrary/ot"
)

var (
	synopsis = `
_orcid_ is a program for harvesting ORCID from the orcid.org API.
`
	description = `
_orcid_ is a command line tool for harvesting ORCID data from the ORCID API.
See http://orcid.org/organizations/integrators for details. It requires
a client id and secret to access. This is set via environment variables
or the command line.

CONFIGURATION

+ ORCID_API_URL - set the URL for accessing the ORCID API (e.g. sandbox or members URL)
+ ORCID_CLIENT_ID - the client id for your registered ORCID app
+ ORCID_SECRET - the client secret needed to aquire an access token for the AP

`

	examples = `
Get an ORCID "works" from the sandbox for a given ORCID id.
` + "```" + `
    export ORCID_API_URL="https://pub.sandbox.orcid.org"
    export ORCID_CLIENT_ID="APP-01XX65MXBF79VJGF"
    export ORCID_CLIENT_SECRET="3a87028d-c84c-4d5f-8ad5-38a93181c9e1"
    orcid -works 0000-0003-0900-6903
` + "```" + `
`

	// Standard Options
	showHelp         bool
	showLicense      bool
	showVersion      bool
	showExamples     bool
	verbose          bool
	outputFName      string
	generateMarkdown bool
	generateManPage  bool

	// Application Options
	showRecord              bool
	showPerson              bool
	showAddress             bool
	showEmail               bool
	showExternalIdentifiers bool
	showKeywords            bool
	showOtherNames          bool
	showPersonalDetails     bool
	showResearcherURLS      bool
	showActivities          bool
	showEducations          bool
	showEmployments         bool
	showFundings            bool
	showPeerReviews         bool
	showProfile             bool
	showWorks               bool
	showWorksDetailed       bool
	searchString            string

	// Required
	apiURL       string
	clientID     string
	clientSecret string
	orcidID      string
)

func init() {
}

func main() {
	appName := path.Base(os.Args[0])
	app := cli.NewCli(ot.Version)
	app.AddHelp("synopsis", []byte(synopsis))
	app.AddHelp("description", []byte(description))
	app.AddHelp("examples", []byte(examples))
	app.AddHelp("license", []byte(fmt.Sprintf(ot.LicenseText, appName, ot.Version)))

	// Standard Options
	app.BoolVar(&showHelp, "h,help", false, "display help")
	app.BoolVar(&showLicense, "l,license", false, "display license")
	app.BoolVar(&showVersion, "v,version", false, "display version")
	app.BoolVar(&showExamples, "examples", false, "display example(s)")
	app.BoolVar(&verbose, "verbose", false, "enable verbose logging")
	app.StringVar(&outputFName, "o,output", "", "set output filename")
	app.BoolVar(&generateMarkdown, "generate-markdown", false, "generate Markdown documentation")
	app.BoolVar(&generateManPage, "generate-manpage", false, "generate man page")

	// Application Options
	app.BoolVar(&showRecord, "record", false, "display record")
	app.BoolVar(&showPerson, "person", false, "display person")
	app.BoolVar(&showAddress, "address", false, "display address")
	app.BoolVar(&showEmail, "email", false, "display email")
	app.BoolVar(&showExternalIdentifiers, "external-ids", false, "display external identifies")
	app.BoolVar(&showKeywords, "keywords", false, "display keywords")
	app.BoolVar(&showOtherNames, "other-names", false, "display other names")
	app.BoolVar(&showPersonalDetails, "personal-details", false, "display personal detials")
	app.BoolVar(&showResearcherURLS, "researcher-urls", false, "display researcher urls")
	app.BoolVar(&showActivities, "activities", false, "display activities")
	app.BoolVar(&showEducations, "educations", false, "display education affiliations")
	app.BoolVar(&showEmployments, "employments", false, "display employment affiliations")
	app.BoolVar(&showFundings, "fundings", false, "display funding activities")
	app.BoolVar(&showPeerReviews, "peer-reviews", false, "display peer review activities")
	app.BoolVar(&showWorks, "works", false, "display works summary")
	app.BoolVar(&showWorksDetailed, "works-detailed", false, "display works in detail")
	app.StringVar(&searchString, "search", "", "search for terms")

	app.StringVar(&orcidID, "O,orcid", "", "use orcid id")

	// Process apps and update the environment as needed.
	app.Parse()
	args := app.Args()

	if generateMarkdown {
		app.GenerateMarkdown(os.Stdout)
		os.Exit(0)
	}
	if generateManPage {
		app.GenerateManPage(os.Stdout)
		os.Exit(0)
	}

	if showHelp {
		if len(args) > 0 {
			fmt.Fprintf(os.Stdout, app.Help(args...))
		} else {
			app.Usage(os.Stdout)
		}
		os.Exit(0)
	}

	if showExamples {
		if len(args) > 0 {
			fmt.Fprintf(app.Out, app.Help(args...))
		}
		os.Exit(0)
	}
	if showLicense {
		fmt.Println(app.License())
		os.Exit(0)
	}
	if showVersion {
		fmt.Println(app.Version())
		os.Exit(0)
	}

	if len(args) > 0 {
		orcidID = args[0]
	}

	if apiURL == "" {
		apiURL = os.Getenv("ORCID_API_URL")
	}
	if clientID == "" {
		clientID = os.Getenv("ORCID_CLIENT_ID")
	}
	if clientSecret == "" {
		clientSecret = os.Getenv("ORCID_CLIENT_SECRET")
	}

	var requestType string

	if showRecord == true {
		requestType = "record"
	}
	if showPerson == true {
		requestType = "person"
	}
	if showAddress == true {
		requestType = "address"
	}
	if showEmail == true {
		requestType = "email"
	}
	if showExternalIdentifiers == true {
		requestType = "external-identifiers"
	}
	if showKeywords == true {
		requestType = "keywords"
	}
	if showOtherNames == true {
		requestType = "other-names"
	}
	if showPersonalDetails == true {
		requestType = "personal-details"
	}
	if showResearcherURLS == true {
		requestType = "researcher-urls"
	}
	if showActivities == true {
		requestType = "activities"
	}
	if showEducations == true {
		requestType = "educations"
	}
	if showEmployments == true {
		requestType = "employments"
	}
	if showFundings == true {
		requestType = "fundings"
	}
	if showPeerReviews == true {
		requestType = "peer-reviews"
	}
	if showWorks == true {
		requestType = "works"
	}
	if showWorksDetailed == true {
		requestType = "works-detailed"
	}
	if requestType == "" && searchString == "" {
		fmt.Fprintf(os.Stderr, "Not sure what to do, see %s -help", appName)
		os.Exit(1)
	}

	out, err := cli.Create(outputFName, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}
	defer cli.CloseFile(outputFName, out)

	// Setup the API access
	api, err := ot.New(apiURL, clientID, clientSecret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	if searchString != "" {
		qry := map[string]string{
			"q": searchString,
		}
		src, err := api.Request("GET", "/v2.0/search/", qry)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", src)
		os.Exit(0)
	}

	if orcidID == "" {
		orcidID = os.Getenv("ORCID_ID")
	}
	if requestType == "works-detailed" {
		src, err := api.Request("GET", fmt.Sprintf("/v2.0/%s/%s", orcidID, "works"), map[string]string{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		rawData := map[string]interface{}{}
		if err := json.Unmarshal(src, &rawData); err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		keyPath := `.group[:]["work-summary"][:]["put-code"]`
		if summary, err := dotpath.EvalJSON(keyPath, src); err == nil {
			workIds := []string{}
			for _, o := range summary.([]interface{}) {
				nInterface := o.([]interface{})
				for _, nVal := range nInterface {
					//s := fmt.Sprintf("%+v", nVal)
					workIds = append(workIds, fmt.Sprintf("%+v", nVal)) //s)
				}
			}
			if len(workIds) > 0 {
				src, err := api.Request("GET", fmt.Sprintf("/v2.0/%s/%s/%s", orcidID, "works", strings.Join(workIds, ",")), map[string]string{})
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s", err)
					os.Exit(1)
				}
				fmt.Fprintf(out, "%s\n", src)
			} else {
				//NOTE: Something went wrong with the detailed work ids, so just return our "works" request value
				fmt.Fprintf(out, "%s\n", src)
			}
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
	} else {
		src, err := api.Request("GET", fmt.Sprintf("/v2.0/%s/%s", orcidID, requestType), map[string]string{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		fmt.Fprintf(out, "%s\n", src)
	}
}
