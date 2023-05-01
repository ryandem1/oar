import datetime
import random
import typing

import pytest

import oar


@pytest.fixture
def valid_test() -> dict[str, typing.Any]:
    """
    Returns
    -------
    test : dict[str, typing.Any]
        random valid OAR test
    """
    test = random.choice([
        {
            "id": 20,
            "summary": "Checks that a valid input produces a valid output",
            "outcome": "Failed",
            "analysis": "FalsePositive",
            "resolution": "TestFixed",
            "created": "2023-05-01T13:53:04.725023Z",
            "modified": "2023-05-01T13:53:04.735876Z",
            "doc": {
                "latency (ms)": {
                    "p50": 254.33,
                    "p75": 332.45,
                    "p95": 501.99,
                    "p99": 676.51
                },
                "maxRPS": 300,
                "owner": "Squidward Tentacles",
                "runtime": "10m",
                "samplePayloads": [
                    {
                        "app_id": "47324033",
                        "status": "APPROVED"
                    },
                    {
                        "app_id": "9948302",
                        "status": "REJECTED"
                    }
                ],
                "service": "application-service",
                "test left merge field": "different value, different type",
                "type": "load"
            }
        },
        {
            "id": 14,
            "summary": "Navbar component link positive test",
            "outcome": "Passed",
            "analysis": "TrueNegative",
            "resolution": "Unresolved",
            "created": "2023-05-01T13:53:04.661289Z",
            "modified": "2023-05-01T13:53:04.661289Z",
            "doc": {
                "latency (ms)": {
                    "p50": 254.33,
                    "p75": 332.45,
                    "p95": 501.99,
                    "p99": 676.51
                },
                "maxRPS": 300,
                "owner": "Squidward Tentacles",
                "runtime": "10m",
                "samplePayloads": [
                    {
                        "app_id": "47324033",
                        "status": "APPROVED"
                    },
                    {
                        "app_id": "9948302",
                        "status": "REJECTED"
                    }
                ],
                "service": "application-service",
                "type": "load"
            }
        },
        {
            "id": 13,
            "summary": "Ensures the /metadata endpoint is functional",
            "outcome": "Passed",
            "analysis": "FalseNegative",
            "resolution": "Unresolved",
            "created": "2023-05-01T13:53:04.660056Z",
            "modified": "2023-05-01T13:53:04.660056Z",
            "doc": {
                "browsers": [
                    "chrome",
                    "firefox",
                    "edge"
                ],
                "owner": "Sandy Cheeks",
                "screenshotURL": "https://some-s3-bucket-that-doesnt-exist.com/714029473432412",
                "type": "UI"
            }
        },
        {
            "id": 12,
            "summary": "Test user insert query is functional",
            "outcome": "Passed",
            "analysis": "NotAnalyzed",
            "resolution": "NotNeeded",
            "created": "2023-05-01T13:53:04.65888Z",
            "modified": "2023-05-01T13:53:04.65888Z",
            "doc": {
                "browsers": [
                    "chrome",
                    "firefox",
                    "edge"
                ],
                "owner": "Sandy Cheeks",
                "screenshotURL": "https://some-s3-bucket-that-doesnt-exist.com/714029473432412",
                "type": "UI"
            }
        },
        {
            "id": 11,
            "summary": "Ensures that publishing a valid Kafka event gets consumed correctly downstream",
            "outcome": "Passed",
            "analysis": "TrueNegative",
            "resolution": "TestFixed",
            "created": "2023-05-01T13:53:04.657613Z",
            "modified": "2023-05-01T13:53:04.657613Z",
            "doc": {
                "app": "user-service",
                "owner": "Patrick Star",
                "testPayload": {
                    "accountStatus": "lock",
                    "id": 1,
                    "username": "someUser48"
                },
                "testResponse": {
                    "responseBody": None,
                    "responseCode": 200
                },
                "type": "integration"
            }
        },
        {
            "id": 10,
            "summary": "Ensures a bad input returns a correct error message",
            "outcome": "Passed",
            "analysis": "TrueNegative",
            "resolution": "NotNeeded",
            "created": "2023-05-01T13:53:04.654884Z",
            "modified": "2023-05-01T13:53:04.654884Z",
            "doc": {
                "latency (ms)": {
                    "p50": 254.33,
                    "p75": 332.45,
                    "p95": 501.99,
                    "p99": 676.51
                },
                "maxRPS": 300,
                "owner": "Squidward Tentacles",
                "runtime": "10m",
                "samplePayloads": [
                    {
                        "app_id": "47324033",
                        "status": "APPROVED"
                    },
                    {
                        "app_id": "9948302",
                        "status": "REJECTED"
                    }
                ],
                "service": "application-service",
                "type": "load"
            }
        },
        {
            "id": 9,
            "summary": "Ensures the /metadata endpoint is functional",
            "outcome": "Failed",
            "analysis": "TruePositive",
            "resolution": "TestFixed",
            "created": "2023-05-01T13:53:04.560903Z",
            "modified": "2023-05-01T13:53:04.560903Z",
            "doc": {
                "browsers": [
                    "chrome",
                    "firefox",
                    "edge"
                ],
                "owner": "Sandy Cheeks",
                "screenshotURL": "https://some-s3-bucket-that-doesnt-exist.com/714029473432412",
                "type": "UI"
            }
        },
        {
            "id": 8,
            "summary": "User service load test",
            "outcome": "Failed",
            "analysis": "FalsePositive",
            "resolution": "KnownIssue",
            "created": "2023-05-01T13:53:04.535689Z",
            "modified": "2023-05-01T13:53:04.535689Z",
            "doc": {
                "latency (ms)": {
                    "p50": 254.33,
                    "p75": 332.45,
                    "p95": 501.99,
                    "p99": 676.51
                },
                "maxRPS": 300,
                "owner": "Squidward Tentacles",
                "runtime": "10m",
                "samplePayloads": [
                    {
                        "app_id": "47324033",
                        "status": "APPROVED"
                    },
                    {
                        "app_id": "9948302",
                        "status": "REJECTED"
                    }
                ],
                "service": "application-service",
                "type": "load"
            }
        },
        {
            "id": 7,
            "summary": "Test user insert query is functional",
            "outcome": "Failed",
            "analysis": "TruePositive",
            "resolution": "Unresolved",
            "created": "2023-05-01T13:53:04.512018Z",
            "modified": "2023-05-01T13:53:04.512018Z",
            "doc": {
                "app": "user-service",
                "owner": "Patrick Star",
                "testPayload": {
                    "accountStatus": "lock",
                    "id": 1,
                    "username": "someUser48"
                },
                "testResponse": {
                    "responseBody": None,
                    "responseCode": 200
                },
                "type": "integration"
            }
        },
        {
            "id": 6,
            "summary": "Navbar component link positive test",
            "outcome": "Failed",
            "analysis": "TruePositive",
            "resolution": "NotNeeded",
            "created": "2023-05-01T13:53:04.48832Z",
            "modified": "2023-05-01T13:53:04.48832Z",
            "doc": {
                "latency (ms)": {
                    "p50": 254.33,
                    "p75": 332.45,
                    "p95": 501.99,
                    "p99": 676.51
                },
                "maxRPS": 300,
                "owner": "Squidward Tentacles",
                "runtime": "10m",
                "samplePayloads": [
                    {
                        "app_id": "47324033",
                        "status": "APPROVED"
                    },
                    {
                        "app_id": "9948302",
                        "status": "REJECTED"
                    }
                ],
                "service": "application-service",
                "type": "load"
            }
        },
        {
            "id": 5,
            "summary": "Ensures that publishing a valid Kafka event gets consumed correctly downstream",
            "outcome": "Failed",
            "analysis": "FalsePositive",
            "resolution": "TicketCreated",
            "created": "2023-05-01T13:53:04.463057Z",
            "modified": "2023-05-01T13:53:04.463057Z",
            "doc": {
                "browsers": [
                    "chrome",
                    "firefox",
                    "edge"
                ],
                "owner": "Sandy Cheeks",
                "screenshotURL": "https://some-s3-bucket-that-doesnt-exist.com/714029473432412",
                "type": "UI"
            }
        },
        {
            "id": 4,
            "summary": "User service load test",
            "outcome": "Passed",
            "analysis": "TrueNegative",
            "resolution": "NotNeeded",
            "created": "2023-05-01T13:53:04.323228Z",
            "modified": "2023-05-01T13:53:04.354813Z",
            "doc": {
                "app": "user-service",
                "owner": "Patrick Star",
                "testPayload": {
                    "accountStatus": "lock",
                    "id": 1,
                    "username": "someUser48"
                },
                "testResponse": {
                    "responseBody": None,
                    "responseCode": 200
                },
                "type": "integration"
            }
        },
        {
            "id": 1,
            "summary": "User service load test",
            "outcome": "Failed",
            "analysis": "TruePositive",
            "resolution": "TicketCreated",
            "created": "2023-05-01T13:53:04.197727Z",
            "modified": "2023-05-01T13:53:04.197727Z",
            "doc": {
                "browsers": [
                    "chrome",
                    "firefox",
                    "edge"
                ],
                "owner": "Sandy Cheeks",
                "screenshotURL": "https://some-s3-bucket-that-doesnt-exist.com/714029473432412",
                "type": "UI"
            }
        }
    ])
    return test


