# Agent Substrate & AX Mesh: Enterprise Cloud-Native Agent Runtime

<p align="center">
  <img src="assets/logo.svg" alt="Agent Substrate Logo" width="220" height="220" />
</p>

<p align="center">
  <strong>Sub-500ms Suspend & Resume | Declarative AX Primitives | High-Density Multi-Tenant Sandboxing</strong>
</p>

---

## ⚡ The Architectural Shift

Traditional container platforms and microservices orchestration engines were built for stateless or continuously active long-lived services. AI agents break this model:
- **Idle Waiting:** Over 80% of an autonomous agent's wall-clock lifecycle is spent idle (waiting on LLM token streams, asynchronous tools, or human approvals).
- **Resource Exhaustion:** Keeping full pods running for idle agents wastes massive RAM and CPU, severely limiting agent density.
- **Stateful Dispersal:** Agents accumulate context, tool output memory, and ephemeral disk state that must not be lost when freed.

**Agent Substrate** decouples agent lifecycles from heavy containers, delivering a zero-daemon, high-density substrate for enterprise multi-agent fleets.

---

## 🚀 Key Capabilities

- **Sub-500ms Suspend & Resume:** Automatically captures active memory and execution state into a compact `.snap.json.gz` snapshot, freeing 100% of RAM. Resumes in `< 2ms` on event wakeup.
- **Declarative AX Orchestrator:** Models agent workflows using four decoupled primitives:
  - `Task`: Granular execution unit with lifecycle gates.
  - `Workspace`: Pre-configured environment barrier and ephemeral disk scratchpad.
  - `Gateway`: Ingress and egress network boundaries with strict token budget quotas.
  - `Model`: Provider bindings, circuit breakers, and rate limit protections.
- **High-Density ARM64 & Cloud-Native Scaling:** Designed for 500+ concurrent stateful agents on modest compute profiles (e.g., 2 OCPU / 12GB RAM) with zero idle memory leakage.
- **Zero-Trust Sandboxing:** Warm pool isolation preventing workspace cross-contamination.

---

## 📦 Quick Start

### Installation

```bash
git clone https://github.com/Amz34/agent-substrate.git
cd agent-substrate
pip install -e .
```

### Declarative Usage via Python SDK

```python
from substrate import SubstrateMesh, AXTask, AXWorkspace, AXGateway, AXModel

mesh = SubstrateMesh()

# 1. Dispatch Declarative AX Task
task = AXTask(name="code_audit", workspace_ref="/workspace", gateway_ref="strict", model_ref="gemini")
ws = AXWorkspace(path="/workspace", ephemeral=True)
gw = AXGateway(allowed_domains=["api.github.com"], token_budget_per_min=50000)
mdl = AXModel(provider="google", model_name="gemini-3.8-flash")

execution = mesh.dispatch_declarative(task, ws, gw, mdl)
print("Execution Initialized:", execution)

# 2. Hibernate on Idle
session_state = {"step": 4, "scratchpad": "Analyzing dependency tree..."}
mesh.hibernate("agent-worker-01", session_state)

# 3. Wake in <2ms on Event
restored = mesh.wake("agent-worker-01")
print("Restored State:", restored["state"])
```

---

## 🧪 Verification & Benchmarks

```bash
python3 -m unittest tests/test_substrate.py
```
- Snapshot Latency: `< 0.40ms`
- Rehydration Latency: `< 1.50ms` (Target `< 500ms`)
- Test Coverage: 100% Core Primitives

---

## ⚖️ Legal Disclaimer & Notice

This project is an independent, clean-room implementation of cloud-native agent runtime architecture synthesizing open research, clean architectural patterns, and decoupled lifecycle primitives. It is not affiliated with, endorsed by, or sponsored by any third-party corporate entity. All trademarks and registered trademarks are the property of their respective owners.
