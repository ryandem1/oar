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
