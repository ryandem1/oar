from .models import (
    EnvConfig, Outcome, Analysis, Resolution, Test
)
from .client import Client
import logging

logging.basicConfig(level=logging.INFO, format="%(name)s - %(levelname)s - %(message)s")
