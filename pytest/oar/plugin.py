import inspect
import json
import logging
import os
import typing
from datetime import datetime
from pathlib import Path

from pydantic import ValidationError
from pytest import fixture, FixtureRequest, hookimpl, Item, CallInfo, StashKey, CollectReport

from oar.client import Client
from oar.config import EnvConfig
from oar.consts import Outcome, Analysis, Resolution
from oar.models import Test
from oar.report import Report

logger = logging.getLogger("oar")

phase_report_key = StashKey[dict[str, CollectReport]]()  # Stores result data to be available at the fixture level


@hookimpl(tryfirst=True, hookwrapper=True)
def pytest_runtest_makereport(item: Item, call: CallInfo) -> None:
    """
    This hook is implemented so that test outcome information is available at the teardown fixture level.

    Parameters
    ----------
    item : Item
        Current PyTest item

    call : CallInfo
        Metadata related to test call

    Returns
    -------
    None
    """
    _ = call
    # execute all other hooks to obtain the report object
    outcome = yield
    rep = outcome.get_result()

    # store test results for each phase of a call, which can be "setup", "call", "teardown"
    item.stash.setdefault(phase_report_key, {})[rep.when] = rep


@fixture
def oar_test(request: FixtureRequest, oar_config, oar_results, oar_client) -> Test:
    """
    Here is the primary fixture to interact with OAR in PyTest. If this fixture is in the fixture list, the result of
    the test will be uploaded to OAR after the test is complete. The OAR client is designed to not fail if something
    goes wrong with the test, so an upload failure would not cause false positives.

    More info on properties:

    1. Tests default to NotAnalyzed and Unresolved
    2. If a test fails in the setup phase, the test's analysis will be marked as a FalsePositive by default
    3. If a test fails in the call phase, the test will remain NotAnalyzed
    4. If a test fails in the teardown phase, the test's analysis will be marked as a FalsePositive by default
    5. If a test passes, it is by default a TrueNegative
    6. If outcome or analysis were set during the test runtime, they will not be set by the above logic
    7. If a test is set as a specific type, the "type" will automatically be added if not already defined

    Yields
    -------
    test : Test
        Current OAR test to add attributes onto.
    """
    # This will get the type hint of the `oar_test` fixture, defaulting to `oar.Test` if one was not provided
    test_type = typing.get_type_hints(request.function).get(oar_test.__name__, Test)

    # Ensures that a test type is actually being used
    if test_type != Test and Test not in inspect.getmro(test_type):
        raise TypeError(f"`oar_test` fixture type must inherit from `oar.Test` Bases: {inspect.getmro(test_type)}")

    # Initialize a base test container to yield back into the test to use, the specific type will be used later
    test = Test()
    yield test

    # Initializes the specific test_type
    try:
        test = test_type(**test.dict())
    except ValidationError as e:
        logger.warning(e)
        return  # Will not report on validation errors

    # Appends a test type onto a test if it is a specific type
    if type(test) != Test and "type" not in test.__dict__:  # noqa
        test.type = type(test).__name__  # noqa

    # request.node is an "item" because we use the default "function" scope
    report = request.node.stash[phase_report_key]

    if not test.summary:
        test.summary = request.node.name

    if report["setup"].failed:
        test.outcome = test.outcome or Outcome.Failed
        test.analysis = test.analysis or Analysis.FalsePositive
        test.resolution = test.resolution or Resolution.Unresolved
    elif ("call" not in report) or report["call"].failed:
        test.outcome = test.outcome or Outcome.Failed
        test.analysis = test.analysis or Analysis.NotAnalyzed
        test.resolution = test.resolution or Resolution.Unresolved
    elif "teardown" in report and report["teardown"].failed:
        test.outcome = test.outcome or Outcome.Failed
        test.analysis = test.analysis or Analysis.FalsePositive
        test.resolution = test.resolution or Resolution.Unresolved
    else:
        test.outcome = test.outcome or Outcome.Passed
        test.analysis = test.analysis or Analysis.TrueNegative
        test.resolution = test.resolution or Resolution.NotNeeded

    if oar_config.send_results:
        test.id_ = oar_client.add_test(test)
        logger.info(f"OAR Test Result ID: {test.id_}")

    # Add Test to results if store_results is being used
    if oar_config.store_results:
        oar_results += test


@fixture(scope="session")
def oar_results(oar_config) -> Report:
    """
    Stores the results of all OAR tests through the session. Will be enriched through other fixtures.

    If the "store_results is False", this will yield a results object, but not do anything with it afterward.

    Will print a summary of tests at the end.

    Yields
    -------
    results : Results
        OAR results to be enriched or analyzed
    """
    results = Report()

    yield results

    if not oar_config.store_results:
        return

    results.completed_time = str(datetime.utcnow())

    # Output JSON file
    if not oar_config.output_file:
        return

    output_dir = Path(os.getcwd()) / oar_config.output_dir
    output_dir.mkdir(exist_ok=True)
    output_file_name = f"oar-results-{int(datetime.utcnow().timestamp())}.json"
    with (output_dir / output_file_name).open("w") as output_file:
        json.dump(results.dict(by_alias=True), output_file, indent=4)


@fixture(scope="session")
def oar_config() -> EnvConfig:
    """
    Will return the default ``EnvConfig`` object

    Returns
    -------
    config : EnvConfig
        OAR default environment configuration
    """
    config = EnvConfig.get()
    return config


@fixture(scope="session")
def oar_client(oar_config) -> Client:
    """
    Will return an initialized OAR client from the default env config.

    Returns
    -------
    client : Client
        Initialized OAR client with information from the ``oar_config``
    """
    client = Client(base_url=oar_config.host)
    return client
