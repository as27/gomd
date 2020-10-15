package main

import (
	"reflect"
	"testing"
)

func TestCmdAutocompleteEmptyString(t *testing.T) {
	a := newApp()
	result := a.cmdAutocomplete("")
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

func TestCmdAutocompleteSingleChar(t *testing.T) {
	a := newApp()
	result := a.cmdAutocomplete("z")
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

func TestCmdAutocompleteNomatch(t *testing.T) {
	a := newApp()
	result := a.cmdAutocomplete("wqoersndflisadflsbzxjchbkjsdabajsbfjkas")
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}
}

func TestCmdAutocompletePrefix(t *testing.T) {
	a := newApp()
	a.commands = []string{"foobaz", "foobar", "foobuh"}
	expected := []string{"foobaz", "foobar"}
	result := a.cmdAutocomplete("fooba")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
