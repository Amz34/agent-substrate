import os
import json
import time
import gzip
from pathlib import Path
from typing import Dict, Any, Optional

class SnapshotEngine:
    """Ultra-low-latency state capture and rehydration."""
    def __init__(self, base_dir: Path):
        self.snapshots_dir = base_dir / "snapshots"
        self.snapshots_dir.mkdir(parents=True, exist_ok=True)

    def snapshot(self, agent_id: str, state: Dict[str, Any]) -> Dict[str, Any]:
        t0 = time.perf_counter()
        target = self.snapshots_dir / f"{agent_id}.snap.json.gz"
        encoded = json.dumps(state, ensure_ascii=False).encode('utf-8')
        compressed = gzip.compress(encoded, compresslevel=1)
        with open(target, "wb") as f:
            f.write(compressed)
        dur_ms = (time.perf_counter() - t0) * 1000.0
        return {
            "status": "SNAPSHOT_COMPACTED",
            "agent_id": agent_id,
            "bytes": len(compressed),
            "latency_ms": round(dur_ms, 2),
            "file": str(target)
        }

    def restore(self, agent_id: str) -> Optional[Dict[str, Any]]:
        t0 = time.perf_counter()
        target = self.snapshots_dir / f"{agent_id}.snap.json.gz"
        if not target.exists():
            return None
        with open(target, "rb") as f:
            decompressed = gzip.decompress(f.read())
            payload = json.loads(decompressed.decode('utf-8'))
        dur_ms = (time.perf_counter() - t0) * 1000.0
        return {
            "status": "RESTORED",
            "agent_id": agent_id,
            "latency_ms": round(dur_ms, 2),
            "state": payload
        }
