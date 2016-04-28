//
// Package ot wraps the data structures and services returned by the Orcid API.
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
package ot

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

const (
	// Version provides the version number for github.com/caltechlibrary/ot
	Version = "0.0.0"
)

// OrcidIdentifier specifies the field orcid-identifier fields in an API response
type OrcidIdentifier struct {
	XMLName xml.Name `json:"-"`
	URI     string   `xml:"uri" json:"uri"`
	Path    string   `xml:"path" json:"path"`
	Host    string   `xml:"host" json:"host"`
}

// PersonalDetails specifies the field personal-details fields in an API response
type PersonalDetails struct {
	XMLName    xml.Name `json:"-"`
	GiveNames  string   `xml:"given-names" json:"given-names"`
	FamilyName string   `xml:"family-name" json:"family-name"`
}

// ContactDetails specifies the field contact-details fields in an API response
type ContactDetails struct {
	XMLName xml.Name `json:"-"`
	EMail   string   `xml:"email" json:"email"`
}

// OrcidBio specifies the field orcid-bio fields in an API response
type OrcidBio struct {
	XMLName         xml.Name         `json:"-"`
	PersonalDetails *PersonalDetails `xml:"personal-details" json:"personal-details"`
	ContactDetails  *ContactDetails  `xml:"contact-details" json:"contact-details"`
}

// OrcidProfile specifies the field orcid-profile fields in an API response
type OrcidProfile struct {
	XMLName         xml.Name         `json:"-"`
	OrcidIdentifier *OrcidIdentifier `xml:"orcid-identifier" json:"orcid-identifier"`
	OrcidBio        *OrcidBio        `xml:"orcid-bio" json:"orcid-bio"`
}

// OrcidSearchResult specifies the individual fields of a single orcid-search-result fields in an API response
type OrcidSearchResult struct {
	XMLName        xml.Name      `json:"-"`
	RelevancyScore float64       `xml:"relevancy-score" json:"relevancy-score"`
	OrchidProfile  *OrcidProfile `xml:"orcid-profile" json:"orcid-profile"`
}

// OrcidSearchResults specifies the field orcid-search-results fields in an API response
type OrcidSearchResults struct {
	XMLName           xml.Name             `json:"-"`
	OrcidSearchResult []*OrcidSearchResult `xml:"orcid-search-result" json:"orcid-search-result"`
}

// OrcidMessage specifies the field orcid-message fields in an API response
type OrcidMessage struct {
	XMLName            xml.Name            `json:"-"`
	MessageVersion     float64             `xml:"message-version" json:"message-version"`
	OrcidSearchResults *OrcidSearchResults `xml:"orcid-search-results" json:"orcid-search-results"`
}

// OrcidAPI holds the details for authenticating against orcid.org's public API
type OrcidAPI struct {
	URL          *url.URL
	ClientID     string
	ClientSecret string
	AccessToken  string
	IsAuth       bool
}

// New creates a new Public API object populated based on any environment variables set.
func New() *OrcidAPI {
	var (
		err error
		ok  bool
	)
	ok = true
	apiURL := os.Getenv("ORCID_API_URL")
	if apiURL == "" {
		ok = false
		log.Println("ORCID_API_URL missing")
	}
	clientID := os.Getenv("ORCID_CLIENT_ID")
	if clientID == "" {
		ok = false
		log.Println("ORCID_CLIENT_ID missing")
	}
	clientSecret := os.Getenv("ORCID_CLIENT_SECRET")
	if clientSecret == "" {
		ok = false
		log.Println("ORCID_CLIENT_SECRET missing")
	}
	u, err := url.Parse(apiURL)
	if err != nil {
		ok = false
		fmt.Printf("ORCID_API_URL malformed %s, %s", apiURL, err)
	}
	if ok == false {
		return nil
	}
	api := new(OrcidAPI)
	api.URL = u
	api.ClientID = clientID
	api.ClientSecret = clientSecret
	return api
}

// Authorize connects with the API to get an authorization token.

// Login connects with the API and gets the necessary access token.
func (api *OrcidAPI) Login() ([]byte, error) {
	//curl -i -L -H "Accept: application/json" -d "client_id=APP-01XX65MXBF79VJGF" \
	//                     -d "client_secret=3a87028d-c84c-4d5f-8ad5-38a93181c9e1" \
	//                     -d "scope=/read-public" \
	//                     -d "grant_type=client_credentials" \
	//                     "https://pub.sandbox.orcid.org/oauth/token"
	u := api.URL
	urlPath := path.Join(u.Path, "oauth", "token")
	u.Path = urlPath
	q := u.Query()
	q.Set("client_id", api.ClientID)
	q.Set("client_secret", api.ClientSecret)
	q.Set("scope", "/read-public")
	q.Set("grant_type", "client_credentials")
	u.RawQuery = q.Encode()
	uri := u.String()
	fmt.Printf("DEBUG u.String() -> %s\n", uri)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requested %s, %s", uri, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http error %s, %s", uri, res.Status)
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("content can't be read %s, %s", uri, err)
	}
	fmt.Printf("DEBUG content: %s\n", content)
	return content, nil
}
