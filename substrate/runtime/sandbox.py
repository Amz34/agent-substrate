import os
import shutil
import tempfile
from pathlib import Path
from typing import Dict, Any

class WarmPoolSandbox:
    """Pre-warmed, ephemeral filesystem and namespace barrier."""
    def __init__(self, pool_dir: Path):
        self.pool_dir = pool_dir / "sandboxes"
        self.pool_dir.mkdir(parents=True, exist_ok=True)

    def provision_sandbox(self, task_id: str) -> Path:
        target = self.pool_dir / task_id
        target.mkdir(parents=True, exist_ok=True)
        # Create standard execution tree
        (target / "workspace").mkdir(exist_ok=True)
        (target / "scratch").mkdir(exist_ok=True)
        (target / "egress_tokens").mkdir(exist_ok=True)
        return target

    def release_sandbox(self, task_id: str):
        target = self.pool_dir / task_id
        if target.exists():
            shutil.rmtree(target, ignore_errors=True)
