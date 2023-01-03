package models

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type Test struct {
	Name    string
	Outcome Outcome
	Action  ActionTaken
	Doc     map[string]any
}

func (t *Test) Validate() error {
	validOutcomes := []Outcome{Passed, Failed, Skipped}
	validActions := []ActionTaken{BugCreated, FalsePositive, QuickFix}

	if !slices.Contains(validOutcomes, t.Outcome) {
		return fmt.Errorf("invalid outcome: '%s', must be one of outcomes: %s", t.Outcome, validOutcomes)
	}

	if !slices.Contains(validActions, t.Action) {
		return fmt.Errorf("invalid action: '%s', must be one of actions: %s", t.Action, validActions)
	}
	return nil
}
