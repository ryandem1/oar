import base64
import datetime
import json
from typing import Any

from pydantic import BaseModel, Field

from oar.consts import Outcome, Analysis, Resolution


class Test(BaseModel):
    """
    Primary structure for OAR test results. This structure is meant to have attributes dynamically added to it
    """
    id_: int = Field(0, alias="id")
    summary: str | None = None
    outcome: "Outcome" | None = None
    analysis: "Analysis" | None = None
    resolution: "Resolution" | None = None

    model_config = {
        "extra": "allow",
        "use_enum_values": True
    }

    def as_request_body(self) -> dict[str, Any]:
        """
        Formats the Test in a format appropriate for the OAR client.

        Returns
        -------
        request_body: dict[str, Any]
            Test as a request body (unmerges the doc attribute)
        """
        return self.model_dump(by_alias=True)


class TestQuery(BaseModel):
    """
    Structure to query for multiple test results and send them to the `/query` endpoint
    """
    ids: list[int] | None = None
    summaries: list[str] | None = None
    outcomes: list["Outcome"] | None = None
    analyses: list["Analysis"] | None = None
    resolutions: list["Resolution"] | None = None
    created_before: datetime.datetime | None = Field(None, alias="createdBefore")
    created_after: datetime.datetime | None = Field(None, alias="createdAfter")
    modified_before: datetime.datetime | None = Field(None, alias="modifiedBefore")
    modified_after: datetime.datetime | None = Field(None, alias="modifiedAfter")
    docs: list[dict[str, Any]] | None = None

    model_config = {
        "populate_by_name": True,
        "use_enum_values": True,
    }

    def __eq__(self, other: 'TestQuery') -> bool:
        """
        Checks for equality without datetime attributes because they are flakey
        """
        return all(
            getattr(self, attr) == getattr(other, attr)
            for attr in ["ids", "summaries", "outcomes", "analyses", "resolutions", "docs"]
        )

    @classmethod
    def from_query_string(cls, query_string: str) -> 'TestQuery':
        """
        Decode/deserializes a base64 encoded query string
        """
        query_object_string = base64.b64decode(query_string.encode("ascii")).decode("ascii")
        query = cls(**json.loads(query_object_string))
        return query

    def as_request_body(self) -> dict[str, Any]:
        """
        Formats the TestQuery in a format appropriate for the OAR client.
        """
        body = self.model_dump(
            by_alias=True,
            exclude={"created_before", "created_after", "modified_before", "modified_after"}
        )
        body.update({
            "createdBefore": self.created_before.strftime("%Y-%m-%dT%H:%M:%SZ") if self.created_before else None,
            "createdAfter": self.created_after.strftime("%Y-%m-%dT%H:%M:%SZ") if self.created_after else None,
            "modifiedBefore": self.modified_before.strftime("%Y-%m-%dT%H:%M:%SZ") if self.modified_before else None,
            "modifiedAfter": self.modified_after.strftime("%Y-%m-%dT%H:%M:%SZ") if self.modified_after else None
        })
        return body

    def as_query_string(self) -> str:
        """
        base64 encodes the query to be used on the query endpoints
        """
        query_string = base64.b64encode(
            self.model_dump(mode="json", by_alias=True, exclude_none=True).encode("ascii")
        ).decode("ascii")
        return query_string


class TestQueryResult(BaseModel):
    """
    Represents a query result. Like the response from the GET ``/tests`` interface.
    """
    count: int
    tests: list[Test]
