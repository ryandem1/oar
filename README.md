# The OAR Framework for Software Test Reporting

## Outcome, Analysis, Resolution

## Background

There seems to be a fundamental flaw with most common test case management systems that are currently widely
used. Many of these systems come from the days when software was tested mostly manually and many of them are bloated and
do not seem to aim to solve any specific issue with software quality. They often serve as a historical ledger of test 
results and a loose bank of test "cases" that are often out-of-date or constantly changing and analysis of such "cases" can
often result in insights that are less than insightful.

The reality is that simply the act of gathering test results and looking at pass/fail 
results are not enough to assess defect presence/risk and software quality. This also does not do anything to improve
the actual tests themselves. The result is a common negative feedback loop where test results become more meaningless 
over time and tests trust falls. Trust must be preserved in tests and tests must be actively valued or not used at all.

Here is a paradigm that uses aims to create a positive feedback loop that increases software and test quality, builds 
trust in tests and software, and promotes active engagement in software quality.

#### Outcome
We start after a test is performed against a system-under-test and we get a "pass" or "failed" result. This is our
**outcome**. An outcome should only ever be binary; if there is any ambiguity in whether a test has failed, then 
there might be unclear requirements or a test that is testing for things that it doesn't need to test for. By keeping 
test outcomes as passed or failed, we are able to extract insights from them in a standard way

#### Analysis
Here comes the more actionable parts. Let's say now that you have a set of test outcomes. Information about outcomes 
themselves are not very meaningful without **analysis**. Tests can pass or fail for many reasons that tell you 
different things about your software/test quality. The OAR framework presents a streamlined way to categorize the 
result of an outcome's  analysis:

- **True Positive**: The test case failed and correctly indicated a defect on the feature-under-test
- **False Positive**: The test case failed, but under further inspection, no defect existed for the feature
- **True Negative**: The test case passed, and the feature-under-test is actually exhibiting correct behavior.
- **False Negative**: The test case passed, but the feature-under-test was found to have defects.

> **_NOTE:_**  For this definition, we say that a test produces a "positive" result when it fails, indicating positive
> presence for defect on the feature that it is testing. A "negative" result is a passed test and indicates a
> "negative" presence for defect on the feature.

> **_NOTE:_**  These basic categories can supply many metrics, but insights from those sort of assume that
> full test coverage exists, which is almost never the case. A metric that is helpful for tracking if there is 
> adequate coverage is **test effectiveness**, which can be thought as:
> 
> 
> (Total number of defects found by tests / Total number of defects found outside of tests) * 100
> 
> The goal is not always to have 100% test effectiveness, but to find an acceptable threshold for the given test suite
> that is proportional to how much effort is being put into maintaining the suite.

By using these classic definitions to categorize software test results, we are able to derive insightful metrics about 
our testing and our software.

#### Resolution
Okay, more action. The OAR framework aims to improve software/test quality, so there must be action taken depending on the
analysis results of the tests. Every analysis result (except true negative), must have a resolution. The 
OAR framework also sets out to streamline the resolutions tests must have:

For True Positives:
- **Ticket Created**: A bug ticket or feature ticket was created to track the defect's future resolution.
- **Quick Fix**: The defect was minor and a quick fix was applied that fixed the defect, there is no need to open a ticket.
- **Known Issue**: A ticket was previously open for the defect and is still pending resolution

For False Positives or False Negatives:
- **Test Fixed**: The test that threw the false positive/negative was fixed and can now work as expected. 
- **Test Disabled**: The test/part of the test that threw the false positive/negative was disabled. Possibly indicating faulty 
test design to begin with, lack of maintenance, or too narrow/broad of a check.

By sticking to these definitions, the team is sticking to actionability and testing stakeholders are actively engaged in
software quality.

## Application

While the OAR framework does not have to necessarily be tied to any one specific implementation, I set out to make an 
implementation myself. This implementation includes:

- A minimal backend written in Go that handles CRUD test/action operations
- A Postgres DB that will be used for minimal relational data, more for the impressive BJSON performance. Most test
results will be schema-less documents.
- A UI that will provide: 
  - an interface into the real-time test result ledger with a JSON filter.
  - a central place for developers/testers to provide analysis on test results
  - a simple resolution workflow.

The hope is that test results have streamlined actionability and triaging software quality issues becomes engaging to all. 

### Concepts

#### Test

In the application, there is only 1 main data structure that get enriched through the OAR process.

A **Test** represents point-in-time information about a **test** that occurred on a feature. While the real size of the 
systems that we test can vary in size, our **feature-under-test** should be a logical piece that has a single purpose.

The **Summary** can be thought of as a title, or a short description describing what the Test accomplished. A good
rule-of-thumb is that if you cannot describe a test in a Summary, then the Test is probably too broad.

The **Outcome** is the 'O' part of the OAR, it is the simple test binary and should remain that way with no ambiguity.
If there is ambiguity whether a test passed or failed, then there could be undefined requirements that the test
should probably not test for.

The **Analysis** is the 'A' and will most likely be done after the initial Test upload. It should be performed by 
someone that has some sort of ownership of the test. An accurate analysis is important for proper statistic and 
resolution. An analysis can change, but should not change after a **resolution** is added.

The **Resolution** is the 'R'. A resolution is the end of the OAR framework, the application follows the pre-defined 
resolutions that are laid out.

The **Doc** is an unstructured part of a test. It is primarily here to store test diagnostic information and test 
metadata. It can be used to filter by in the UI, and can include helpful information for the analysis/resolution 
portion of the process. This also allows the test to become further enriched with more data at the analysis/resolution 
phase. This can include information like ticket number/comments/trace links.


#### Service

The backend is quite simple, it is there to provide an interface for the CRUD test operations. It also has endpoints to 
facilitate test enrichment in the analysis/resolution phases. It connects to a Postgres database and stores tests as 
partially structured, partially unstructured data. 

> **POST /test**  
> Sending a post request to the ``/test`` endpoint will create a new test result. The fields:
> "id", "summary", "outcome", "analysis", and "resolution" are the structured part of the test and will be treated 
> differently from other fields.
> 
> You are able to also send any arbitrary JSON data in the request body, and it will be stored as the unstructured data.
> This can include any relevant test metadata and helpful diagnostic information for the analysis/resolution.
