from datetime import datetime
from typing import TypeVar

from pydantic import BaseModel

from oar.models import Test

AnyTest = TypeVar("AnyTest", bound=Test)


class Report(BaseModel):
    """
    Aggregate OAR result information for a run.
    """
    start_time: str = str(datetime.utcnow())
    completed_time: str | None = None
    tests: list[AnyTest] = []

    def __iadd__(self, other: AnyTest) -> None:
        """
        You can use code like: ``oar.Results += oar.Test`` to add a new test to the Results.tests

        Parameters
        ----------
        other : AnyTest
            Any OAR test instance to add to the results

        Returns
        -------
        None
        """
        self.tests.append(other)
