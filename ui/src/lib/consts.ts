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
    TicketCreated = "TicketCreated",
    QuickFix = "QuickFix",
    KnownIssue = "KnownIssue",
    TestFixed = "TestFixed",
    TestDisabled = "TestDisabled"
}
