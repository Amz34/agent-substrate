#!/usr/bin/env python3
import sys
import argparse
import json
from pathlib import Path
from substrate import SubstrateMesh, AXTask, AXWorkspace, AXGateway, AXModel

def main():
    parser = argparse.ArgumentParser(description="Agent Substrate Enterprise CLI (AX)")
    sub = parser.add_subparsers(dest="cmd")

    snap_p = sub.add_parser("hibernate", help="Snapshot agent state to storage")
    snap_p.add_argument("--agent-id", required=True)
    snap_p.add_argument("--state-json", default="{}")

    wake_p = sub.add_parser("wake", help="Restore agent state in sub-500ms")
    wake_p.add_argument("--agent-id", required=True)

    dispatch_p = sub.add_parser("dispatch", help="Dispatch declarative AX task")
    dispatch_p.add_argument("--task-name", required=True)
    dispatch_p.add_argument("--workspace", default="/home/ubuntu")

    args = parser.parse_args()
    mesh = SubstrateMesh()

    if args.cmd == "hibernate":
        state = json.loads(args.state_json)
        res = mesh.hibernate(args.agent_id, state)
        print(json.dumps(res, indent=2))
    elif args.cmd == "wake":
        res = mesh.wake(args.agent_id)
        print(json.dumps(res, indent=2))
    elif args.cmd == "dispatch":
        task = AXTask(name=args.task_name, workspace_ref=args.workspace, gateway_ref="strict", model_ref="gemini")
        ws = AXWorkspace(path=args.workspace)
        gw = AXGateway(allowed_domains=["*"])
        mdl = AXModel(provider="google", model_name="gemini-3.8-flash")
        res = mesh.dispatch_declarative(task, ws, gw, mdl)
        print(json.dumps(res, indent=2))
    else:
        parser.print_help()

if __name__ == "__main__":
    main()
