import logging

from pytest import fixture, FixtureRequest, hookimpl, Item, CallInfo, StashKey, CollectReport
from oar.client import Client
from oar.models import EnvConfig, Test, Outcome, Analysis, Resolution


logger = logging.getLogger("oar messenger")

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
def oar_test(request: FixtureRequest, oar_client) -> Test:
    """
    Here is the primary fixture to interact with OAR in PyTest. If this fixture is in the fixture list, the result of
    the test will be uploaded to OAR after the test is complete. The OAR client is designed to not fail if something
    goes wrong with the test, so an upload failure would not cause false positives.

    More info on properties:

    1. Tests begin NotAnalyzed and Unresolved
    2. If a test fails in the setup phase, the test's analysis will be marked as a FalsePositive by default
    3. If a test fails in the call phase, the test will remain NotAnalyzed
    4. If a test fails in the teardown phase, the test's analysis will be marked as a FalsePositive by default
    5. If a test passes, it is by default a TrueNegative
    6. If outcome or analysis were set during the test runtime, they will not be set by the above logic

    Yields
    -------
    test : Test
        Current OAR test to add attributes onto
    """
    test = Test(analysis=Analysis.NotAnalyzed, resolution=Resolution.Unresolved)
    yield test
    # request.node is an "item" because we use the default "function" scope
    report = request.node.stash[phase_report_key]

    if not test.summary:
        test.summary = request.node.name

    if report["setup"].failed:
        if test.outcome is None:
            test.outcome = Outcome.Failed
        if test.analysis == Analysis.NotAnalyzed:
            test.analysis = Analysis.FalsePositive
    elif ("call" not in report) or report["call"].failed:
        if test.outcome is None:
            test.outcome = Outcome.Failed
    elif "teardown" in report and report["teardown"].failed:
        if test.outcome is None:
            test.outcome = Outcome.Failed
        if test.analysis == Analysis.NotAnalyzed:
            test.analysis = Analysis.FalsePositive
    else:
        if test.outcome is None:
            test.outcome = Outcome.Passed
        if test.analysis == Analysis.NotAnalyzed:
            test.analysis = Analysis.TrueNegative

    test_id = oar_client.add_test(test)
    logger.info(f"OAR Test Result ID: {test_id}")


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
