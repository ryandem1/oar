import json
import logging
import os
from pathlib import Path

import tomli
from pydantic import BaseSettings

logger = logging.getLogger("oar")


class EnvConfig(BaseSettings):
    """
    Primary environment configuration for the OAR PyTest plugin. Both is the structure for the environment and provides
    methods to easily access the environment.
    """
    host: str = "oar-service:8080"  # Base URL of the OAR instance to send results to
    send_results: bool = False  # This is what will control sending the results to the OAR instance
    store_results: bool = True  # This will enable the `oar_results` fixture, will not prevent sending results to OAR
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
