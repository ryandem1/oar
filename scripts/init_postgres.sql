/*
Will initialize the needed table in postgres for the OAR app to function. This is the primary staging place for test
results. If desired, Snowflake is a good place to store test data for archiving/analysis; in this case, data should be
periodically transferred over to Snowflake.
*/

create table if not exists OAR_TESTS
(
    id         bigserial
    constraint id
        primary key,
    summary    text        not null,
    outcome    varchar(6)  not null,
    analysis   varchar(13) not null,
    resolution varchar(20) not null,
    doc        jsonb,
    constraint analysis
        check (analysis in ('NotAnalyzed', 'TruePositive', 'FalsePositive', 'TrueNegative', 'FalseNegative')),
    constraint outcome
        check (outcome in ('Passed', 'Failed')),
    constraint resolution
        check (resolution in ('Unresolved', 'NotNeeded', 'TicketCreated', 'QuickFix', 'KnownIssue', 'TestFixed', 'TestDisabled'))
);

comment on table OAR_TESTS is 'tests is the core test ledger where results will be stored. Contains both structured test data and unstructured data that will be stored in BJSON';

comment on constraint ID on OAR_TESTS is 'Unique identifier of a test result';

comment on column OAR_TESTS.summary is 'Short description of the test that was performed';

comment on column OAR_TESTS.outcome is 'Either "Passed" or "Failed". Binary outcome of the test';

comment on column OAR_TESTS.analysis is 'The analysis conclusion of the test outcome';

comment on column OAR_TESTS.resolution is 'The resolution of an actionable test analysis';

comment on column OAR_TESTS.doc is 'Unstructured document for any additional test result data';

comment on constraint analysis on OAR_TESTS is 'Ensures that the analysis is a valid analysis option';

comment on constraint outcome on OAR_TESTS is 'Ensures that an OUTCOME is either "Passed" or "Failed"';

comment on constraint resolution on OAR_TESTS is 'Ensures that a resolution is a valid value';
