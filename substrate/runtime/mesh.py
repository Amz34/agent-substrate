from pathlib import Path
from typing import Dict, Any, Optional
from .snapshot import SnapshotEngine
from .sandbox import WarmPoolSandbox
from ..orchestrator.ax_primitives import AXTask, AXWorkspace, AXGateway, AXModel

class SubstrateMesh:
    """Enterprise AX & Substrate Execution Mesh."""
    def __init__(self, data_dir: Optional[Path] = None):
        self.data_dir = data_dir or Path("/home/ubuntu/ops/data/substrate")
        self.snapshot_engine = SnapshotEngine(self.data_dir)
        self.sandbox_engine = WarmPoolSandbox(self.data_dir)

    def dispatch_declarative(self, task: AXTask, workspace: AXWorkspace, gateway: AXGateway, model: AXModel) -> Dict[str, Any]:
        box_path = self.sandbox_engine.provision_sandbox(task.name)
        return {
            "execution_id": f"exec-{task.name}",
            "status": "INITIALIZED",
            "sandbox_path": str(box_path),
            "gateway_barrier": {
                "allowed_domains": gateway.allowed_domains,
                "circuit_breaker": gateway.circuit_breaker_error_threshold
            },
            "model_binding": f"{model.provider}:{model.model_name}",
            "ephemeral": workspace.ephemeral
        }

    def hibernate(self, agent_id: str, current_state: Dict[str, Any]) -> Dict[str, Any]:
        return self.snapshot_engine.snapshot(agent_id, current_state)

    def wake(self, agent_id: str) -> Optional[Dict[str, Any]]:
        return self.snapshot_engine.restore(agent_id)
