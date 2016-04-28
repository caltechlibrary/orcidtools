package ot

import (
	"fmt"
	"log"
	"testing"
)

func TestEnvironmentSetup(t *testing.T) {
	api := New()
	fmt.Printf("DEBUG api.URL: %s\n", api.URL)
	if api == nil {
		log.Fatalf("Environment is not setup properly")
	}
}

func TestBasic(t *testing.T) {
	api := New()
	data, err := api.Login()
	if err != nil {
		t.Errorf("api.Login(), %s", err)
	}
	fmt.Printf("DEBUG data: %s\n", data)
}
