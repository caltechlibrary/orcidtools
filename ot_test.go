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
	"testing"
)

func TestAPI(t *testing.T) {
	apiURL := "https://pub.sandbox.orcid.org"
	clientID := "APP-01XX65MXBF79VJGF"
	clientSecret := "3a87028d-c84c-4d5f-8ad5-38a93181c9e1"
	testORCID := "0000-0003-0900-6903"

	// Test setup
	api, err := New(apiURL, clientID, clientSecret)
	if err != nil {
		t.Errorf("Can't create API, %s", err)
		t.FailNow()
	}
	if api == nil {
		t.Errorf("API shouldn't be nil")
		t.FailNow()
	}

	// Test internal login method
	err = api.login()
	if err != nil {
		t.Errorf("Can't authenticate, %s", err)
		t.FailNow()
	}
	api.token = ""

	// Test request method
	data, err := api.Request("get", "/v1.2/"+testORCID+"/orcid-profile", map[string]string{})
	if err != nil {
		t.Errorf("request profile failed, %s", err)
		t.FailNow()
	}
	t.Fail() // FIXME: Need to finish writing the test
	t.Logf("%s\n", data)
	//fmt.Printf("DEBUG Data %s\n", data)
}
