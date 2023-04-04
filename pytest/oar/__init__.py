import logging

from .client import Client
from .config import EnvConfig
from .consts import Outcome, Analysis, Resolution
from .results import Results
from .test import Test

logging.basicConfig(level=logging.INFO, format="%(message)s")
