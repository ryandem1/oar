package models

import (
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

// A Test represents point-in-time information about a test that occurred on a subject.
// The Summary can be thought of as a title, or a short description describing what the Test accomplished. A good
// rule-of-thumb is that if you cannot describe a test in a Summary, then the Test is probably too broad.
// The Outcome is the 'O' part of the OAR, it is the simple test binary and should remain that way with no ambiguity.
// The Analysis is the 'A' and will most likely be done after the initial Test upload.
// The Resolution is the 'R' and will most likely take place after the Analysis
// The Doc is a free form JSON document that can be used to store any sort of metadata about the Test
type Test struct {
	ID         int            `json:"id"`
	Summary    string         `json:"summary"`
	Outcome    Outcome        `json:"outcome"`
	Analysis   Analysis       `json:"analysis"`
	Resolution Resolution     `json:"resolution"`
	Doc        map[string]any `json:"doc"`
}

// Validate will ensure that a Test has a valid Outcome, Analysis, and Resolution and a non-blank Summary.
func (t *Test) Validate() error {
	validOutcomes := []Outcome{Passed, Failed}

	if !slices.Contains(validOutcomes, t.Outcome) {
		return fmt.Errorf("invalid outcome: '%s', must be one of outcomes: %s", t.Outcome, validOutcomes)
	}

	validAnalyses := []Analysis{NotAnalyzed}

	switch t.Outcome {
	case Passed:
		validAnalyses = append(validAnalyses, TrueNegative, FalseNegative)
	case Failed:
		validAnalyses = append(validAnalyses, TruePositive, FalsePositive)
	default:
		return fmt.Errorf("unrecognized test outcome: %s", t.Outcome)
	}

	if !slices.Contains(validAnalyses, t.Analysis) {
		return fmt.Errorf("invalid analysis: '%s', must be one of analyses: %s", t.Analysis, validAnalyses)
	}

	validResolutions := []Resolution{Unresolved, TicketCreated, QuickFix, KnownIssue, TestFixed, TestDisabled}
	if !slices.Contains(validResolutions, t.Resolution) {
		return fmt.Errorf(
			"invalid resolution: '%s', must be one of resultions: %s",
			t.Resolution,
			validResolutions,
		)
	}

	if len(strings.TrimSpace(t.Summary)) < 1 {
		return fmt.Errorf("summary cannot be blank")
	}

	return nil
}

// Clean will trim the whitespace around a Test's Summary
func (t *Test) Clean() {
	t.Summary = strings.TrimSpace(t.Summary)
	if t.Analysis == "" {
		t.Analysis = NotAnalyzed
	}

	if t.Resolution == "" {
		t.Resolution = Unresolved
	}
}
