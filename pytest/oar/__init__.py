import logging

from .client import Client
from .config import EnvConfig
from .consts import Outcome, Analysis, Resolution
from .report import Report
from .result import Test, TestQuery, TestQueryResult

logging.basicConfig(level=logging.INFO, format="%(message)s")
