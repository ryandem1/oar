# OAR Test Manager

## Outcome, Analysis, Resolution

## Background

There seems to be a fundamental flaw with most common test case management systems that are currently widely
used. Many of these systems come from the days where software was tested mostly manually and many of them are bloated and
do not seem to aim to solve any specific issue with software quality. They often serve as a historical ledger of test 
results and a loose bank of test "cases" that are often out-of-date or constantly changing and analysis of such "cases" can
often result in insights that are less than insightful. Part of the change in ideology is that tests are point-in-time, 
there is no need to store test definitions in a separate places.

The reality is that simply the act of gathering test results and looking at pass/fail 
results are not enough to determine defect presence/risk and software quality. This also does not do anything to improve
the actual tests themselves. The result is a common negative feedback loop where test results become more meaningless 
over time and tests trust falls. Trust must be preserved in tests and tests must be actively valued or not used at all.

Here is a paradigm that uses aims to create a positive feedback loop that increases software and test quality, builds 
trust in tests and software, and promotes active engagement in software quality.

#### Outcome
We start after a test is performed against a system-under-test and we get a "pass" or "failed" result.

### Concepts

#### Test
A **Test** represents point-in-time information about a **test** that occurred on a system. While the real size of the 
systems that we test can vary in size, our **system-under-test** should be a logical piece that has a single purpose.

The **Summary** can be thought of as a title, or a short description describing what the Test accomplished. A good
rule-of-thumb is that if you cannot describe a test in a Summary, then the Test is probably too broad.

The **Outcome** is the 'O' part of the OAR, it is the simple test binary and should remain that way with no ambiguity.
If there is ambiguity whether a test passed or failed, then there could be 