@pytest.fixture
def valid_query() -> dict[str, typing.Any]:
    """
    Returns
    -------
    query : dict[str, typing.Any]
        Valid test query
    """
    query = {
        "ids": [random.randint(0, 50) for _ in range(random.randint(0, 25))],
        "summaries": random.choices(["error message", "load test", "Navbar", "/metadata"], k=random.randint(0, 3)),
        "outcomes": random.choices(list(oar.Outcome), k=random.randint(1, len(list(oar.Outcome)))),
        "analyses": random.choices(list(oar.Analysis), k=random.randint(1, len(list(oar.Analysis)))),
        "resolutions": random.choices(list(oar.Resolution), k=random.randint(1, len(list(oar.Resolution)))),
        "createdBefore": datetime.datetime.utcnow() + datetime.timedelta(days=random.randint(0, 2)),
        "createdAfter": datetime.datetime.utcnow() - datetime.timedelta(days=random.randint(0, 2)),
        "modifiedBefore": datetime.datetime.utcnow() + datetime.timedelta(days=random.randint(0, 2)),
        "modifiedAfter": datetime.datetime.utcnow() - datetime.timedelta(days=random.randint(0, 2)),
        "docs": random.choices([
            {"type": "UI"},
            {"owner": "Sandy Cheeks"},
            {"browsers": ["chrome"]},
            {"app": "users-service"}
        ], k=random.randint(1, 2))
    }

    # Makes some fields None because all query fields are optional
    for k in query:
        if random.getrandbits(1):
            query[k] = None

    return query
