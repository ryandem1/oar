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
// The Doc is a free form JSON document that can be used to store any sort of metadata about the Test
type Test struct {
	ID      int
	Summary string
	Outcome Outcome
	Doc     map[string]any
}

// Validate will ensure that a Test has a valid Outcome and a non-blank Summary.
func (t *Test) Validate() error {
	validOutcomes := []Outcome{Passed, Failed}

	if !slices.Contains(validOutcomes, t.Outcome) {
		return fmt.Errorf("invalid outcome: '%s', must be one of outcomes: %s", t.Outcome, validOutcomes)
	}

	if len(strings.TrimSpace(t.Summary)) < 1 {
		return fmt.Errorf("summary cannot be blank")
	}

	return nil
}

// Clean will trim the whitespace around a Test's Summary
func (t *Test) Clean() {
	t.Summary = strings.TrimSpace(t.Summary)
}
