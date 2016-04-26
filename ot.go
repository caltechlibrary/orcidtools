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
