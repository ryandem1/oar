from pytest import fixture
from oar.client import Client
from oar.models import EnvConfig


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
