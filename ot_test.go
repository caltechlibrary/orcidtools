package ot

import (
	"fmt"
	"os"
	"testing"
)

func TestBasic(t *testing.T) {
	os.Setenv("ORCID_PUBLIC_API_URL", "https://sandbox.orcid.org")
	api := New()
	data, err := api.Login()
	if err != nil {
		t.Errorf("api.Login(), %s", err)
	}
	fmt.Printf("DEBUG data: %s\n", data)
}
