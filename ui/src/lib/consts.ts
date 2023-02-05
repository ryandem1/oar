export enum Outcome {
    Passed = "Passed",
    Failed = "Failed"
}

export enum Analysis {
    NotAnalyzed = "Not Analyzed",
    TruePositive = "True Positive",
    FalsePositive = "False Positive",
    TrueNegative = "True Negative",
    FalseNegative = "False Negative"
}

export enum Resolution {
    Unresolved = "Unresolved",
    TicketCreated = "Ticket Created",
    QuickFix = "Quick Fix",
    KnownIssue = "Known Issue",
    TestFixed = "Test Fixed",
    TestDisabled = "Test Disabled"
}
