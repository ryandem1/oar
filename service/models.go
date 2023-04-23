package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/slices"
	"strings"
	"time"
)

// A Test represents point-in-time information about a test that occurred on a subject.
// The Summary can be thought of as a title, or a short description describing what the Test accomplished. A good
// rule-of-thumb is that if you cannot describe a test in a Summary, then the Test is probably too broad.
// The Outcome is the 'O' part of the OAR, it is the simple test binary and should remain that way with no ambiguity.
// The Analysis is the 'A' and will most likely be done after the initial Test upload.
// The Resolution is the 'R' and will most likely take place after the Analysis
// The Doc is a free form JSON document that can be used to store any sort of metadata about the Test
type Test struct {
	ID         uint64         `json:"id"`
	Summary    string         `json:"summary"`
	Outcome    Outcome        `json:"outcome"`
	Analysis   Analysis       `json:"analysis"`
	Resolution Resolution     `json:"resolution"`
	Created    time.Time      `json:"created"`
	Modified   time.Time      `json:"modified"`
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

	validResolutions := []Resolution{Unresolved, NotNeeded, TicketCreated, QuickFix, KnownIssue, TestFixed, TestDisabled}
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

// Merge will right-merge the current test instance with a different instance of a test. All values that are in the test
// to be merged with will be preferred.
func (t *Test) Merge(testPatch *Test) {
	if testPatch.Summary != "" {
		t.Summary = testPatch.Summary
	}
	if testPatch.Outcome != "" {
		t.Outcome = testPatch.Outcome
	}
	if testPatch.Analysis != "" {
		t.Analysis = testPatch.Analysis
	}
	if testPatch.Resolution != "" {
		t.Resolution = testPatch.Resolution
	}
	if testPatch.Doc != nil && len(testPatch.Doc) > 0 {
		for k, v := range testPatch.Doc {
			t.Doc[k] = v
		}
	}
}

// Equal will deeply check the comparedTest against the current instance of the test and return a bool if the test is
// equal in values. Does not check timestamps
func (t *Test) Equal(comparedTest *Test) bool {
	oarDetailsAreEqual := t.ID == comparedTest.ID &&
		t.Summary == comparedTest.Summary &&
		t.Outcome == comparedTest.Outcome &&
		t.Analysis == comparedTest.Analysis &&
		t.Resolution == comparedTest.Resolution

	if !oarDetailsAreEqual {
		return false
	}

	return cmp.Equal(t.Doc, comparedTest.Doc)
}

// TestQuery is a structure that can define a query request for existing tests. It is the same as the Test
// structure, except it takes array inputs for each field. Passing multiple values within an array will be treated
// as a logical 'OR' for querying that field. Multiple attributes passed in the query will be treated as logical
// 'AND'
//
// Additionally, Doc can be queried, it will partially match with the Postgres "contains (@>)" operator.
// For more information, see: https://www.postgresql.org/docs/current/functions-json.html
type TestQuery struct {
	IDs            []uint64       `json:"ids,omitempty"`
	Summaries      []string       `json:"summaries,omitempty"`
	Outcomes       []Outcome      `json:"outcomes,omitempty"`
	Analyses       []Analysis     `json:"analyses,omitempty"`
	Resolutions    []Resolution   `json:"resolutions,omitempty"`
	CreatedBefore  *time.Time     `json:"createdBefore,omitempty"`
	CreatedAfter   *time.Time     `json:"createdAfter,omitempty"`
	ModifiedBefore *time.Time     `json:"modifiedBefore,omitempty"`
	ModifiedAfter  *time.Time     `json:"modifiedAfter,omitempty"`
	Doc            map[string]any `json:"doc,omitempty"`
}
