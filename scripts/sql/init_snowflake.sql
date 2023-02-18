/*
Initializes a Snowflake table, this is designed to be a space for test data archive/analysis. It should not be used as
the first place for test data, as Snowflake is crazy expensive.
*/

create table if not exists OAR_TESTS
(
    id         bigint
        constraint id
            primary key,
    summary    text        not null,
    outcome    varchar(6)  not null,
    analysis   varchar(13) not null,
    resolution varchar(20) not null,
    doc        variant
);

comment on table OAR_TESTS is 'tests is the core test ledger where results will be stored. Contains both structured test data and unstructured data that will be stored in BJSON';

comment on column OAR_TESTS.id is 'Unique identifier of a test result';

comment on column OAR_TESTS.summary is 'Short description of the test that was performed';

comment on column OAR_TESTS.outcome is 'Either "Passed" or "Failed". Binary outcome of the test';

comment on column OAR_TESTS.analysis is 'The analysis conclusion of the test outcome';

comment on column OAR_TESTS.resolution is 'The resolution of an actionable test analysis';

comment on column OAR_TESTS.doc is 'Unstructured document for any additional test result data';
