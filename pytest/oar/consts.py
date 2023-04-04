from enum import Enum


class Outcome(str, Enum):
    Passed = "Passed"
    Failed = "Failed"


class Analysis(str, Enum):
    NotAnalyzed = "NotAnalyzed"
    TruePositive = "TruePositive"
    FalsePositive = "FalsePositive"
    TrueNegative = "TrueNegative"
    FalseNegative = "FalseNegative"


class Resolution(str, Enum):
    Unresolved = "Unresolved"
    NotNeeded = "NotNeeded"
    TicketCreated = "TicketCreated"
    QuickFix = "QuickFix"
    KnownIssue = "KnownIssue"
    TestFixed = "TestFixed"
    TestDisabled = "TestDisabled"
