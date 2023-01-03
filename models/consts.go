package models

type Outcome int32

const (
	Passed Outcome = iota
	Failed
	Skipped
)

type ActionTaken int32

const (
	BugCreated ActionTaken = iota
	FalsePositive
	QuickFix
)
