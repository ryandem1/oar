import os
import json
import tomli

from pydantic import BaseModel, BaseSettings
from enum import Enum
from typing import Any
from pathlib import Path


class EnvConfig(BaseSettings):
    """
    Primary environment configuration for the OAR PyTest plugin. Both is the structure for the environment and provides
    methods to easily access the environment.
    """
    host = "oar-service:8080"

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
    id_: int
    summary: str
    outcome: Outcome
    analysis: Analysis
    resolution: Resolution
    doc: dict[str, Any] | None
