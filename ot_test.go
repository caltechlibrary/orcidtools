package ot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

func TestEnvironmentSetup(t *testing.T) {
	api := New()
	if api == nil {
		log.Fatalf("Environment is not setup properly")
	}
}

func TestLogin(t *testing.T) {
	api := New()
	data, err := api.Login()
	if err != nil {
		t.Errorf("api.Login(), %s", err)
	}
	if data.Error != "" {
		t.Errorf("api.Login() returned an error reply %s", data)
	}
	if data.AccessToken == "" {
		t.Errorf("api.Login() missing access_token %s", data)
	}
	if data.TokenType == "" {
		t.Errorf("api.Login() missing token_type %s", data)
	}
}

func TestGetAPI(t *testing.T) {
	var err error
	api := New()
	_, err = api.Login()
	if err != nil {
		t.Errorf("api.Login() error %s", err)
	}
	orcid := "0000-0002-2389-8429"
	_, err = api.GetBio(orcid)
	if err != nil {
		t.Errorf("api.GetBio(%q) error %s", orcid, err)
	}
	_, err = api.GetWorks(orcid)
	if err != nil {
		t.Errorf("api.GetWorks(%q) error %s", orcid, err)
	}
	_, err = api.GetProfile(orcid)
	if err != nil {
		t.Errorf("api.GetProfile(%q) error %s", orcid, err)
	}
}

// Test that orcid message examples can be unmarshalled
func TestORCIDMessage(t *testing.T) {
	var orcidMessage *OrcidMessage

	noError := func(err error, msg string, failNow bool) {
		if err != nil {
			t.Errorf("%s, %s", msg, err)
			if failNow == true {
				t.FailNow()
			}
		}
	}

	fname := path.Join("testdata", "0000-0003-0900-6903", "orcid-profile-message.json")
	src, err := ioutil.ReadFile(fname)
	noError(err, fmt.Sprintf("Expected to open %s", fname), true)

	err = json.Unmarshal(src, &orcidMessage)
	noError(err, "orcid-profile-message.json", false)

	fname = path.Join("testdata", "0000-0003-0900-6903", "orcid-bio-message.json")
	src, err = ioutil.ReadFile(fname)
	noError(err, fmt.Sprintf("Expected to open %s", fname), true)

	err = json.Unmarshal(src, &orcidMessage)
	noError(err, "orcid-bio-message.json", false)

	fname = path.Join("testdata", "0000-0003-0900-6903", "orcid-works-message.json")
	src, err = ioutil.ReadFile(fname)
	noError(err, fmt.Sprintf("Expected to open %s", fname), true)

	err = json.Unmarshal(src, &orcidMessage)
	noError(err, "orcid-works-message.json", false)
}

func setup() {
	os.Setenv("ORCID_API_URL", "https://pub.sandbox.orcid.org")
	os.Setenv("ORCID_CLIENT_ID", "APP-01XX65MXBF79VJGF")
	os.Setenv("ORCID_CLIENT_SECRET", "3a87028d-c84c-4d5f-8ad5-38a93181c9e1")
	os.Setenv("ORCID_ACCESS_TOKEN", "")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
