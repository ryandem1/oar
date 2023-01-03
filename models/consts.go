package models

type Outcome string

const (
	Passed  Outcome = "Passed"
	Failed  Outcome = "Failed"
	Skipped Outcome = "Skipped"
)

type ActionTaken string

const (
	BugCreated    ActionTaken = "BugCreated"
	FalsePositive ActionTaken = "FalsePositive"
	QuickFix      ActionTaken = "QuickFix"
)
