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
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

const (
	// Version provides the version number for github.com/caltechlibrary/ot
	Version = "0.0.0"
)

// OrcidProfile specifies the field orcid-profile fields in an API response
type OrcidProfile struct {
	XMLName          xml.Name                `json:"-"`
	ORCID            *string                 `xml:"orcid" json:"orcid"`
	OrcidID          *string                 `xml:"orcid-id" json:"orcid-id"`
	OrcidIdentifier  *OrcidIdentifier        `xml:"orcid-identifier" json:"orcid-identifier"`
	OrcidBio         *OrcidBio               `xml:"orcid-bio" json:"orcid-bio"`
	OrcidDepreciated *interface{}            `xml:"orcid-deprecated" json:"orcid-deprecated"`
	OrcidPreferences *map[string]interface{} `xml:"orcid-preferences" json:"orcid-preferences"`
	OrcidHistory     *map[string]interface{} `xml:"orcid-history" json:"orcid-history"`
	OrcidActivities  *interface{}            `xml:"orcid-activities" json:"orcid-activities"`
	OrcidInternal    *interface{}            `xml:"orcid-internal" json:"orcid-internal"`
	Type             *string                 `xml:"type" json:"type"`
	GroupType        *interface{}            `xml:"group-type" json:"group-type"`
	ClientType       *interface{}            `xml:"client-type" json:"client-type"`
}

// OrcidIdentifier specifies the field orcid-identifier fields in an API response
type OrcidIdentifier struct {
	XMLName xml.Name     `json:"-"`
	URI     string       `xml:"uri" json:"uri"`
	Path    string       `xml:"path" json:"path"`
	Host    string       `xml:"host" json:"host"`
	Value   *interface{} `xml:"value,omitempty" json:"value,omitempty"`
}

// OrcidBio specifies the field orcid-bio fields in an API response
type OrcidBio struct {
	XMLName         xml.Name         `json:"-"`
	PersonalDetails *PersonalDetails `xml:"personal-details" json:"personal-details"`
	Biography       *string          `xml:"biography,omitempty" json:"biography,omitempty"`
	ContactDetails  *ContactDetails  `xml:"contact-details" json:"contact-details"`
}

type Name struct {
	Value      string  `xml:"value,omitempty" json:"value,omitempty"`
	Visibility *string `xml:"visibility,omitempty" json:"visibility,omitempty"`
}

// PersonalDetails specifies the field personal-details fields in an API response
type PersonalDetails struct {
	XMLName    xml.Name `json:"-"`
	GiveNames  *Name    `xml:"given-names,omitempty" json:"given-names,omitempty"`
	FamilyName *Name    `xml:"family-name,omitempty" json:"family-name,omitempty"`
	CreditName *Name    `xml:"credit-name,omitempty"`
	OtherNames *[]Name  `xml:"other-names,omitempty" json:"other-names,omitempty"`
}

// ContactDetails specifies the field contact-details fields in an API response
type ContactDetails struct {
	XMLName xml.Name `json:"-"`
	EMail   string   `xml:"email" json:"email"`
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
	MessageVersion     *string             `xml:"message-version" json:"message-version"`
	OrcidSearchResults *OrcidSearchResults `xml:"orcid-search-results" json:"orcid-search-results"`
	OrcidProfile       *OrcidProfile       `xml:"orcid-profile" json:"orcid-profile"`
	Error              *string             `xml:"error,omitempty" json:"error,omitempty"`
	ErrorDesc          *string             `xml:"error-desc" json:"error-desc"`
}

