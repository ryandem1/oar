package models

import (
	"fmt"
	"golang.org/x/exp/slices"
)

// TestResolution represents a single Resolution for a test. A test should only have a single Resolution at a time.
type TestResolution struct {
	TestID     int        `json:"testId"`
	Resolution Resolution `json:"resolution"`
	Comment    string     `json:"comment"`
}

// Validate will ensure that a TestResolution has a proper Resolution
func (tr *TestResolution) Validate() error {
	validResolutions := []Resolution{TicketCreated, QuickFix, KnownIssue, TestFixed, TestDisabled}
	if !slices.Contains(validResolutions, tr.Resolution) {
		return fmt.Errorf(
			"invalid resolution: '%s', must be one of resultions: %s",
			tr.Resolution,
			validResolutions,
		)
	}

	return nil
}
