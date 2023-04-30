import base64
import datetime
import json
from typing import Any

from pydantic import BaseModel, Field, Extra

from oar.consts import Outcome, Analysis, Resolution


class Test(BaseModel):
    """
    Primary structure for OAR test results. This structure is meant to have attributes dynamically added to it
    """
    id_: int = Field(0, alias="id")  # Sometimes ID will be ignored
    summary: str | None = None
    outcome: Outcome | None = None
    analysis: Analysis | None = None
    resolution: Resolution | None = None

    class Config:
        extra = Extra.allow
        use_enum_values = True

    def as_request_body(self) -> dict[str, Any]:
        """
        Formats the Test in a format appropriate for the OAR client.

        Returns
        -------
        request_body: dict[str, Any]
            Test as a request body (unmerges the doc attribute)
        """
        return self.dict(by_alias=True)


class TestQuery(BaseModel):
    """
    Structure to query for multiple test results and send them to the ``/query`` endpoint
    """
    ids: list[int] | None = None
    summaries: list[str] | None = None
    outcomes: list[Outcome] | None = None
    analyses: list[Analysis] | None = None
    resolutions: list[Resolution] | None = None
    created_before: datetime.datetime | None = Field(None, alias="createdBefore")
    created_after: datetime.datetime | None = Field(None, alias="createdAfter")
    modified_before: datetime.datetime | None = Field(None, alias="modifiedBefore")
    modified_after: datetime.datetime | None = Field(None, alias="modifiedAfter")
    docs: list[dict[str, Any]] | None = None

    @classmethod
    def from_query_string(cls, query_string: str) -> 'TestQuery':
        """
        Will return a decoded/deserialized ``TestQuery`` object from a base64 encoded test query string. Like the string
        you would get if you called the ``/query`` endpoint.

        Parameters
        ----------
        query_string : str
            base64 encoded query string

        Returns
        -------
        query : TestQuery
            Query object created from the base64 encoded string
        """
        query_object_string = base64.b64decode(query_string.encode("ascii")).decode("ascii")
        query = cls(**json.loads(query_object_string))
        return query

    def as_request_body(self) -> dict[str, Any]:
        """
        Formats the TestQuery in a format appropriate for the OAR client.

        Returns
        -------
        request_body: dict[str, Any]
            TestQuery as a request body
        """
        return self.dict(by_alias=True)

    def as_query_string(self) -> str:
        """
        base64 encodes the query to be used on the query endpoints. This is equivalent to calling the ``/query`` method.
        So, if using in Python, better to use this method to save the call.

        Returns
        -------
        query_string : str
            Encoded Query string that is compatible for the various ``/tests`` endpoint
        """
        query_string = base64.b64encode(s=self.json(by_alias=True).encode("ascii")).decode("ascii")
        return query_string


class TestQueryResult(BaseModel):
    """
    Represents a query result. Like the response from the GET ``/tests`` interface.
    """
    count: int
    tests: list[Test]
