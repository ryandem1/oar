import json
import logging
import os
from datetime import datetime
from enum import Enum
from pathlib import Path
from typing import Any, TypeVar

import tomli
from pydantic import BaseModel, BaseSettings, Field, Extra

logger = logging.getLogger("oar")


class EnvConfig(BaseSettings):
    """
    Primary environment configuration for the OAR PyTest plugin. Both is the structure for the environment and provides
    methods to easily access the environment.
    """
    host: str = "oar-service:8080"  # Base URL of the OAR instance to send results to
    send_results: bool = False  # This is what will control sending the results to the OAR instance
    store_results: bool = True  # This will enable the `oar_results` fixture, will not prevent sending results to OAR
    log_summary: bool = True  # This will control the logging of summary statistics in a run
    output_file: bool = True  # This will output a JSON results file with name `oar-results-<utc-timestamp>.json`
    output_dir: str = "oar-results"  # Controls where JSON results files will be stored. Relative to CWD

    class Config:
        env_prefix = "OAR_"

    @classmethod
    def from_file(cls, config_file_path: Path) -> 'EnvConfig':
        """
        Will return a new ``EnvConfig`` file by reading from an environment configuration file by path

        Parameters
        ----------
        config_file_path : Path
            Full path of the config file to initialize an ``EnvConfig`` from.

        Returns
        -------
        config : EnvConfig
            Representation of environment configuration
        """
        if not config_file_path.exists():
            return cls()  # If config file is not found, do not error, fallback on default

        with config_file_path.open("rb") as config_file:
            match config_file_path.suffix:
                case ".toml":
                    return cls(**tomli.load(config_file))
                case ".json":
                    return cls(**json.load(config_file))
                case _:
                    raise FileNotFoundError("Config file must be a .json or a .toml file!")

    @classmethod
    def get(cls) -> 'EnvConfig':
        """
        Will return the default ``EnvConfig``. This is defined as the config .json or .toml file that is located at the
        ``OAR_CONFIG_PATH`` environment variable (whose default location is the root of the project/oar-config.toml).

        Returns
        -------
        config : EnvConfig
            Default ``EnvConfig`` object
        """
        config_file_path = Path(os.getcwd()) / os.environ.get("OAR_CONFIG_PATH", "oar-config.toml")
        return cls.from_file(config_file_path)


class Outcome(str, Enum):
    Passed = "Passed"
    Failed = "Failed"


class Analysis(str, Enum):
    NotAnalyzed = "NotAnalyzed"
    TruePositive = "TruePositive"
    FalsePositive = "FalsePositive"
    TrueNegative = "TrueNegative"
    FalseNegative = "FalseNegative"


class Resolution(str, Enum):
    Unresolved = "Unresolved"
    NotNeeded = "NotNeeded"
    TicketCreated = "TicketCreated"
    QuickFix = "QuickFix"
    KnownIssue = "KnownIssue"
    TestFixed = "TestFixed"
    TestDisabled = "TestDisabled"


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

    def as_request_body(self) -> dict[str, Any]:
        """
        Formats the Test in a format appropriate for the OAR client.

        Returns
        -------
        request_body: dict[str, Any]
            Test as a request body (unmerges the doc attribute)
        """
        return self.dict(by_alias=True)


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
