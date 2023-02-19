package main

import (
	"strings"
	"testing"
	"time"
)

// TestValidTestsPassValidate will ensure that the Test.Validate() function correctly accepts valid Test objects.
func TestValidTestsPassValidate(t *testing.T) {
	validTests := multiple(150, Fake.test)

	for _, validTest := range validTests {
		err := validTest.Validate()
		if err != nil {
			t.Error(err)
		}
	}
}

// TestInvalidTestsDoNotPassValidate will ensure that the Test.Validate() does not accept invalid Test objects.
func TestInvalidTestsDoNotPassValidate(t *testing.T) {
	validTests := multiple(5, Fake.test) // Start will valid tests and make them invalid

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
		"invalid outcome test":      validTests[0],
		"invalid analysis test":     validTests[1],
		"invalid resolution test":   validTests[2],
		"outcome/analysis mismatch": validTests[3],
		"empty summary test":        validTests[4],
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

// TestTestClean will ensure that the Test.Clean() function cleans a test and does not alter it in any unexpected ways
func TestTestClean(t *testing.T) {
	validTests := multiple(150, Fake.test)
	validTestsCopy := make([]*Test, len(validTests)) // Copy to validate the tests after the clean operation
	copy(validTestsCopy, validTests)

	t.Run("valid tests do not get altered", func(t *testing.T) {
		for i, validTest := range validTests {
			validTest.Clean()
			if validTest != validTestsCopy[i] {
				t.Error("clean, valid test got altered after Test.Clean()")
			}
		}
	})

	uncleanTest := &Test{
		ID:         Fake.testID(),
		Summary:    Fake.testSummary(),
		Outcome:    Passed,
		Analysis:   "",
		Resolution: "",
		Created:    time.Now(),
		Modified:   time.Now(),
		Doc:        nil,
	}
	t.Run("Empty Analysis and Resolution get converted to NotAnalyzed and Unresolved", func(t *testing.T) {
		uncleanTest.Clean()
		if uncleanTest.Analysis != NotAnalyzed || uncleanTest.Resolution != Unresolved {
			t.Error("test without analysis/resolution did not get set to NotAnalyzed/Unresolved after clean")
		}
	})

	uncleanTest = &Test{
		ID:         Fake.testID(),
		Summary:    "   " + Fake.testSummary() + "     ",
		Outcome:    Passed,
		Analysis:   NotAnalyzed,
		Resolution: Unresolved,
		Created:    time.Now(),
		Modified:   time.Now(),
		Doc:        nil,
	}
	t.Run("Test with unclean summary gets cleaned up", func(t *testing.T) {
		uncleanTest.Clean()
		if uncleanTest.Summary != strings.TrimSpace(uncleanTest.Summary) {
			t.Error("the unclean test did not have its summary cleaned")
		}
	})
}
