enum Outcome {
  Passed,
  Failed
}

enum Analysis {
  NotAnalyzed,
  TruePositive,
  FalsePositive,
  TrueNegative,
  FalseNegative
}

enum Resolution {
  Unresolved,
  NotNeeded,
  TicketCreated,
  QuickFix,
  KnownIssue,
  TestFixed,
  TestDisabled
}
