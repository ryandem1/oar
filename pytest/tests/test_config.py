import json
import os
import pathlib
from typing import Literal

import pytest

import oar


@pytest.fixture
def config_file(request: pytest.FixtureRequest) -> tuple[dict, pathlib.Path]:
    """
    Creates an OAR config JSON file and returns a path to the created file. Will delete file
    after test.

    Parameters
    ----------
    request : pytest.FixtureRequest
        PyTest fixture request to get file extension parameter

    Yields
    -------
    path : tuple[dict, pathlib.Path]
        Config dictionary and path to file
    """
    extension: Literal["toml", "json"] = getattr(request, "param")
    config_file_content = {
        "host": "http://localhost:8080",
        "send_results": True,
        "store_results": False,
        "output_file": False,
        "output_dir": "test_dir"
    }

    path = pathlib.Path(__file__).parent.parent / f"oar-config.{extension}"

    match extension:

        case "json":
            with path.open("w") as config_file:
                json.dump(config_file_content, config_file)

        case "toml":
            contents = f"""
            host = "{config_file_content['host']}"
            send_results = "{config_file_content['send_results']}"
            store_results = "{config_file_content['store_results']}"
            output_file = "{config_file_content['output_file']}"
            output_dir = "{config_file_content['output_dir']}"
            """
            path.write_text(contents)

        case _:
            raise NotImplementedError

    yield config_file_content, path

    os.remove(path)


@pytest.fixture
def invalid_config_file(request: pytest.FixtureRequest) -> pathlib.Path:
    """
    Parameters
    ----------
    request : pytest.FixtureRequest
        PyTest fixture request to get file extension parameter

    Yields
    -------
    path : pathlib.Path
        Path to a config file of an unsupported format.
    """
    extension = getattr(request, "param")
    path = pathlib.Path(__file__).parent.parent / f"invalid-config.{extension}"
    path.write_text("")

    yield path

    os.remove(path)


class TestConfig:

    @pytest.mark.parametrize("config_file", ["toml", "json"], indirect=True)
    def test_from_file(self, config_file: tuple[dict, pathlib.Path]):
        """
        Tests config creation from file.

        Parameters
        ----------
        config_file : tuple[dict, pathlib.Path]
            See fixture
        """
        contents, path = config_file

        config = oar.EnvConfig.from_file(path)
        assert config.dict() == contents

    @pytest.mark.parametrize("invalid_config_file", ["ini", ".env", ".cfg"], indirect=True)
    def test_from_file_non_supported_format(self, invalid_config_file: pathlib.Path):
        """
        Ensures config files that are not of the supported types are not accepted.
        """
        with pytest.raises(FileNotFoundError):
            oar.EnvConfig.from_file(invalid_config_file)

    def test_no_config_file(self):
        """
        Ensures no config file makes `.from_file` fall back on default
        """
        non_existent_file_path = pathlib.Path(__file__).parent.parent / "my-oar-config.toml"
        oar.EnvConfig.from_file(non_existent_file_path)

    @pytest.mark.parametrize("config_file", ["toml", "json"], indirect=True)
    def test_get(self, config_file: tuple[dict, pathlib.Path]):
        """
        Tests retrieving environment config from file at environment variable location

        Parameters
        ----------
        config_file : tuple[dict, pathlib.Path]
            See fixture
        """
        contents, path = config_file
        os.environ["OAR_CONFIG_PATH"] = path.stem + path.suffix

        config = oar.EnvConfig.get()
        assert config.dict() == contents
