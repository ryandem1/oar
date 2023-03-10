package main

type Outcome string

const (
	Passed Outcome = "Passed"
	Failed Outcome = "Failed"
)

type Analysis string

const (
	NotAnalyzed   Analysis = "NotAnalyzed"
	TruePositive  Analysis = "TruePositive"
	FalsePositive Analysis = "FalsePositive"
	TrueNegative  Analysis = "TrueNegative"
	FalseNegative Analysis = "FalseNegative"
)

type Resolution string

const (
	Unresolved    Resolution = "Unresolved"
	NotNeeded     Resolution = "NotNeeded"
	TicketCreated Resolution = "TicketCreated"
	QuickFix      Resolution = "QuickFix"
	KnownIssue    Resolution = "KnownIssue"
	TestFixed     Resolution = "TestFixed"
	TestDisabled  Resolution = "TestDisabled"
)
