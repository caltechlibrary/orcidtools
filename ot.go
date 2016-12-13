//
// Package ot is a library for working with the ORCID API
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2016, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of ot nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package ot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// Version of ot package
	Version = "v0.0.01"

	// License string suitable to populate via cli
	LicenseText = `
%s %s

Copyright (c) 2016, Caltech
All rights not granted herein are expressly reserved by Caltech.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of ot nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`
)

type API struct {
	u      *url.URL
	id     string
	secret string
	token  string

	// Timeout is the client time out period, default is 10 seconds
	Timeout time.Duration
}

func New(apiURL, clientID, clientSecret string) (*API, error) {
	u, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}
	return &API{
		u:       u,
		id:      clientID,
		secret:  clientSecret,
		token:   "",
		Timeout: 10 * time.Second,
	}, nil
}

func (api *API) login() error {
	if api.token == "" {
		client := &http.Client{
			Timeout: api.Timeout,
		}
		u := api.u
		u.Path = "/oauth/token"

		// OAuth2 authentication is usually done with a POST, need to setup the form values
		// and URL encode the results.
		payload := map[string]string{
			"client_id":     api.id,
			"client_secret": api.secret,
			"scope":         "/read-public",
			"grant_type":    "client_credentials",
		}
		form := u.Query()
		for key, value := range payload {
			form.Add(key, value)
		}

		// OK, we're ready to setup our request, send it and get on our way.
		req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
		if err != nil {
			return err
		}
		// Need to set the mime type for the content we're sending to the API
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		// Get the text response for API
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		src, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		// We have a throw away response object from authenticating in
		data := &struct {
			AccessToken  string `json:"access_token"`
			Bearer       string `json:"bearer"`
			RefreshToken string `json:"refresh_token"`
			ExpiresIn    int    `json:"expires_in"`
			Scope        string `json:"scope"`
		}{}
		err = json.Unmarshal(src, &data)
		if err != nil {
			return err
		}
		if data.AccessToken == "" {
			return fmt.Errorf("Could not decode access token from %q, %+v", src, data)
		}
		api.token = data.AccessToken
	}
	return nil
}

// Request contacts the ORCID API and returns the full read response body, and error
func (api *API) Request(method, docPath string, payload map[string]string) ([]byte, error) {
	var (
		req *http.Request
		err error
	)
	// Create a http client
	client := &http.Client{
		Timeout: api.Timeout,
	}

	// NOT: if api.token not set we should just go ahead and login.
	if api.token == "" {
		err := api.login()
		if err != nil {
			return nil, err
		}
	}

	// NOTE: we want a copy the URL in API object and update copy with the docPath
	u := api.u
	u.Path = docPath

	// NOTE: Based the HTTP method we want, we build our request appropriately
	switch strings.ToUpper(method) {
	case "GET":
		req, err = http.NewRequest("GET", u.String(), nil)
		if err != nil {
			return nil, err
		}
		// NOTE: If we've authenticated we need to path the auth token
		if len(api.token) > 0 {
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api.token))
		}
		// NOTE: We need to indicate the format we want
		req.Header.Add("Accept", "application/json")

		// NOTE: Build our payload to pass in the URL since this is a GET
		qry := req.URL.Query()
		for key, value := range payload {
			qry.Add(key, value)
		}
		req.URL.RawQuery = qry.Encode()
	default:
		return nil, fmt.Errorf("Do not know how to make a %s request", method)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	src, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return src, nil
}
