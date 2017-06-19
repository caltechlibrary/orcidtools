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
	src, err := api.Request("get", "/v1.2/"+testORCID+"/orcid-profile", map[string]string{})
	if err != nil {
		t.Errorf("request profile failed, %s", err)
		t.FailNow()
	}

	data := map[string]interface{}{}
	if err := json.Unmarshal(src, &data); err != nil {
		t.Errorf("Can't unmashall JSON response, %s", err)
		t.FailNow()
	}

	if val, ok := data["message-version"]; ok == true {
		if val.(string) != "1.2" {
			t.Errorf("expected 1.2, got %s\n", val)
		}
	} else {
		t.Errorf("missing message-version key")
		t.FailNow()
	}

	// FIXME: Need to finish writing the test
	/*
	   {
	     "message-version": "1.2",
	     "orcid-profile": {
	       "orcid": null,
	       "orcid-id": null,
	       "orcid-identifier": {
	         "value": null,
	         "uri": "http://orcid.org/0000-0003-0900-6903",
	         "path": "0000-0003-0900-6903",
	         "host": "orcid.org"
	       },
	       "orcid-deprecated": null,
	       "orcid-preferences": {
	         "locale": "EN"
	       },
	       "orcid-history": {
	         "creation-method": "DIRECT",
	         "completion-date": null,
	         "submission-date": {
	           "value": 1461871747241
	         },
	         "last-modified-date": {
	           "value": 1497891064296
	         },
	         "claimed": {
	           "value": true
	         },
	         "source": null,
	         "deactivation-date": null,
	         "verified-email": {
	           "value": true
	         },
	         "verified-primary-email": {
	           "value": true
	         },
	         "visibility": null
	       },
	       "orcid-bio": {
	         "personal-details": {
	           "given-names": {
	             "value": "Robert",
	             "visibility": null
	           },
	           "family-name": {
	             "value": "Doiel",
	             "visibility": null
	           },
	           "credit-name": {
	             "value": "R. S. Doiel",
	             "visibility": "PUBLIC"
	           },
	           "other-names": {
	             "other-name": [
	               {
	                 "value": "R. Doiel"
	               },
	               {
	                 "value": "R. S. Doiel"
	               },
	               {
	                 "value": "Robert Doiel"
	               },
	               {
	                 "value": "Robert S. Doiel"
	               }
	             ],
	             "visibility": "PUBLIC"
	           }
	         },
	         "biography": null,
	         "researcher-urls": {
	           "researcher-url": [
	             {
	               "url-name": {
	                 "value": "Personal website"
	               },
	               "url": {
	                 "value": "http://rsdoiel.github.io"
	               }
	             },
	             {
	               "url-name": {
	                 "value": "Work related Open Source Projects"
	               },
	               "url": {
	                 "value": "https://github.com/caltechlibrary"
	               }
	             },
	             {
	               "url-name": {
	                 "value": "Personal Open Source projects"
	               },
	               "url": {
	                 "value": "https://github.com/rsdoiel"
	               }
	             }
	           ],
	           "visibility": "PUBLIC"
	         },
	         "contact-details": {
	           "email": [
	             {
	               "value": "rsdoiel@caltech.edu",
	               "primary": true,
	               "current": true,
	               "verified": true,
	               "visibility": "PUBLIC",
	               "source": "0000-0003-0900-6903",
	               "source-client-id": null
	             }
	           ],
	           "address": {
	             "country": {
	               "value": "US",
	               "visibility": "PUBLIC"
	             }
	           }
	         },
	         "keywords": {
	           "keyword": [
	             {
	               "value": "Archives"
	             },
	             {
	               "value": "Developer"
	             },
	             {
	               "value": "Library"
	             },
	             {
	               "value": "Software"
	             }
	           ],
	           "visibility": "PUBLIC"
	         },
	         "external-identifiers": null,
	         "delegation": null,
	         "scope": null
	       },
	       "orcid-activities": {
	         "affiliations": {
	           "affiliation": [
	             {
	               "type": "EMPLOYMENT",
	               "department-name": "Library",
	               "role-title": "Software Developer",
	               "start-date": {
	                 "year": {
	                   "value": "2015"
	                 },
	                 "month": {
	                   "value": "04"
	                 },
	                 "day": {
	                   "value": "20"
	                 }
	               },
	               "end-date": null,
	               "organization": {
	                 "name": "California Institute of Technology",
	                 "address": {
	                   "city": "Pasadena",
	                   "region": "CA",
	                   "country": "US"
	                 },
	                 "disambiguated-organization": {
	                   "disambiguated-organization-identifier": "6469",
	                   "disambiguation-source": "RINGGOLD"
	                 }
	               },
	               "source": {
	                 "source-orcid": {
	                   "value": null,
	                   "uri": "http://orcid.org/0000-0003-0900-6903",
	                   "path": "0000-0003-0900-6903",
	                   "host": "orcid.org"
	                 },
	                 "source-client-id": null,
	                 "source-name": {
	                   "value": "R. S. Doiel"
	                 },
	                 "source-date": {
	                   "value": 1461872412194
	                 }
	               },
	               "created-date": {
	                 "value": 1461872412194
	               },
	               "last-modified-date": {
	                 "value": 1461878696569
	               },
	               "visibility": "PUBLIC",
	               "put-code": "1788050"
	             },
	             {
	               "type": "EMPLOYMENT",
	               "department-name": "ITS Web Services",
	               "role-title": "Software Developer",
	               "start-date": {
	                 "year": {
	                   "value": "1992"
	                 },
	                 "month": {
	                   "value": "01"
	                 },
	                 "day": null
	               },
	               "end-date": {
	                 "year": {
	                   "value": "2015"
	                 },
	                 "month": {
	                   "value": "04"
	                 },
	                 "day": null
	               },
	               "organization": {
	                 "name": "University of Southern California",
	                 "address": {
	                   "city": "Los Angeles",
	                   "region": "CA",
	                   "country": "US"
	                 },
	                 "disambiguated-organization": {
	                   "disambiguated-organization-identifier": "5116",
	                   "disambiguation-source": "RINGGOLD"
	                 }
	               },
	               "source": {
	                 "source-orcid": {
	                   "value": null,
	                   "uri": "http://orcid.org/0000-0003-0900-6903",
	                   "path": "0000-0003-0900-6903",
	                   "host": "orcid.org"
	                 },
	                 "source-client-id": null,
	                 "source-name": {
	                   "value": "R. S. Doiel"
	                 },
	                 "source-date": {
	                   "value": 1461872486190
	                 }
	               },
	               "created-date": {
	                 "value": 1461872486190
	               },
	               "last-modified-date": {
	                 "value": 1462302682388
	               },
	               "visibility": "PUBLIC",
	               "put-code": "1788055"
	             },
	             {
	               "type": "EDUCATION",
	               "department-name": null,
	               "role-title": "M.A. Humanities",
	               "start-date": null,
	               "end-date": {
	                 "year": {
	                   "value": "2010"
	                 },
	                 "month": null,
	                 "day": null
	               },
	               "organization": {
	                 "name": "Mount Saint Mary's College",
	                 "address": {
	                   "city": "Los Angeles",
	                   "region": "CA",
	                   "country": "US"
	                 },
	                 "disambiguated-organization": {
	                   "disambiguated-organization-identifier": "5144",
	                   "disambiguation-source": "RINGGOLD"
	                 }
	               },
	               "source": {
	                 "source-orcid": {
	                   "value": null,
	                   "uri": "http://orcid.org/0000-0003-0900-6903",
	                   "path": "0000-0003-0900-6903",
	                   "host": "orcid.org"
	                 },
	                 "source-client-id": null,
	                 "source-name": {
	                   "value": "R. S. Doiel"
	                 },
	                 "source-date": {
	                   "value": 1461872058612
	                 }
	               },
	               "created-date": {
	                 "value": 1461872058612
	               },
	               "last-modified-date": {
	                 "value": 1462302674625
	               },
	               "visibility": "PUBLIC",
	               "put-code": "1788042"
	             },
	             {
	               "type": "EDUCATION",
	               "department-name": null,
	               "role-title": "B. S. Computer Science",
	               "start-date": null,
	               "end-date": {
	                 "year": {
	                   "value": "2004"
	                 },
	                 "month": null,
	                 "day": null
	               },
	               "organization": {
	                 "name": "University of Southern California",
	                 "address": {
	                   "city": "Los Angeles",
	                   "region": "CA",
	                   "country": "US"
	                 },
	                 "disambiguated-organization": {
	                   "disambiguated-organization-identifier": "5116",
	                   "disambiguation-source": "RINGGOLD"
	                 }
	               },
	               "source": {
	                 "source-orcid": {
	                   "value": null,
	                   "uri": "http://orcid.org/0000-0003-0900-6903",
	                   "path": "0000-0003-0900-6903",
	                   "host": "orcid.org"
	                 },
	                 "source-client-id": null,
	                 "source-name": {
	                   "value": "R. S. Doiel"
	                 },
	                 "source-date": {
	                   "value": 1461872112462
	                 }
	               },
	               "created-date": {
	                 "value": 1461872112462
	               },
	               "last-modified-date": {
	                 "value": 1462302675853
	               },
	               "visibility": "PUBLIC",
	               "put-code": "1788043"
	             },
	             {
	               "type": "EDUCATION",
	               "department-name": null,
	               "role-title": "A.A. Humanities",
	               "start-date": null,
	               "end-date": {
	                 "year": {
	                   "value": "1989"
	                 },
	                 "month": null,
	                 "day": null
	               },
	               "organization": {
	                 "name": "College of the Canyons",
	                 "address": {
	                   "city": "Santa Clarita",
	                   "region": "CA",
	                   "country": "US"
	                 },
	                 "disambiguated-organization": {
	                   "disambiguated-organization-identifier": "17072",
	                   "disambiguation-source": "RINGGOLD"
	                 }
	               },
	               "source": {
	                 "source-orcid": {
	                   "value": null,
	                   "uri": "http://orcid.org/0000-0003-0900-6903",
	                   "path": "0000-0003-0900-6903",
	                   "host": "orcid.org"
	                 },
	                 "source-client-id": null,
	                 "source-name": {
	                   "value": "R. S. Doiel"
	                 },
	                 "source-date": {
	                   "value": 1461872250080
	                 }
	               },
	               "created-date": {
	                 "value": 1461872250080
	               },
	               "last-modified-date": {
	                 "value": 1462302677090
	               },
	               "visibility": "PUBLIC",
	               "put-code": "1788046"
	             }
	           ]
	         },
	         "orcid-works": {
	           "orcid-work": [
	             {
	               "put-code": "25388820",
	               "work-title": {
	                 "title": {
	                   "value": "Digital Video and Internet 2: Growing Up Together"
	                 },
	                 "subtitle": null,
	                 "translated-title": null
	               },
	               "journal-title": {
	                 "value": "Syllabus Magazine"
	               },
	               "short-description": null,
	               "work-citation": {
	                 "work-citation-type": "BIBTEX",
	                 "citation": "@Article{Doiel2000,\n  author    = {Doiel, R. and Lunsten, A.},\n  title     = {Digital Video and Internet 2: Growing Up Together},\n  journal   = {Syllabus Magazine},\n  year      = {2000},\n  publisher = {1105 Media Inc, Ed-Tech Group.},\n  url       = {https://campustechnology.com/articles/2001/07/digital-video-and-internet2-growing-up-together.aspx},\n}\n"
	               },
	               "work-type": "JOURNAL_ARTICLE",
	               "publication-date": {
	                 "year": {
	                   "value": "2001"
	                 },
	                 "month": {
	                   "value": "07"
	                 },
	                 "day": null,
	                 "media-type": null
	               },
	               "work-external-identifiers": null,
	               "url": {
	                 "value": "https://campustechnology.com/articles/2001/07/digital-video-and-internet2-growing-up-together.aspx"
	               },
	               "work-contributors": null,
	               "work-source": null,
	               "source": {
	                 "source-orcid": {
	                   "value": null,
	                   "uri": "http://orcid.org/0000-0003-0900-6903",
	                   "path": "0000-0003-0900-6903",
	                   "host": "orcid.org"
	                 },
	                 "source-client-id": null,
	                 "source-name": {
	                   "value": "R. S. Doiel"
	                 },
	                 "source-date": {
	                   "value": 1468447356184
	                 }
	               },
	               "created-date": {
	                 "value": 1468447356184
	               },
	               "last-modified-date": {
	                 "value": 1468447356184
	               },
	               "language-code": "en",
	               "country": {
	                 "value": "US",
	                 "visibility": "PUBLIC"
	               },
	               "visibility": "PUBLIC"
	             },
	             {
	               "put-code": "25388814",
	               "work-title": {
	                 "title": {
	                   "value": "Getting Ready for Internet2"
	                 },
	                 "subtitle": null,
	                 "translated-title": null
	               },
	               "journal-title": {
	                 "value": "Syllabus Magazine"
	               },
	               "short-description": null,
	               "work-citation": {
	                 "work-citation-type": "BIBTEX",
	                 "citation": "@Article{Doiel2001,\n  author    = {Doiel, R. and Lunsten, A.},\n  title     = {Getting Ready for Internet2},\n  journal   = {Syllabus Magazine},\n  year      = {2001},\n  month     = apr,\n  publisher = {1105 Media Inc, Ed-Tech Group.},\n  url       = {https://campustechnology.com/articles/2001/04/getting-ready-for-internet2.aspx},\n}\n"
	               },
	               "work-type": "JOURNAL_ARTICLE",
	               "publication-date": {
	                 "year": {
	                   "value": "2001"
	                 },
	                 "month": {
	                   "value": "04"
	                 },
	                 "day": null,
	                 "media-type": null
	               },
	               "work-external-identifiers": null,
	               "url": {
	                 "value": "https://campustechnology.com/articles/2001/04/getting-ready-for-internet2.aspx"
	               },
	               "work-contributors": null,
	               "work-source": null,
	               "source": {
	                 "source-orcid": {
	                   "value": null,
	                   "uri": "http://orcid.org/0000-0003-0900-6903",
	                   "path": "0000-0003-0900-6903",
	                   "host": "orcid.org"
	                 },
	                 "source-client-id": null,
	                 "source-name": {
	                   "value": "R. S. Doiel"
	                 },
	                 "source-date": {
	                   "value": 1468447392830
	                 }
	               },
	               "created-date": {
	                 "value": 1468447392830
	               },
	               "last-modified-date": {
	                 "value": 1468447392836
	               },
	               "language-code": "en",
	               "country": {
	                 "value": "US",
	                 "visibility": "PUBLIC"
	               },
	               "visibility": "PUBLIC"
	             },
	             {
	               "put-code": "25388837",
	               "work-title": {
	                 "title": {
	                   "value": "Internet 2: better, stronger, faster"
	                 },
	                 "subtitle": null,
	                 "translated-title": null
	               },
	               "journal-title": {
	                 "value": "CNET/Builder.com"
	               },
	               "short-description": null,
	               "work-citation": {
	                 "work-citation-type": "BIBTEX",
	                 "citation": "@Article{Doiel1999,\n  author  = {Doiel, R. and Lunsten, A.},\n  title   = {Internet 2: better, stronger, faster},\n  journal = {CNET/Builder.com},\n  year    = {1999},\n}\n"
	               },
	               "work-type": "JOURNAL_ARTICLE",
	               "publication-date": {
	                 "year": {
	                   "value": "1999"
	                 },
	                 "month": null,
	                 "day": null,
	                 "media-type": null
	               },
	               "work-external-identifiers": null,
	               "url": null,
	               "work-contributors": null,
	               "work-source": null,
	               "source": {
	                 "source-orcid": {
	                   "value": null,
	                   "uri": "http://orcid.org/0000-0003-0900-6903",
	                   "path": "0000-0003-0900-6903",
	                   "host": "orcid.org"
	                 },
	                 "source-client-id": null,
	                 "source-name": {
	                   "value": "R. S. Doiel"
	                 },
	                 "source-date": {
	                   "value": 1468447657846
	                 }
	               },
	               "created-date": {
	                 "value": 1468447657846
	               },
	               "last-modified-date": {
	                 "value": 1468447657846
	               },
	               "language-code": "en",
	               "country": {
	                 "value": "US",
	                 "visibility": "PUBLIC"
	               },
	               "visibility": "PUBLIC"
	             },
	             {
	               "put-code": "25388835",
	               "work-title": {
	                 "title": {
	                   "value": "SMIL, web sites of the future"
	                 },
	                 "subtitle": null,
	                 "translated-title": null
	               },
	               "journal-title": {
	                 "value": "CNET/Builder.com"
	               },
	               "short-description": null,
	               "work-citation": {
	                 "work-citation-type": "BIBTEX",
	                 "citation": "@Article{Doiel1999a,\n  author  = {Doiel, R. and Lunsten, A.},\n  title   = {SMIL, web sites of the future},\n  journal = {CNET/Builder.com},\n  year    = {1999},\n}\n"
	               },
	               "work-type": "JOURNAL_ARTICLE",
	               "publication-date": {
	                 "year": {
	                   "value": "1999"
	                 },
	                 "month": null,
	                 "day": null,
	                 "media-type": null
	               },
	               "work-external-identifiers": null,
	               "url": null,
	               "work-contributors": null,
	               "work-source": null,
	               "source": {
	                 "source-orcid": {
	                   "value": null,
	                   "uri": "http://orcid.org/0000-0003-0900-6903",
	                   "path": "0000-0003-0900-6903",
	                   "host": "orcid.org"
	                 },
	                 "source-client-id": null,
	                 "source-name": {
	                   "value": "R. S. Doiel"
	                 },
	                 "source-date": {
	                   "value": 1468447599144
	                 }
	               },
	               "created-date": {
	                 "value": 1468447599144
	               },
	               "last-modified-date": {
	                 "value": 1468447599144
	               },
	               "language-code": "en",
	               "country": {
	                 "value": "US",
	                 "visibility": "PUBLIC"
	               },
	               "visibility": "PUBLIC"
	             }
	           ],
	           "scope": null
	         },
	         "funding-list": null
	       },
	       "orcid-internal": null,
	       "type": "USER",
	       "group-type": null,
	       "client-type": null
	     },
	     "orcid-search-results": null,
	     "error-desc": null
	   }
	*/
}
