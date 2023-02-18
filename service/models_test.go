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
	validTests := multiple(15, Fake.test) // Start will valid tests and make them invalid

	// Alter tests to make them invalid (each test can be thought of as its own scenario)
	validTests[0].Outcome = "Skipped" // OAR doesn't accept skipped tests
	validTests[1].Analysis = "InvalidAnalysis"
	validTests[2].Resolution = "InvalidResolution"

	if validTests[3].Outcome == "Passed" {
		validTests[3].Analysis = TruePositive // Passed tests cannot have a "positive" analysis
	} else {
		validTests[3].Analysis = TrueNegative // Failed tests cannot have a "Negative" analysis
	}

	validTests[4].Summary = ""

	invalidTests := map[string]*Test{
		"Invalid Outcome Test":      validTests[0],
		"Invalid Analysis Test":     validTests[1],
		"Invalid Resolution Test":   validTests[2],
		"Outcome/Analysis Mismatch": validTests[3],
		"Empty Summary Test":        validTests[4],
	}

	for testName, invalidTest := range invalidTests {
		t.Run(testName, func(t *testing.T) {
			err := invalidTest.Validate()
			if err == nil {
				t.Errorf("Scenario %s failed, invalid test did not throw error", testName)
			}
		})
	}
}
