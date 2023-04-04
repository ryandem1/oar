import logging
from datetime import datetime
from typing import TypeVar

from pydantic import BaseModel

from oar.consts import Outcome, Analysis, Resolution
from oar.test import Test

logger = logging.getLogger("oar")


AnyTest = TypeVar("AnyTest", bound=Test)


class Results(BaseModel):
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

    @property
    def passed_ids(self) -> list[int]:
        return [test.id_ for test in self.tests if test.outcome == Outcome.Passed]

    @property
    def failed_ids(self) -> list[int]:
        return [test.id_ for test in self.tests if test.outcome == Outcome.Failed]

    @property
    def all_ids(self) -> list[int]:
        return [test.id_ for test in self.tests]

    @property
    def need_analysis_ids(self) -> list[int]:
        return [test.id_ for test in self.tests if test.analysis == Analysis.NotAnalyzed]

    @property
    def need_resolution_ids(self) -> list[int]:
        return [test.id_ for test in self.tests if test.resolution == Resolution.Unresolved]

    def log_summary_statistics(self) -> None:
        """
        Will log out the summary statistic attributes through the ``oar messanger`` logger

        Returns
        -------
        None
        """
        logger.info("\n============OAR SUMMARY===============")
        logger.info(f"Passed IDs: {self.passed_ids}")
        logger.info(f"Failed IDs: {self.failed_ids}")
        logger.info(f"Tests that need analysis: {self.need_analysis_ids}")
        logger.info(
            f"Tests that need resolution: {self.need_resolution_ids}" +
            "\n======================================"
        )
