from dataclasses import dataclass, field
from typing import Dict, Any, Optional
import time

@dataclass
class AXTask:
    name: str
    workspace_ref: str
    gateway_ref: str
    model_ref: str
    metadata: Dict[str, Any] = field(default_factory=dict)

@dataclass
class AXWorkspace:
    path: str
    read_only: bool = False
    ephemeral: bool = True

@dataclass
class AXGateway:
    allowed_domains: list = field(default_factory=list)
    token_budget_per_min: int = 100000
    circuit_breaker_error_threshold: int = 5

@dataclass
class AXModel:
    provider: str
    model_name: str
    temperature: float = 0.2
    max_tokens: int = 4096