// LoginResponseMessage structure that can hold a success or failing login response body
type LoginResponseMessage struct {
	AccessToken      string `json:"access_token,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	ExpiresIn        int    `json:"json,omitempty"`
	Scope            string `json:"scope,omitempty"`
	ORCID            string `json:"orcid,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

// OrcidAPI holds the details for authenticating against orcid.org's public API
type OrcidAPI struct {
	URL          *url.URL
	ClientID     string
	ClientSecret string
	ApiVersion   string
	AccessToken  string
	TokenType    string
	RefreshToken string
	ExpiresIn    int
	Scope        string
	ORCID        string
	IsAuth       bool
}

// String representation of an OrcidMessage
func (om *OrcidMessage) String() string {
	src, _ := json.Marshal(om)
	return fmt.Sprintf("%s", src)
}

// String representation of the API info
func (api *OrcidAPI) String() string {
	secret := ""
	if api.ClientSecret != "" {
		secret = "***************"
	}
	return fmt.Sprintf(`
URL: %s
ClientID: %s
ClientSecret: %s
`, api.URL.String(), api.ClientID, secret)
}

// String representation of OrcidIdentifier
func (oi *OrcidIdentifier) String() string {
	src, _ := json.Marshal(oi)
	return fmt.Sprintf("%s", src)
}

// String representation of OrcidProfile
func (op *OrcidProfile) String() string {
	src, _ := json.Marshal(op)
	return fmt.Sprintf("%s", src)
}

// String representation of LoginResponseMessage
func (lrm *LoginResponseMessage) String() string {
	src, _ := json.Marshal(lrm)
	return fmt.Sprintf("%s", src)
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
	apiVersion := os.Getenv("ORCID_API_VERSION")
	if apiVersion == "" {
		apiVersion = "v1.2"
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
	api.ApiVersion = apiVersion
	return api
}

// Login connects with the API and gets the necessary access token.
func (api *OrcidAPI) Login() (*LoginResponseMessage, error) {
	u := api.URL
	urlPath := path.Join(u.Path, "oauth", "token")
	u.Path = urlPath
	form := url.Values{}
	form.Add("client_id", api.ClientID)
	form.Add("client_secret", api.ClientSecret)
	form.Add("scope", "/read-public")
	form.Add("grant_type", "client_credentials")

	client := &http.Client{}
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf(`http.NewRequest("POST", %q, %q)`, u.String(), strings.NewReader(form.Encode()))
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requested %s, %s", u, err)
	}
	defer res.Body.Close()
	api.IsAuth = false
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf(`content can't be read %s, status code:%d, error message: %q`, u.String(), res.StatusCode, err)
	}
	var data *LoginResponseMessage
	err = json.Unmarshal(content, &data)
	if err != nil {
		return data, fmt.Errorf(`content can't be read %s, status code:%d, error message: %q`, u.String(), res.StatusCode, err)
	}
	if res.StatusCode != 200 {
		return data, fmt.Errorf(`status code:%d, error message: %q`, res.StatusCode, err)
	}
	api.AccessToken = data.AccessToken
	api.TokenType = data.TokenType
	api.RefreshToken = data.RefreshToken
	api.ExpiresIn = data.ExpiresIn
	api.Scope = data.Scope
	api.ORCID = data.ORCID
	api.IsAuth = true
	return data, nil
}

// Get returns a response body for a given ORCID message request
func (api *OrcidAPI) Get(p, orcid string) (*OrcidMessage, error) {
	u := api.URL
	urlPath := path.Join(api.ApiVersion, orcid, p) + "/"
	u.Path = urlPath
	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf(`http.NewRequest("GET", %q, nil)`, u.String())
	}
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", api.TokenType, api.AccessToken))
	req.Header.Set("Accept", "application/json")
	// 	fmt.Printf(`
	// curl -L -H "Authorization: %s %s" -H "Accept: application/json" %s | jq
	// `, api.TokenType, api.AccessToken, u.String())

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requested %s, %s", u, err)
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf(`content can't be read %s, status code:%d, error message: %s`, u.String(), res.StatusCode, err)
	}
	var orcidMessage *OrcidMessage
	err = json.Unmarshal(content, &orcidMessage)
	if err != nil {
		return nil, fmt.Errorf(`cannot unmarshal content: %s`, content)
	}
	if res.StatusCode != 200 {
		return orcidMessage, fmt.Errorf(`status code %d returned`, res.StatusCode)
	}
	return orcidMessage, nil
}

// GetBio return an ORCID Bio
func (api *OrcidAPI) GetBio(orcid string) (*OrcidMessage, error) {
	return api.Get("/orcid-bio", orcid)
}

// GetWorks return an ORCID Works
func (api *OrcidAPI) GetWorks(orcid string) (*OrcidMessage, error) {
	return api.Get("/orcid-works", orcid)
}

// GetProfile return an ORCID Profile
func (api *OrcidAPI) GetProfile(orcid string) (*OrcidMessage, error) {
	return api.Get("/orcid-profile", orcid)
}
