package models

type Outcome string

const (
	Passed Outcome = "Passed"
	Failed Outcome = "Failed"
)

type Analysis string

const (
	TruePositive  Analysis = "TruePositive"
	FalsePositive Analysis = "FalsePositive"
	TrueNegative  Analysis = "TrueNegative"
	FalseNegative Analysis = "FalseNegative"
)

type Resolution string

const (
	TicketCreated Resolution = "BugCreated"
	QuickFix      Resolution = "QuickFix"
	TestFixed     Resolution = "TestFixed"
)
