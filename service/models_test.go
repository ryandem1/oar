package main

import (
	"github.com/google/go-cmp/cmp"
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

// TestTest_Clean will ensure that the Test.Clean() function cleans a test and does not alter it in any unexpected ways
func TestTest_Clean(t *testing.T) {
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

// TestTest_Equal ensures that Test.Equal correctly identifies tests that are equal or not
func TestTest_Equal(t *testing.T) {
	t.Run("identical tests should be equal", func(t *testing.T) {
		test := Fake.test()
		copiedTest := test

		if !test.Equal(copiedTest) {
			t.Error("identical tests were marked as not equal")
		}
	})

	t.Run("different tests should be marked unequal", func(t *testing.T) {
		test := Fake.test()
		comparedTest := Fake.test()

		if test.Equal(comparedTest) {
			t.Error("different tests were marked equal")
		}
	})
}

// TestTest_Merge will check that a test can be successfully right-merged with another test
func TestTest_Merge(t *testing.T) {
	t.Run("valid test merge with valid test", func(t *testing.T) {
		targetTest := Fake.test()
		originalTargetTest := targetTest
		sourceTest := Fake.test()
		targetTest.Merge(sourceTest)
		oarDetailsAreEqual := sourceTest.Summary == targetTest.Summary &&
			sourceTest.Outcome == targetTest.Outcome &&
			sourceTest.Analysis == targetTest.Analysis &&
			sourceTest.Resolution == targetTest.Resolution

		if !oarDetailsAreEqual {
			t.Error("right merge failed")
		}

		for k, v := range sourceTest.Doc {
			if !cmp.Equal(targetTest.Doc[k], v) {
				t.Error("right merge on Doc failed")
			}
		}

		for k, v := range originalTargetTest.Doc {
			if !cmp.Equal(targetTest.Doc[k], v) {
				t.Error("right merge on Doc failed. source values did not get preserved")
			}
		}
	})

	t.Run("invalid test details do not get merged", func(t *testing.T) {
		targetTest := Fake.test()
		originalTargetTest := targetTest
		invalidSourceTest := &Test{
			ID:         Fake.testID(),
			Summary:    "",
			Outcome:    "",
			Analysis:   "",
			Resolution: "",
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc:        nil,
		}
		targetTest.Merge(invalidSourceTest)
		oarDetailsAreEqual := originalTargetTest.Summary == targetTest.Summary &&
			originalTargetTest.Outcome == targetTest.Outcome &&
			originalTargetTest.Analysis == targetTest.Analysis &&
			originalTargetTest.Resolution == targetTest.Resolution

		if !oarDetailsAreEqual || targetTest.Doc == nil {
			t.Error("target test got merged with invalid details")
		}
	})
}
