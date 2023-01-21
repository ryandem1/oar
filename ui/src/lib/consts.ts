export enum Outcome {
    PASSED = "PASSED",
    FAILED = "FAILED",
}

export enum Analysis {
    NotAnalyzed = "NotAnalyzed",
    TruePositive = "TruePositive",
    FalsePositive = "FalsePositive",
    TrueNegative = "TrueNegative",
    FalseNegative = "FalseNegative"
}

export enum Resolution {
    Unresolved = "Unresolved",
    TicketCreated = "TicketCreated",
    QuickFix = "QuickFix",
    KnownIssue = "KnownIssue",
    TestFixed = "TestFixed",
    TestDisabled = "TestDisabled"
}
