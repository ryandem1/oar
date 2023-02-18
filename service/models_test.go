package main

import (
	"testing"
)

var Fake = newFaker()

// TestValidTestsPassValidate will ensure that the Test.Validate() function correctly accepts valid Test objects.
func TestValidTestsPassValidate(t *testing.T) {
	validTests := multiple(15, Fake.test)

	for _, validTest := range validTests {
		err := validTest.Validate()
		if err != nil {
			t.Error(err)
		}
	}
}

// TestInvalidTestsDoNotPassValidate will ensure that the Test.Validate() does not accept invalid Test objects.
func TestInvalidTestsDoNotPassValidate(t *testing.T) {
	tests := multiple(15, Fake.test) // Start will valid tests and make them invalid

	invalidTest1 := tests[0]
	invalidTest1.Outcome = "Skipped" // OAR doesn't accept skipped tests

	invalidTest2 := tests[1]
	invalidTest2.ID = -5 // IDs should always be positive
}
