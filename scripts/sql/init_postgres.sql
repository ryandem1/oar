/*
Will initialize the needed table in postgres for the OAR app to function. This is the primary staging place for test
results. If desired, Snowflake is a good place to store test data for archiving/analysis; in this case, data should be
periodically transferred over to Snowflake.
*/
create or replace function update_modified_column()
returns trigger as $$
begin
    new.modified = (now() at time zone 'utc');
    return new;
end;
$$ language 'plpgsql';
comment on function update_modified_column() is 'Updates the "modified" column of "oar_tests" with the latest UTC timestamp every time a record gets changed';

create table if not exists oar_tests
(
    id          bigserial   constraint id primary key,
    summary     text        not null,
    outcome     varchar(6)  not null,
    analysis    varchar(13) not null,
    resolution  varchar(20) not null,
    created     timestamp not null default (now() at time zone 'utc'),
    modified    timestamp not null default (now() at time zone 'utc'),
    doc         jsonb,
    constraint analysis
        check (analysis in ('NotAnalyzed', 'TruePositive', 'FalsePositive', 'TrueNegative', 'FalseNegative')),
    constraint outcome
        check (outcome in ('Passed', 'Failed')),
    constraint resolution
        check (resolution in ('Unresolved', 'NotNeeded', 'TicketCreated', 'QuickFix', 'KnownIssue', 'TestFixed', 'TestDisabled'))
);

-- Will add the trigger that updates the modified column automatically on every update.
create trigger update_modified
before update on oar_tests
for each row execute procedure update_modified_column();

comment on table oar_tests
    is 'tests is the core test ledger where results will be stored. Contains both structured test data and unstructured data that will be stored in BJSON';

comment on constraint ID on oar_tests
    is 'Unique identifier of a test result';

comment on column oar_tests.summary
    is 'Short description of the test that was performed';

comment on column oar_tests.outcome
    is 'Either "Passed" or "Failed". Binary outcome of the test';

comment on column oar_tests.analysis
    is 'The analysis conclusion of the test outcome';

comment on column oar_tests.resolution
    is 'The resolution of an actionable test analysis';

comment on column oar_tests.created
    is 'UTC timestamp of when the test result was reported. Will automatically be set and should not be changed';

comment on column oar_tests.modified
    is 'UTC timestamp of when the test result was last enriched/modified. Will automatically be set on every update.';

comment on column oar_tests.doc
    is 'Unstructured document for any additional test result data';

comment on constraint analysis on oar_tests
    is 'Ensures that the analysis is a valid analysis option';

comment on constraint outcome on oar_tests
    is 'Ensures that an OUTCOME is either "Passed" or "Failed"';

comment on constraint resolution on oar_tests
    is 'Ensures that a resolution is a valid value';
