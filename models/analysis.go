package models

import (
	"fmt"
	"golang.org/x/exp/slices"
)

// TestAnalysis represents a single analysis of a test. A test can only have a single analysis at a time.
type TestAnalysis struct {
	TestID   int      `json:"testId"`
	Analysis Analysis `json:"analysis"`
	Comment  string   `json:"comment"`
}

// Validate will ensure that the analysis is a valid option according to the Test's Outcome
func (ta *TestAnalysis) Validate(test *Test) error {
	var validAnalyses []Analysis

	switch test.Outcome {
	case Passed:
		validAnalyses = []Analysis{TrueNegative, FalseNegative}
	case Failed:
		validAnalyses = []Analysis{TruePositive, FalsePositive}
	default:
		return fmt.Errorf("Unrecognized test outcome: %s", test.Outcome)
	}

	if !slices.Contains(validAnalyses, ta.Analysis) {
		return fmt.Errorf("invalid analysis: '%s', must be one of analyses: %s", ta.Analysis, validAnalyses)
	}

	return nil
}
