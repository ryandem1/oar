package models

import (
	"fmt"
	"golang.org/x/exp/slices"
)

// TestAnalysis represents a single analysis of a test. A test can only have a single analysis at a time.
type TestAnalysis struct {
	TestID   int      `json:"testId"`
	Analysis Analysis `json:"analysis"`
}

// Validate will ensure that the analysis is a valid option
func (ta *TestAnalysis) Validate() error {
	validAnalyses := []Analysis{TruePositive, FalsePositive, TrueNegative, FalseNegative}

	if !slices.Contains(validAnalyses, ta.Analysis) {
		return fmt.Errorf("invalid analysis: '%s', must be one of analyses: %s", ta.Analysis, validAnalyses)
	}

	return nil
}
